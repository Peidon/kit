class FormBot {
    constructor() {
        // Mapping of field titles to Field arrays
        this.memoryStates = new Map();
        // Map to track seen field IDs for titles quick lookup
        this.seen = new Map();
        // Track learned field IDs to avoid redundant learning
        this.learned = new Set();
        // to track which fields have been filled on the current page
        this.fillStates = new Map();
        this.reviewOverlayId = "formbot-review-overlay";

        this.init();
    }

    init() {
        // initialize memory states from storage
        this.getMemory().then((memory) => {

            if (!memory || !Array.isArray(memory)) {
                return;
            }
            memory.forEach((field) => {
                if (field.title && field.value) {
                    if (!this.fillStates.has(field.title)) {
                        // initialize fill state for each title to 0 (index of the value to use for filling)
                        this.fillStates.set(field.title, 0);
                    }
                    if (this.memoryStates.has(field.title)) {
                        this.memoryStates.get(field.title).push(field);
                    } else {
                        this.memoryStates.set(field.title, [field]);
                    }
                }
            });
            this.memoryStates.forEach((fields) => {
                if (Array.isArray(fields)) {
                    // sort by rank for prioritization
                    fields.sort((a, b) => a.rank - b.rank);
                } else {
                    console.warn("Invalid field entry in memory states:", fields);
                }
            });

        });
    }

    async getMemory() {
        return new Promise((resolve) => {
            // Retrieve the memory object from chrome.storage.local
            chrome.storage.local.get(["memory"], (result) => { resolve(result.memory); });
        });
    }

    async saveMemory(memory) {
        return new Promise((resolve) => {
            // Save the memory object to chrome.storage.local
            // This will overwrite the existing memory with the new one provided
            chrome.storage.local.set({ memory }, resolve);
        });
    }

    snapshotMemory() {
        const entries = [];
        const default_time = Date.now();

        this.memoryStates.forEach((fields, title) => {
            if (!Array.isArray(fields)) {
                return;
            }

            fields.forEach((field, index) => {
                entries.push({
                    id: `${title}::${field.rank}::${index}`,
                    title,
                    value: field.value,
                    rank: Number.isFinite(field.rank) ? field.rank : index,
                    create_at: field.create_at || default_time
                });
            });
        });

        return entries.sort((a, b) => {
            if (a.create_at !== b.create_at) {
                return b.create_at - a.create_at;
            }
            if (a.rank !== b.rank) {
                return a.rank - b.rank;
            }
            return a.title.localeCompare(b.title);
        });
    }

    rehydrateMemory(memory) {
        this.memoryStates = new Map();
        this.fillStates = new Map();
        const create_at = Date.now();
        memory.forEach((field) => {
            if (!field?.title || !field?.value) {
                return;
            }

            const normalizedField = new Field(
                field.title,
                field.value,
                Number.isFinite(field.rank) ? field.rank : 0,
                field.create_at || create_at
            );

            if (!this.memoryStates.has(normalizedField.title)) {
                this.memoryStates.set(normalizedField.title, []);
                this.fillStates.set(normalizedField.title, 0);
            }

            this.memoryStates.get(normalizedField.title).push(normalizedField);
        });

        this.memoryStates.forEach((fields) => {
            fields.sort((a, b) => a.rank - b.rank);
        });
    }

    async persistMemoryEntries(entries) {
        const sanitized = entries
            .map((entry, index) => ({
                title: (entry.title || "").trim(),
                value: (entry.value || "").trim(),
                rank: Number.isFinite(entry.rank) ? entry.rank : index
            }))
            .filter((entry) => entry.title && entry.value)
            .sort((a, b) => {
                if (a.rank !== b.rank) {
                    return a.rank - b.rank;
                }
                return a.title.localeCompare(b.title);
            });

        this.rehydrateMemory(sanitized);
        await this.saveMemory(sanitized);
    }

    buildDetectRequestBody(inputs) {
        const params = [];
        inputs.forEach((input) => {

            const f_id = field_id(input);
            if (!f_id || this.seen.has(f_id)) {
                return;
            }

            const labels = collectLabels(input);
            if (labels.length === 0) {
                return;
            }

            params.push({
                id: f_id,
                labels: labels
            });
        });

        return params;
    }

    async linkTitles(fieldsLabels) {
        if (fieldsLabels.length === 0) {
            return new Map();
        }
        const targetTitles = Array.from(this.memoryStates.keys());
        const linked = await extensionApiFetch(
            `/link_titles`,
            {
                method: 'POST',
                body: JSON.stringify({ "source": fieldsLabels, "target": targetTitles })
            }
        );

        if (!linked || !linked.result) {
            throw new Error("Invalid response from link_titles API");
        }

        return new Map(Object.entries(linked.result));
    }

    titleFromSeen(fieldId, titles) {
        if (this.seen.has(fieldId)) {
            return this.seen.get(fieldId);
        }
        const title = titles.get(fieldId) || "";
        this.seen.set(fieldId, title);
        return title;
    }

    fillText(input, value) {
        input.focus();
        input.value = value;
        input.classList.add("formbot-filled");
        input.dispatchEvent(new Event("input", { bubbles: true }));
        input.dispatchEvent(new Event("change", { bubbles: true }));
        input.blur();
    }

    async fill() {
        const inputs = Array.from(document.querySelectorAll("input, textarea"));
        const inputsToFill = inputs.filter(input => {
            if (input.disabled || input.readOnly || input.value.trim() !== "" || 
            input.type === "file" || input.type === "checkbox" || input.type === "radio") {
                return false;
            }
            const f_id = field_id(input);
            return f_id;
        });

        if (inputsToFill.length === 0) {
            alert("I already have filled all available fields.", "OK");
            return;
        }
        const filled = new Set();
        const params = this.buildDetectRequestBody(inputsToFill);
        await this.linkTitles(params).then((titles) => {
            inputsToFill.forEach((input) => {
                const f_id = field_id(input);
                if (!f_id) {
                    return;
                }
                const title = this.titleFromSeen(f_id, titles);
                if (!this.memoryStates.has(title)) {
                    return;
                }
                if (!this.memoryStates.has(title) || this.memoryStates.get(title).length === 0) {
                    return; // no known values for this title
                }
                const fields = this.memoryStates.get(title);
                const fillIndex = this.fillStates.get(title) || 0;
                if (fillIndex >= fields.length) {
                    return; // no more values to fill for this title
                }
                const fieldToFill = fields[fillIndex];
                this.fillText(input, fieldToFill.value);
                this.fillStates.set(title, fillIndex + 1); // move to next value for next time
                filled.add(f_id);
            });
        }).catch((error) => {
            console.error("Failed to autofill:", error);
        }).then(()=>{
            if (filled.size == 0) {
                alert("No matching memory entries found to auto fill. Please fill manually and click the [Learn] button.", "OK");
            }
        });
    }

    async learn() {
        const inputs = Array.from(document.querySelectorAll("input, textarea"));
        const inputsToLearn = inputs.filter(input => {
            const f_id = field_id(input);
            if (input.value.trim() === "" || input.disabled || input.readOnly || 
                input.type === "file" || input.type === "checkbox" || input.type === "radio" || 
                this.learned.has(f_id)) {
                return false;
            }
            return f_id;
        });

        if (inputsToLearn.length === 0) {
            alert("No new inputs to learn from. Please input some information first.", "OK");
            return;
        }
        const params = this.buildDetectRequestBody(inputs);
        const create_time = Date.now();
        await this.linkTitles(params).then((titles) => {
            inputsToLearn.forEach((input) => {
                const f_id = field_id(input);
                if (!f_id || this.learned.has(f_id)) {
                    return;
                }
                const title = this.titleFromSeen(f_id, titles);
                if(title==""){
                    return;
                }
                const value = input.value;
                if (this.memoryStates.has(title)) {
                    const existing = this.memoryStates.get(title);
                    if (existing.some(field => field.value === value)) {
                        this.learned.add(f_id);
                        return; // already have this value for the title, skip
                    }
                    const newField = new Field(title, value, existing[existing.length - 1].rank + 1, create_time);
                    existing.push(newField);
                } else {
                    const newField = new Field(title, value, 0, create_time);
                    this.memoryStates.set(title, [newField]);
                }
                this.learned.add(f_id);
            });
        }).catch((error) => {
            console.error("Failed to learn:", error);
        }).then(async () => {

            const memoryObjs = [];
            const deduped = new Set();
            this.memoryStates.forEach((fields) => {
                if (!Array.isArray(fields)) {
                    return;
                }
                fields.forEach((field) => {
                    const key = `${field.title}::${field.value}`;
                    if (deduped.has(key)) {
                        return;
                    }
                    deduped.add(key);
                    memoryObjs.push(field);
                });
            });

            await this.saveMemory(memoryObjs);
        });
    }

    review() {
        const existingOverlay = document.getElementById(this.reviewOverlayId);
        if (existingOverlay) {
            existingOverlay.remove();
        }

        const overlay = document.createElement("div");
        overlay.id = this.reviewOverlayId;
        overlay.className = "formbot-review-overlay";

        const modal = document.createElement("section");
        modal.className = "formbot-review-modal";
        overlay.appendChild(modal);

        const searchWrap = document.createElement("div");
        searchWrap.className = "formbot-review-search";
        const searchInput = document.createElement("input");
        searchInput.type = "search";
        searchInput.placeholder = "Search...";
        searchInput.className = "formbot-review-search-input";
        searchWrap.appendChild(searchInput);
        modal.appendChild(searchWrap);

        const body = document.createElement("div");
        body.className = "formbot-review-body";
        modal.appendChild(body);

        const rows = document.createElement("div");
        rows.className = "formbot-review-rows";
        body.appendChild(rows);

        const footer = document.createElement("div");
        footer.className = "formbot-review-footer";
        const saveButton = document.createElement("button");
        saveButton.type = "button";
        saveButton.className = "formbot-review-save";
        saveButton.textContent = "Save";
        const nextButton = document.createElement("button");
        nextButton.type = "button";
        nextButton.className = "formbot-review-next";
        nextButton.textContent = "Next Rank";
        footer.appendChild(saveButton);
        footer.appendChild(nextButton);
        modal.appendChild(footer);

        document.body.appendChild(overlay);

        let page = 0;
        const pageSize = 20;
        const entries = this.snapshotMemory();
        let filteredEntries = entries.slice();
        let closed = false;
        const markedForDeletion = new Set();

        const scheduleSave = () => {
            const nextEntries = entries.filter((entry) => !markedForDeletion.has(entry.id));
            return this.persistMemoryEntries(nextEntries);
        };

        const refreshPageBounds = () => {
            const totalPages = Math.max(1, Math.ceil(filteredEntries.length / pageSize));
            if (page >= totalPages) {
                page = totalPages - 1;
            }
            nextButton.disabled = filteredEntries.length <= pageSize;
            nextButton.textContent = totalPages > 1
                ? `Next Page (${page + 1}/${totalPages})`
                : "Next Page";
        };

        const createTrashButton = (entryId, row) => {
            const button = document.createElement("button");
            button.type = "button";
            button.className = "formbot-review-delete";
            button.setAttribute("aria-label", "Delete memory row");
            button.innerHTML = `
                <svg viewBox="0 0 24 24" aria-hidden="true">
                    <path d="M9 3h6l1 2h4v2H4V5h4l1-2zm-2 6h2v8H7V9zm4 0h2v8h-2V9zm4 0h2v8h-2V9zM6 7h12v12a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V7z"/>
                </svg>
            `;
            button.addEventListener("click", () => {
                if (markedForDeletion.has(entryId)) {
                    markedForDeletion.delete(entryId);
                    row.classList.remove("formbot-review-row--pending-delete");
                    button.classList.remove("formbot-review-delete--active");
                } else {
                    markedForDeletion.add(entryId);
                    row.classList.add("formbot-review-row--pending-delete");
                    button.classList.add("formbot-review-delete--active");
                }
            });
            return button;
        };

        const renderRows = () => {
            refreshPageBounds();
            rows.replaceChildren();

            const start = page * pageSize;
            const visibleEntries = filteredEntries.slice(start, start + pageSize);

            if (visibleEntries.length === 0) {
                const emptyState = document.createElement("div");
                emptyState.className = "formbot-review-empty";
                emptyState.textContent = "No memory entries match this search.";
                rows.appendChild(emptyState);
                return;
            }

            visibleEntries.forEach((entry) => {
                const row = document.createElement("div");
                row.className = "formbot-review-row";

                const keyInput = document.createElement("input");
                keyInput.type = "text";
                keyInput.className = "formbot-review-field";
                keyInput.placeholder = "Key";
                keyInput.value = entry.title;
                keyInput.addEventListener("input", (event) => {
                    entry.title = event.target.value;
                });

                const valueInput = document.createElement("input");
                valueInput.type = "text";
                valueInput.className = "formbot-review-field";
                valueInput.placeholder = "Value";
                valueInput.value = entry.value;
                valueInput.addEventListener("input", (event) => {
                    entry.value = event.target.value;
                });

                if (markedForDeletion.has(entry.id)) {
                    row.classList.add("formbot-review-row--pending-delete");
                }

                row.appendChild(keyInput);
                row.appendChild(valueInput);
                row.appendChild(createTrashButton(entry.id, row));
                rows.appendChild(row);
            });
        };

        const applyFilter = () => {
            const needle = normalize(searchInput.value || "");
            filteredEntries = entries.filter((entry) => {
                if (!needle) {
                    return true;
                }
                return normalize(`${entry.title} ${entry.value}`).includes(needle);
            });
            page = 0;
            renderRows();
        };

        searchInput.addEventListener("input", applyFilter);
        saveButton.addEventListener("click", async () => {
            saveButton.disabled = true;
            saveButton.textContent = "Saving...";
            try {
                await scheduleSave();
                const remainingEntries = entries.filter((entry) => !markedForDeletion.has(entry.id));
                entries.splice(0, entries.length, ...remainingEntries);
                markedForDeletion.clear();
                applyFilter();
            } catch (error) {
                console.error("Failed to save reviewed memory:", error);
            } finally {
                saveButton.disabled = false;
                saveButton.textContent = "Save";
            }
        });
        nextButton.addEventListener("click", () => {
            const totalPages = Math.max(1, Math.ceil(filteredEntries.length / pageSize));
            page = (page + 1) % totalPages;
            renderRows();
        });

        const closeOnEscape = (event) => {
            if (event.key === "Escape") {
                closeReview();
            }
        };

        const closeReview = async () => {
            if (closed) {
                return;
            }
            closed = true;
            document.removeEventListener("keydown", closeOnEscape);
            overlay.remove();
        };

        overlay.addEventListener("click", (event) => {
            if (event.target === overlay) {
                void closeReview();
            }
        });

        document.addEventListener("keydown", closeOnEscape);
        applyFilter();
        searchInput.focus();
    }
}

var bot = new FormBot();
class Field {
    /**
     * 
     * @param {string} title The title of the field, e.g. "first name", "email", "company"
     * @param {string} value The value of the field, e.g. "John", "Acme Inc."
     * @param {number} rank The importance rank of the field, used for prioritization when filling forms. Lower is more important. Default is 0.
     * @param {number} create_at The create_at when the field was created.
     */
    constructor(title, value, rank = 0, create_time = Date.now()) {
        this.title = title;
        this.value = value;
        this.rank = rank;
        this.create_at = create_time;
    }
}

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
        this.reviewStyleId = "formbot-review-style";

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
                create_at
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
        const seen = new Set();
        inputs.forEach((input) => {

            const f_id = field_id(input);
            if (!f_id || this.seen.has(f_id)) {
                return;
            }

            const labels = collectLabels(input, seen);
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
        const title = titles.get(fieldId) || fieldId;
        this.seen.set(fieldId, title);
        return title;
    }

    fillText(input, value) {
        input.focus();
        input.value = value;
        input.style.backgroundColor = "#e6ffed";
        input.dispatchEvent(new Event("input", { bubbles: true }));
        input.dispatchEvent(new Event("change", { bubbles: true }));
        input.blur();
    }

    fill() {
        const inputs = Array.from(document.querySelectorAll("input, textarea"));
        const params = this.buildDetectRequestBody(inputs);
        const inputsToFill = inputs.filter(input => { return input.value.trim() == ""; }).filter(input => {
            const f_id = field_id(input);
            return f_id;
        });

        if (inputsToFill.length === 0) {
            alert("I already have filled all available fields.", "OK");
            return;
        }
        const filled = new Set();
        this.linkTitles(params).then((titles) => {
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
            console.log("seen inputs: ", this.seen)
            if (filled.size == 0) {
                alert("No matching memory entries found to auto fill. Please fill manually and click the [Learn] button.", "OK");
            }
        });
    }

    learn() {
        const inputs = Array.from(document.querySelectorAll("input, textarea"));
        const params = this.buildDetectRequestBody(inputs);
        const create_time = Date.now();
        const inputsToLearn = inputs.filter(input => { return input.value.trim() !== ""; }).filter(input => {
            const f_id = field_id(input);
            return f_id && !this.learned.has(f_id);
        });

        if (inputsToLearn.length === 0) {
            alert("No new inputs to learn from. Please input some information first.", "OK");
            return;
        }
        this.linkTitles(params).then((titles) => {
            inputsToLearn.forEach((input) => {
                const f_id = field_id(input);
                if (!f_id || this.learned.has(f_id)) {
                    return;
                }
                const title = this.titleFromSeen(f_id, titles);
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

            console.log("Memory states before saving:", this.memoryStates.keys());
            const memoryObjs = [];
            this.memoryStates.forEach((fields) => {
                memoryObjs.push(...fields);
            });

            await this.saveMemory(memoryObjs);
            console.log("Memory states after learning:", memoryObjs);
        });
    }

    review() {
        const existingOverlay = document.getElementById(this.reviewOverlayId);
        if (existingOverlay) {
            existingOverlay.remove();
        }

        this.injectReviewStyles();

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

    injectReviewStyles() {
        if (document.getElementById(this.reviewStyleId)) {
            return;
        }

        const style = document.createElement("style");
        style.id = this.reviewStyleId;
        style.textContent = `
            .formbot-review-overlay {
                position: fixed;
                inset: 0;
                z-index: 2147483647;
                padding: 44px 32px 24px;
                background: rgba(191, 191, 191, 0.55);
                backdrop-filter: blur(2px);
                box-sizing: border-box;
            }

            .formbot-review-modal {
                height: 100%;
                border-radius: 22px;
                background: #ffffff;
                box-shadow: 0 24px 60px rgba(15, 23, 42, 0.16);
                overflow: hidden;
                display: flex;
                flex-direction: column;
                font-family: "Segoe UI", "Helvetica Neue", Arial, sans-serif;
            }

            .formbot-review-search {
                padding: 32px 32px 22px;
                border-bottom: 1px solid #e4e8f0;
            }

            .formbot-review-search-input,
            .formbot-review-field {
                width: 100%;
                box-sizing: border-box;
                border: 2px solid #d5dce8;
                border-radius: 18px;
                background: #ffffff;
                color: #505050;
                font-size: 24px;
                line-height: 1.2;
                outline: none;
            }

            .formbot-review-search-input {
                height: 82px;
                padding: 0 32px;
            }

            .formbot-review-search-input::placeholder,
            .formbot-review-field::placeholder {
                color: #8d8d8d;
            }

            .formbot-review-body {
                flex: 1;
                overflow: hidden;
                padding: 34px 20px 24px 24px;
            }

            .formbot-review-rows {
                height: 100%;
                overflow-y: auto;
                overflow-x: hidden;
                direction: rtl;
                padding-left: 8px;
                scrollbar-width: thin;
                scrollbar-color: #c8cfda transparent;
            }

            .formbot-review-rows::-webkit-scrollbar {
                width: 12px;
            }

            .formbot-review-rows::-webkit-scrollbar-track {
                background: transparent;
            }

            .formbot-review-rows::-webkit-scrollbar-thumb {
                background: #c8cfda;
                border-radius: 999px;
            }

            .formbot-review-row,
            .formbot-review-empty {
                direction: ltr;
            }

            .formbot-review-row {
                display: grid;
                grid-template-columns: minmax(220px, 1fr) minmax(220px, 1fr) 50px;
                gap: 22px;
                align-items: center;
                margin-bottom: 24px;
            }

            .formbot-review-row--pending-delete .formbot-review-field {
                border-color: #ffbaba;
                background: #fff6f6;
                color: #b44545;
            }

            .formbot-review-field {
                height: 84px;
                padding: 0 26px;
            }

            .formbot-review-field:focus,
            .formbot-review-search-input:focus {
                border-color: #9eb4ff;
                box-shadow: 0 0 0 4px rgba(84, 112, 255, 0.12);
            }

            .formbot-review-delete {
                width: 44px;
                height: 44px;
                border: none;
                background: transparent;
                color: #959191ff;
                cursor: pointer;
                padding: 0;
                display: inline-flex;
                align-items: center;
                justify-content: center;
            }

            .formbot-review-delete svg {
                width: 34px;
                height: 34px;
                fill: currentColor;
            }

            .formbot-review-delete:hover {
                color: #dd2020;
            }

            .formbot-review-delete--active {
                color: #ff3434;
            }

            .formbot-review-empty {
                height: 100%;
                display: flex;
                align-items: center;
                justify-content: center;
                color: #8d8d8d;
                font-size: 22px;
            }

            .formbot-review-footer {
                padding: 0 32px 28px;
                display: flex;
                gap: 14px;
                justify-content: flex-end;
            }

            .formbot-review-save,
            .formbot-review-next {
                min-width: 200px;
                height: 56px;
                padding: 0 28px;
                border: 2px solid #d5dce8;
                border-radius: 999px;
                background: #ffffff;
                color: #4b5563;
                font-size: 18px;
                font-weight: 600;
                cursor: pointer;
            }

            .formbot-review-save {
                border-color: #9eb4ff;
                background: #5470ff;
                color: #ffffff;
            }

            .formbot-review-save:hover:not(:disabled) {
                background: #4360f0;
                border-color: #4360f0;
            }

            .formbot-review-next:hover:not(:disabled) {
                border-color: #9eb4ff;
                color: #1f3dab;
            }

            .formbot-review-save:disabled,
            .formbot-review-next:disabled {
                opacity: 0.55;
                cursor: default;
            }
        `;

        document.head.appendChild(style);
    }
}

bot = new FormBot();

function field_id(input) {
    if (input.id == "" && input.name == "") {
        return null;
    }
    // input unique identifier based on its attributes and position in the DOM
    // parent element index is used to differentiate between multiple similar fields (e.g. multiple "email" fields for different purposes)
    const parts = [];
    let el = input.parentElement;
    let levels = 10; // limit how far up the DOM we go to avoid overly long identifiers
    while (el && el.parentElement && levels > 0) {
        const siblings = Array.from(el.parentElement.children);
        // only include index if there are multiple similar siblings, and not too many to avoid noise
        if (siblings.length > 1 && siblings.length < 10) {
            const index = siblings.indexOf(el);
            parts.unshift(`${levels}.${index}`);
        }
        el = el.parentElement;
        levels--;
    }
    parts.unshift(input.id || input.name || "");
    return parts.join("_").trim();
}

function splitIntoPhrases(text) {
    if (!text) {
        return [];
    }
    text = text.replace(/[^a-zA-Z\s]+/g, "#").trim();
    return text.split(/#/).map((part) => {
        part = part.replace(/([a-z])([A-Z])/g, '$1 $2');
        part = part.replace(/([A-Z]+)([A-Z][a-z])/g, '$1 $2');
        return part.toLowerCase();
    });
}

function collectLabels(input, text_seen = new Set()) {

    const labels = [];
    const seen = new Set();

    const pushLabel = (text) => {
        if (!text || text.trim() === "" || text_seen.has(text)) {
            return;
        }
        text_seen.add(text);
        const parts = splitIntoPhrases(text).filter((part) => !seen.has(part) && part.length > 1);
        seen.add(...parts);
        labels.push(...parts);
    };

    pushLabel(input.placeholder);
    pushLabel(input.name);
    pushLabel(input.id);

    // 1. Standard label
    if (input.id) {
        const label = document.querySelector(`label[for="${input.id}"]`);
        pushLabel(label?.innerText);
    }

    // 2. aria-label
    pushLabel(input.getAttribute("aria-label"));

    // 3. aria-labelledby
    const labelledBy = input.getAttribute("aria-labelledby");
    if (labelledBy) {
        labelledBy.split(/\s+/).forEach((id) => {
            pushLabel(document.getElementById(id)?.innerText);
        });
    }

    // 4. Fieldset legend is common for textarea / radio / checkbox groups
    const fieldset = input.closest("fieldset");
    if (fieldset) {
        pushLabel(fieldset.querySelector("legend")?.innerText);
    }

    const nearbySelector = [
        "legend",
        "label",
        "[role='heading']",
        "h1",
        "h2",
        "h3",
        "h4",
        "h5",
        "h6",
        ".label",
        ".title",
        "[class*='heading' i]",
        "[class*='title' i]"
    ].join(", ");

    // 5. Walk up DOM and find nearby text
    let el = input.parentElement;
    for (let i = 0; i < 4; i++) {
        if (!el) break;

        const directHeading = el.querySelector?.(`:scope > ${nearbySelector.replace(/, /g, ", :scope > ")}`);
        pushLabel(directHeading?.innerText);

        const textNodes = Array.from(el.childNodes || [])
            .filter(n => n.nodeType === Node.TEXT_NODE)
            .map(n => n.textContent.trim())
            .join(" ");
        pushLabel(textNodes);

        let prev = el.previousElementSibling;
        let steps = 0;
        while (prev && steps < 3) {
            pushLabel(prev.querySelector?.(nearbySelector)?.innerText);
            if (!prev.querySelector?.("input, textarea, select")) {
                pushLabel(prev.innerText);
            }
            prev = prev.previousElementSibling;
            steps++;
        }

        el = el.parentElement;
    }

    return labels;
}

function normalize(text) {
    if (!text) {
        return "";
    }
    // for CaptalizedWords, like "URLAddress", we want to split into "URL Address" to improve matching, but we want to keep all uppercase words together, so we use a regex that splits before uppercase letters but not between consecutive uppercase letters, so "URLAddress" becomes "URL Address" but "MyURLAddress" becomes "My URL Address"
    text = text.replace(/([a-z])([A-Z])/g, '$1 $2'); // split camelCase and PascalCase
    text = text.replace(/([A-Z]+)([A-Z][a-z])/g, '$1 $2'); // split before uppercase followed by lowercase, but keep consecutive uppercase together
    // convert to lowercase and remove non-alphanumeric characters for better matching
    return text.toLowerCase().replace(/[^a-z0-9]/g, " ").trim();
}

async function extensionApiFetch(path, options = {}) {
    return new Promise((resolve, reject) => {
        chrome.runtime.sendMessage(
            {
                action: 'API_FETCH',
                path,
                options
            },
            (response) => {
                if (chrome.runtime.lastError) {
                    reject(new Error(chrome.runtime.lastError.message));
                    return;
                }

                if (!response?.ok) {
                    reject(new Error(response?.error || 'Extension API request failed'));
                    return;
                }

                resolve(response.data);
            }
        );
    });
}

chrome.runtime.onMessage.addListener((msg) => {
    if (msg.action === "FILL_FORM") {
        bot.fill();
    }
    if (msg.action === "LEARN") {
        bot.learn();
    }
    if (msg.action === "REVIEW") {
        // Popup a window with the current memory for review and editing
        bot.review();
    }
});

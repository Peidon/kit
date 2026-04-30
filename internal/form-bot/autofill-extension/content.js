class Field {
    /**
     * 
     * @param {string} title The title of the field, e.g. "first name", "email", "company"
     * @param {string} value The value of the field, e.g. "John", "Acme Inc."
     * @param {number} rank The importance rank of the field, used for prioritization when filling forms. Lower is more important. Default is 0.
     */
    constructor(title, value, rank = 0) {
        this.title = title;
        this.value = value;
        this.rank = rank;
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

        this.init();
    }

    init() {
        // initialize memory states from storage
        this.getMemory().then((memory) => {

            if (!memory || !Array.isArray(memory)) {
                console.warn("No valid memory found in storage, starting with empty memory.");
                return;
            }
            memory.forEach((field) => {
                if (field.title && field.value) {
                    if (!this.fillStates.has(field.title)) {
                        // initialize fill state for each title to 0 (index of the value to use for filling)
                        this.fillStates.set(field.title, 0);
                    }
                    if (this.memoryStates.has(field.title)) {
                        this.memoryStates.get(field.title).push(new Field(field.title, field.value, field.rank));
                    } else {
                        this.memoryStates.set(field.title, [new Field(field.title, field.value, field.rank)]);
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

            console.log("Memory states initialized:", this.memoryStates);

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
        this.linkTitles(params).then((titles) => {
            inputs.filter(input => { return input.value.trim() == ""; }).forEach((input) => {
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
            });
        }).catch((error) => {
            console.error("Failed to detect field:", error);
        }).then(()=>{
            console.log("seen inputs: ", this.seen)
        });
    }

    learn() {
        const inputs = Array.from(document.querySelectorAll("input, textarea"));
        const params = this.buildDetectRequestBody(inputs);
        this.linkTitles(params).then((titles) => {
            inputs.filter(input => { return input.value.trim() !== ""; }).forEach((input) => {
                const f_id = field_id(input);
                if (!f_id || this.learned.has(f_id)) {
                    return;
                }
                const title = this.titleFromSeen(f_id, titles);
                const value = input.value;
                if (this.memoryStates.has(title)) {
                    const existing = this.memoryStates.get(title);
                    if (existing.some(field => field.value === value)) {
                        return; // already have this value for the title, skip
                    }
                    const newField = new Field(title, value, existing[existing.length - 1].rank + 1)
                    existing.push(newField);
                } else {
                    const newField = new Field(title, value);
                    this.memoryStates.set(title, [newField]);
                }
                this.learned.add(f_id);
            });
        }).catch((error) => {
            console.error("Failed to detect field:", error);
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

function collectLabels(input) {

    const labels = [];
    const seen = new Set();

    const pushLabel = (text) => {
        const value = normalize(text);
        if (!value || seen.has(value)) {
            return;
        }
        seen.add(value);
        labels.push(value);
    };

    pushLabel(input.placeholder);
    pushLabel(input.name);

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

// convert hump word to normal word
function splitHumpWord(word) {
    if (!word) {
        return "";
    }
    const parts = word.split(/(?=[A-Z])/g); // split before uppercase letters
    return parts.join(' ');
}

function normalize(text) {
    text = splitHumpWord(text);
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
});

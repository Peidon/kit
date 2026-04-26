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

    fill() {
        const inputs = Array.from(document.querySelectorAll("input, textarea"));
        const params = this.buildDetectRequestBody(inputs);
        detectFieldsMean(params).then((titles) => {
            inputs.filter(input => { return input.value.trim() == ""; }).forEach((input) => {
                const f_id = field_id(input);
                if (!f_id) {
                    return;
                }
                if (!this.seen.has(f_id)) {
                    const title = titles.get(f_id) || f_id;
                    this.seen.set(f_id, title);
                }
                const title = this.seen.get(f_id) || f_id;
                if (!this.memoryStates.has(title)) {
                    return;
                }
                const fields = this.memoryStates.get(title);
                const fillIndex = this.fillStates.get(title) || 0;
                if (fillIndex >= fields.length) {
                    return; // no more values to fill for this title
                }
                const fieldToFill = fields[fillIndex];
                input.focus();
                input.value = fieldToFill.value;
                input.style.backgroundColor = "#e6ffed";
                input.dispatchEvent(new Event("input", { bubbles: true }));
                input.dispatchEvent(new Event("change", { bubbles: true }));
                this.fillStates.set(title, fillIndex + 1); // move to next value for next time
            });
        }).catch((error) => {
            console.error("Failed to detect field:", error);
        });
    }

    learn() {
        const inputs = Array.from(document.querySelectorAll("input, textarea"));
        const params = this.buildDetectRequestBody(inputs);
        detectFieldsMean(params).then((titles) => {
            inputs.filter(input => { return input.value.trim() !== ""; }).forEach((input) => {
                const f_id = field_id(input);
                if (!f_id || this.learned.has(f_id)) {
                    return;
                }
                if (!this.seen.has(f_id)) {
                    const title = titles.get(f_id) || f_id;
                    this.seen.set(f_id, title);
                }
                const title = this.seen.get(f_id) || f_id;
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

    const placeholder = normalize(input.placeholder);
    if (placeholder) {
        labels.push(placeholder);
    }

    const name = normalize(input.name);
    if (name) {
        labels.push(name);
    }

    // 1. Standard label
    if (input.id) {
        const label = document.querySelector(`label[for="${input.id}"]`);
        const labelValue = normalize(label?.innerText);
        if (labelValue) {
            labels.push(labelValue);
        }
    }

    // 2. aria-label
    const aria = normalize(input.getAttribute("aria-label"));
    if (aria) labels.push(aria);

    // 3. Walk up DOM and find nearby text
    let el = input;
    for (let i = 0; i < 3; i++) {
        if (!el) break;

        // Check for text nodes directly under the parent element
        // This is important for modern UIs where labels are often siblings rather than parents
        const textNodes = Array.from(el.parentElement?.childNodes || [])
            .filter(n => n.nodeType === Node.TEXT_NODE)
            .map(n => n.textContent.trim())
            .join(" ");

        if (textNodes.length > 0 && textNodes.length < 100) {
            labels.push(normalize(textNodes));
        }

        el = el.parentElement;
    }

    // 4. Look for preceding sibling text
    let prev = input.previousElementSibling;
    while (prev) {
        if (prev.innerText && prev.innerText.length < 100) {
            labels.push(normalize(prev.innerText));
        }
        prev = prev.previousElementSibling;
    }

    // 5. nearby text (very important for modern UIs)
    const parent = input.parentElement;
    if (parent) {
        const text = parent.innerText;
        if (text && text.length < 100) {
            labels.push(normalize(text));
        }
    }

    return labels;
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

/**
 * Detect the meaning of form fields by sending their labels to the API
 * @param {array} fields_labels An array of field labels to be detected
 * @returns {object} A mapping of field labels to their detected meanings
 */
async function detectFieldsMean(fields_labels) {
    if (!Array.isArray(fields_labels) || fields_labels.length === 0) {
        return new Map();
    }
    const detected = await extensionApiFetch(
        `/detect_fields`,
        {
            method: 'POST',
            body: JSON.stringify({ "fields": fields_labels })
        }
    );

    if (!detected || !detected.result) {
        throw new Error("Invalid response from detect_fields API");
    }

    return new Map(Object.entries(detected.result));
}

chrome.runtime.onMessage.addListener((msg) => {
    if (msg.action === "FILL_FORM") {
        bot.fill();
    }
    if (msg.action === "LEARN") {
        bot.learn();
    }
});

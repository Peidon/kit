class Field {
    /**
     * 
     * @param {string} title The title of the field, e.g. "first name", "email", "company"
     * @param {string} value The value of the field, e.g. "John", "Acme Inc."
     * @param {number} rank The importance rank of the field, used for prioritization when filling forms. Lower is more important. Default is 0.
     * @param {number} create_at The create_at when the field was created.
     */
    constructor(title, value, rank = 0, create_at = Date.now()) {
        this.title = title;
        this.value = value;
        this.rank = rank;
        this.create_at = create_at;
    }
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

function collectLabels(input) {

    const labels = [];
    const seen = new Set();

    const pushLabel = (text) => {
        text = text?.replace(/\s+/g, " ").trim();
        if (!text || seen.has(text)) {
            return;
        }
        seen.add(text);
        const parts = splitIntoPhrases(text).filter((part) => !seen.has(part) && part.length > 1);
        seen.add(...parts);
        labels.push(...parts);
    };

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
        if (directHeading?.innerText) {
            break;
        }

        const textNodes = Array.from(el.childNodes || [])
            .filter(n => n.nodeType === Node.TEXT_NODE)
            .map(n => n.textContent.trim())
            .join(" ");
        pushLabel(textNodes);
        if (textNodes) {
            break;
        }

        let prev = el.previousElementSibling;
        let steps = 0;
        while (prev && steps < 3) {
            pushLabel(prev.querySelector?.(nearbySelector)?.innerText);
            if (prev.querySelector?.(nearbySelector)?.innerText) {
                break;
            }
            if (!prev.querySelector?.("input, textarea, select")) {
                pushLabel(prev.innerText);
                if (prev.innerText) {
                    break;
                }
            }
            prev = prev.previousElementSibling;
            steps++;
        }

        el = el.parentElement;
    }

    pushLabel(input.name);
    pushLabel(input.id);

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

chrome.runtime.onMessage.addListener((msg, sender, sendResponse) => {

    if (msg.action === "FILL_FORM") {
        
        // Execute the async operation
        (async () => {
           try {
            if (msg.action === "FILL_FORM") {
                // 1. Wait for bot.fill() to complete (which awaits get_data())
                await bot.fill(); 
            }
            // 2. Send a success response back to the popup
            sendResponse({ status: "fill success" });
        } catch (error) {
            console.error(error, ". sender id:", sender.id);
            sendResponse({ status: "fill error", error: error.message });
        }
        })();

        return true; // 3. CRUCIAL! Keeps the channel open so sendResponse can fire later
    }

    if (msg.action === "LEARN") {
        (async () => {
           try {
            if (msg.action === "LEARN") {
                await bot.learn();
            }
            sendResponse({ status: "learn success" });
        } catch (error) {
            console.error(error, ". sender id:", sender.id);
            sendResponse({ status: "error", error: error.message });
        }
        })();

        return true; // 3. CRUCIAL! Keeps the channel open so sendResponse can fire later
    }
    if (msg.action === "REVIEW") {
        // Popup a window with the current memory for review and editing
        bot.review();
    }
});

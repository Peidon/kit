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


function fetchInputsLabels(inputs) {
    const fieldsList = [];
    inputs.forEach((input) => {

        const f_id = normalize(field_id(input));
        if (!f_id) {
            return;
        }

        labels = collectLabels(input);
        if (labels.length === 0) {
            return;
        }

        fieldsList.push({
            id: f_id,
            labels: labels
        });
    });

    detectFieldsMean(fieldsList).then(({ result }) => {
        console.log("Detected field meanings:", result);
        return result;
    }).catch((error) => {
        console.error("Failed to detect field meanings:", error);
    });
    return {};
}

function allFieldsInfo() {
    const inputs = document.querySelectorAll("input, textarea");
    const fieldsDict = new Map();
    const labelMap = fetchInputsLabels(inputs);
    inputs.forEach((input) => {
        const f_id = normalize(field_id(input));
        if (!f_id) {
            return;
        }

        const title = labelMap.get(f_id) || f_id;
        const value = input.value;
        if (fieldsDict.has(f_id)) {
            const existing = fieldsDict.get(f_id);
            existing.push(new Field(title, value, existing[existing.length - 1].rank + 1)); // increase rank for duplicates
            return;
        }
        fieldsDict.set(f_id, [new Field(title, value)]);
    });
    return fieldsDict.values();
}

async function getMemory() {
    return new Promise((resolve) => {
        // Retrieve the memory object from chrome.storage.local
        chrome.storage.local.get(["memory"], (result) => {
            // If there's an error, resolve with an empty object
            if (chrome.runtime.lastError) {
                console.error("Error retrieving memory:", chrome.runtime.lastError);
                resolve({});
                return;
            }
            // Resolve with the retrieved memory or an empty object if it doesn't exist
            resolve(result.memory || {});
        });
    });
}

async function saveMemory(memory) {
    return new Promise((resolve) => {
        // Save the memory object to chrome.storage.local
        // This will overwrite the existing memory with the new one provided
        chrome.storage.local.set({ memory }, () => {
            if (chrome.runtime.lastError) {
                console.error("Error saving memory:", chrome.runtime.lastError);
                resolve(false);
            } else {
                resolve(true);
            }
        });
    });
}

function fillExperiences(experiences) {
    const section = findSectionByKeyword('experience');
    if (!section || !Array.isArray(experiences) || experiences.length === 0) {
        return;
    }

    const addButton = findButtonByKeyword(section, 'add');

    const getExperienceForm = () => {
        return section.querySelector('form') || section;
    };

    const matchExperienceField = (input, experience) => {
        const label = normalize(getLabel(input));

        if (label.includes('company')) return experience.company || '';
        if (label.includes('title') || label.includes('position')) return experience.title || '';
        if (label.includes('start date') || label.includes('from')) return experience.startDate || '';
        if (label.includes('end date') || label.includes('to')) return experience.endDate || '';
        if (label.includes('description') || label.includes('detail')) return experience.description || '';
        return null;
    };

    const fillExperienceForm = (experience) => {
        const form = getExperienceForm();
        const inputs = form.querySelectorAll('input, textarea');

        inputs.forEach((input) => {
            const value = matchExperienceField(input, experience);
            if (!value) return;
            input.focus();
            input.value = value;
            input.dispatchEvent(new Event('input', { bubbles: true }));
            input.dispatchEvent(new Event('change', { bubbles: true }));
        });
    };

    experiences.forEach((experience) => {

        if (!addButton.disabled) {
            addButton.focus();
            addButton.click();
        }

        if (!addButton.disabled) {
            // retry click until button disabled
            let clickCount = 0;
            const maxClicks = 5;
            const clickInterval = setInterval(() => {
                if (addButton.disabled || clickCount >= maxClicks) {
                    clearInterval(clickInterval);
                } else {
                    addButton.click();
                    clickCount++;
                }
            }, 300);
        }

        fillExperienceForm(experience);

        const saveButton = findButtonByKeyword(section, 'save');
        if (saveButton && !saveButton.disabled) {
            saveButton.focus();
            saveButton.click();
            // Check if save button still exists and retry
            setTimeout(() => {
                const retryButton = findButtonByKeyword(section, 'save');
                if (retryButton && !retryButton.disabled && document.contains(retryButton)) {
                    retryButton.focus();
                    retryButton.click();
                }
            }, 300);
        }
    });
}

/**
 * Implementation for filling education details
 * step 1: find the education section, if not found, return.
 * step 2: iterate through educations list, for each detail, click `add` button of the education section, fill the form, and click `save` button. 
 * @param {array} educations The list of education details
 */
function fillEducations(educations) {
    const section = findSectionByKeyword('Education');
    if (!section || !Array.isArray(educations) || educations.length === 0) {
        return;
    }

    const getEducationForm = () => {
        return section.querySelector('form') || section;
    };

    const matchEducationField = (input, education) => {
        const label = normalize(getLabel(input));

        if (label.includes('school')) return education.school || '';
        if (label.includes('field') || label.includes('study')) return education.field || education.fieldOfStudy || '';
        if (label.includes('degree')) return education.degree || '';
        if (label.includes('start date') || label.includes('from')) return education.startDate || '';
        if (label.includes('end date') || label.includes('to')) return education.endDate || '';
        return null;
    };

    const fillEducationForm = (education) => {
        const form = getEducationForm();
        const inputs = form.querySelectorAll('input, textarea');

        inputs.forEach((input) => {
            const value = matchEducationField(input, education);
            if (!value) return;
            input.focus();
            input.value = value;
            input.dispatchEvent(new Event('input', { bubbles: true }));
            input.dispatchEvent(new Event('change', { bubbles: true }));
        });
    };

    const addButton = findButtonByKeyword(section, 'add');

    // Use async function to handle synchronous iteration
    (async () => {
        for (const education of educations) {
            // Click add button to open form for this education
            if (addButton && !addButton.disabled) {
                addButton.focus();
                addButton.click();
            }

            fillEducationForm(education);

            const saveButton = findButtonByKeyword(section, 'save');
            if (saveButton && !saveButton.disabled) {
                saveButton.focus();
                // Wait for the save process to complete
                await new Promise((resolve) => {
                    let clickCount = 0;
                    const maxClicks = 5;
                    const clickInterval = setInterval(() => {
                        if (!addButton.disabled || clickCount >= maxClicks) {
                            clearInterval(clickInterval);
                            resolve();
                        } else {
                            saveButton.click();
                            clickCount++;
                        }
                    }, 300);
                });
            }
        }
    })();
}

function findSectionByKeyword(keyword) {
    const lower = keyword.toLowerCase();
    // First try direct matches on id and class for better performance
    // Using case-insensitive attribute selector
    // This will match any element whose id or class contains the keyword, ignoring case
    const query = `[id*="${keyword}" i], [class*="${keyword}" i]`;
    const direct = document.querySelector(query);
    if (direct) {
        return direct;
    }

    const candidates = Array.from(document.querySelectorAll('section, div, form, article'));
    return candidates.find((el) => {
        return el.innerText && el.innerText.toLowerCase().includes(lower);
    }) || null;
}

function findButtonByKeyword(container, keyword) {
    const lower = keyword.toLowerCase();
    const candidates = Array.from(container.querySelectorAll('button, [role="button"], input[type="button"], input[type="submit"]'));
    return candidates.find((el) => {
        const text = (el.innerText || el.value || el.getAttribute('aria-label') || el.getAttribute('title') || '').toLowerCase();
        const attrs = `${el.id || ''} ${el.className || ''} ${el.getAttribute('name') || ''}`.toLowerCase();
        return text.includes(lower) || attrs.includes(lower);
    }) || null;
}

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
    return text.toLowerCase().replace(/[^a-z0-9]/g, " ");
}

// Example function to handle dropdowns - this is highly dependent on the specific implementation of the dropdown in the target website
function handleDropdown(labelText, value) {
    const elements = document.querySelectorAll("div");

    elements.forEach(el => {
        if (!(el.innerText && el.innerText.includes(labelText))) {
            return;
        }
        el.click();
        setTimeout(() => {
            const options = document.querySelectorAll("div[role='option']");
            options.forEach(opt => {
                if (opt.innerText.includes(value)) {
                    opt.click();
                }
            });
        }, 500);
    });
}

function field_id(input) {
    return input.name || input.id || "";
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

function getLabel(input) {
    // 1. label[for=id]
    if (input.id) {
        const label = document.querySelector(`label[for="${input.id}"]`);
        if (label) return label.innerText;
    }

    // 2. wrapped label
    const parentLabel = input.closest("label");
    if (parentLabel) return parentLabel.innerText;

    // 3. aria-label
    if (input.getAttribute("aria-label")) {
        return input.getAttribute("aria-label");
    }

    // 4. placeholder
    if (input.placeholder) {
        return input.placeholder;
    }

    // 5. nearby text (very important for modern UIs)
    const parent = input.parentElement;
    if (parent) {
        const text = parent.innerText;
        if (text && text.length < 100) return text;
    }

    // 6. fallback
    return input.name || input.id || "";
}

/**
 * Matches a form field to the corresponding profile information
 * @param {object} the html field input component
 * @param {object} profile The user's profile information
 * @returns {string|null} The matching profile information or null if no match is found
 */
function matchField(input, profile) {
    const label = normalize(getLabel(input));

    if (label.includes("email")) return profile.email;
    if (label.includes("first name")) return profile.firstname;
    if (label.includes("last name")) return profile.lastname;
    if (label.includes("full name")) return profile.firstname + " " + profile.lastname;
    if (label.includes("phone")) return profile.phone;
    if (label.includes("location") || label.includes("address")) return profile.location;
    if (label.includes("city")) return profile.city;
    if (label.includes("country")) return profile.country;
    if (label.includes("linkedin")) return profile.linkedin;

    return null;
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
    const detected = await extensionApiFetch(
        `/detect_fields`,
        {
            method: 'POST',
            body: JSON.stringify({ "fields": fields_labels })
        }
    );

    return detected.result;
}

/**
 * 
 * @param {string} userId user ID
 * @returns {object} user information from API
 */
async function fetchInfo(userId) {
    const data = await extensionApiFetch(
        `/user_info?user_id=${encodeURIComponent(userId)}`,
        {
            method: 'GET'
        }
    );

    return {
        profile: data.profile,
        educations: data.educations,
        experiences: data.experiences
    };
}

function fillProfile(profile) {
    const inputs = document.querySelectorAll("input, textarea");

    inputs.forEach((input) => {
        const value = matchField(input, profile);
        if (value) {
            input.focus();
            input.value = value;
            input.style.backgroundColor = "#e6ffed";
            input.dispatchEvent(new Event("input", { bubbles: true }));
            input.dispatchEvent(new Event("change", { bubbles: true }));
        }
    });
}

function fillForm() {

    const userId = "e865c637-aabf-4d80-ab58-423f6904805c";

    fetchInfo(userId).then(({ profile, educations }) => {
        fillProfile(profile);
        fillEducations(educations);
        // fillExperiences(experiences);
    }).catch((error) => {
        console.error("Failed to fetch user info:", error);
    });
}

// const observer = new MutationObserver(() => {
//     fillForm(); // re-run when DOM changes
// });

// observer.observe(document.body, {
//     childList: true,
//     subtree: true
// });

chrome.runtime.onMessage.addListener((msg) => {
    if (msg.action === "FILL_FORM") {
        fillForm();
    }
    if (msg.action === "LIST_FIELDS") {
        allFieldsInfo();
        console.log("All detected fields:", info);
        // alert("Detected fields:\n" + Object.entries(info).map(([k, v]) => `${k}: ${v}`).join("\n"));
    }
});

const profile = {
    firstname: "Peidong",
    familyname: "Xu",
    email: "peidong.xu94@gmail.com",
    phone: "0417909812",
    location: "11 Hillcrest Ave, Epping, NSW 2121",
    city: "Sydney",
    country: "Australia",
    linkedin: "https://www.linkedin.com/in/peidong-xu-9aab4414b/"
};

const educations = [{
    school: "Nanjing University",
    field: "Software Engineering",
    degree: "Master",
    endDate: "10-06-2020",
    startDate: "10-09-2017"
}, {
    school: "East China Institute of Geology",
    field: "Embedded Systems",
    degree: "Bachelor",
    endDate: "10-06-2017",
    startDate: "01-09-2013"
}];

const experiences = [{
    company: "Shopee",
    title: "Software Engineer",
    startDate: "17-06-2020",
    endDate: "18-06-2025",
    description: "I mainly worked on the Chatbot platform, focused on high-traffic backend systems and large-scale data processing, achieving significant performance improvements."
}];

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
        const attrs = `${el.id || ''} ${el.className || ''} ${el.getAttribute('data-testid') || ''} ${el.getAttribute('name') || ''}`.toLowerCase();
        return text.includes(lower) || attrs.includes(lower);
    }) || null;
}


function normalize(text) {
    return text.toLowerCase().replace(/[^a-z0-9]/g, " ");
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
    if (label.includes("last name")) return profile.familyname;
    if (label.includes("full name")) return profile.firstname + " " + profile.familyname;
    if (label.includes("phone")) return profile.phone;
    if (label.includes("location") || label.includes("address")) return profile.location;
    if (label.includes("city")) return profile.city;
    if (label.includes("country")) return profile.country;
    if (label.includes("linkedin")) return profile.linkedin;

    return null;
}

function fillForm() {
    const inputs = document.querySelectorAll("input, textarea");

    inputs.forEach((input) => {
        const value = matchField(input, profile);
        if (value) {
            input.focus();
            input.value = value;
            input.dispatchEvent(new Event("input", { bubbles: true }));
            input.dispatchEvent(new Event("change", { bubbles: true }));
        }
    });

    fillEducations(educations);
    // fillExperiences(experiences);
}

chrome.runtime.onMessage.addListener((msg) => {
    if (msg.action === "FILL_FORM") {
        fillForm();
    }
});
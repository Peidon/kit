const versionElement = document.getElementById("version");
const manifest = chrome.runtime.getManifest();
const brandIcon = document.getElementById("brand-icon");

if (versionElement && manifest?.version) {
    versionElement.textContent = manifest.version;
}

if (brandIcon) {
    const staticSrc = brandIcon.dataset.staticSrc || brandIcon.getAttribute("src");
    const animatedSrc = brandIcon.dataset.animatedSrc;

    const showStaticIcon = () => {
        if (staticSrc) {
            brandIcon.setAttribute("src", staticSrc);
        }
    };

    const playAnimatedIcon = () => {
        if (!animatedSrc) {
            return;
        }

        brandIcon.setAttribute("src", `${animatedSrc}?t=${Date.now()}`);
    };

    brandIcon.addEventListener("mouseenter", playAnimatedIcon);
    brandIcon.addEventListener("mouseleave", showStaticIcon);
    brandIcon.addEventListener("focus", playAnimatedIcon);
    brandIcon.addEventListener("blur", showStaticIcon);
}

async function withLoading(buttonId, action) {
    const button = document.getElementById(buttonId);

    if (!button) {
        return;
    }

    const originalLabel = button.textContent;
    const loadingElement = document.createElement("span");
    loadingElement.className = "loading";
    loadingElement.textContent = "...";

    button.disabled = true;
    button.appendChild(loadingElement);
    button.blur();

    try {
        await action();
    } finally {
        loadingElement.remove();
        button.disabled = false;
        button.textContent = originalLabel;
    }
}

async function sendAction(actionName) {
    const [tab] = await chrome.tabs.query({ active: true, currentWindow: true });

    if (!tab?.id) {
        return;
    }

    await chrome.tabs.sendMessage(tab.id, { action: actionName });
}

document.getElementById("fill")?.addEventListener("click", async () => {
    await withLoading("fill", () => sendAction("FILL_FORM"));
});

document.getElementById("learn")?.addEventListener("click", async () => {
    await withLoading("learn", () => sendAction("LEARN"));
});

document.getElementById("review")?.addEventListener("click", async () => {
    await withLoading("review", () => sendAction("REVIEW"));
});

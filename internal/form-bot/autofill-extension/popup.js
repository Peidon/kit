const versionElement = document.getElementById("version");
const manifest = chrome.runtime.getManifest();

if (versionElement && manifest?.version) {
    versionElement.textContent = manifest.version;
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

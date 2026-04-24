document.getElementById("fill").addEventListener("click", async () => {
    const fillButton = document.getElementById("fill");
    const loadingElement = document.createElement("span");
    loadingElement.className = "loading";
    loadingElement.textContent = "ing...";
    fillButton.disabled = true;
    fillButton.appendChild(loadingElement);
    fillButton.blur();

    const [tab] = await chrome.tabs.query({ active: true, currentWindow: true });

    chrome.tabs.sendMessage(tab.id, { action: "FILL_FORM" }, () => {
        fillButton.removeChild(loadingElement);
        fillButton.disabled = false;
    });
});

document.getElementById("learn").addEventListener("click", async () => {
    const listButton = document.getElementById("learn");
    const loadingElement = document.createElement("span");
    loadingElement.className = "loading";
    loadingElement.textContent = "ing...";
    listButton.disabled = true;
    listButton.appendChild(loadingElement);
    listButton.blur();

    const [tab] = await chrome.tabs.query({ active: true, currentWindow: true });

    chrome.tabs.sendMessage(tab.id, { action: "LEARN" }, () => {
        listButton.removeChild(loadingElement);
        listButton.disabled = false;
    });
});
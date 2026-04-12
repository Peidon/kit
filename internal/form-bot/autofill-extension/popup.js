document.getElementById("fill").addEventListener("click", async () => {
    const fillButton = document.getElementById("fill");
    const loadingElement = document.createElement("span");
    loadingElement.className = "loading";
    loadingElement.textContent = " Loading...";
    fillButton.disabled = true;
    fillButton.appendChild(loadingElement);
    fillButton.blur();

    const [tab] = await chrome.tabs.query({ active: true, currentWindow: true });

    chrome.tabs.sendMessage(tab.id, { action: "FILL_FORM" }, () => {
        fillButton.removeChild(loadingElement);
        fillButton.disabled = false;
    });
});
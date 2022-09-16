chrome.runtime.onInstalled.addListener(() => {
  chrome.contextMenus.create({
    "id": "wordbook",
    "title": chrome.i18n.getMessage("extTabText"),
    "contexts": ["selection"]
  });
});

chrome.contextMenus.onClicked.addListener((data, tab) => {
  chrome.tabs.sendMessage(tab.id, {name: "wordbook"}, () => {
  });
})

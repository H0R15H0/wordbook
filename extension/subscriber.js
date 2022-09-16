const DIALOG_OFFSET = 5

chrome.runtime.onMessage.addListener((request) => {
  switch (request.name)  {
    case "wordbook":
      const selection = window.getSelection()
      console.log(selection.toString())
      break;
  
    default:
      break;
  }
});

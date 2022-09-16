const DIALOG_OFFSET = 5
const url = ""

chrome.runtime.onMessage.addListener((request) => {
  switch (request.name)  {
    case "wordbook":
      const selectedText = window.getSelection().toString()
      fetch(url, {
        method: 'POST',
        mode: 'no-cors',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          text: selectedText,
          source_url: window.location.toString() + "#:~:text=" + selectedText
        }),
      })
      .then(() => {
        console.log('wordbook: success!');
      })
      .catch((error) => {
        alert(error)
      });
      break;
  
    default:
      break;
  }
});

const DIALOG_OFFSET = 5
const url = ""

chrome.runtime.onMessage.addListener((request) => {
  switch (request.name)  {
    case "wordbook":
      const selection = window.getSelection()
      fetch(url, {
        method: 'POST',
        mode: 'no-cors',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          text: selection.toString(),
          source_url: window.location.toString()
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

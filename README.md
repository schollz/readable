# readability-bookmarklet
A simple link you can drag into your bookmarks that will make any page readable


```html
<a onclick="alert('Drag this link to your browser\'s toolbar to create the bookmarklet.'); return false;" href="javascript: (function () {

var xhr = new XMLHttpRequest();
xhr.open("POST", "https://readable.schollz.com/", true);
xhr.setRequestHeader("Content-type", "application/json");
xhr.onreadystatechange = function () {
    if (xhr.readyState === 4 && xhr.status === 200) {
        var json = JSON.parse(xhr.responseText);
        document.write(json.html);
        document.close();
    }
};
var data = JSON.stringify({"url": window.location.href});
xhr.send(data);

 })()">Make readable</a>
```

<a onclick="alert('Drag this link to your browser\'s toolbar to create the bookmarklet.'); return false;" href="javascript: !function(){var e=new XMLHttpRequest;e.open("POST","https://readable.schollz.com/",!0),e.setRequestHeader("Content-type","application/json"),e.onreadystatechange=function(){if(4===e.readyState&&200===e.status){var t=JSON.parse(e.responseText);document.write(t.html),document.close()}};var t=JSON.stringify({url:window.location.href});e.send(t)}();">Make readable</a>

## Server

```
$ export MERCURY_API_KEY=XX
$ go build && ./readability-bookmarklet
```
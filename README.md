# readability-bookmarklet
A simple link you can drag into your bookmarks that will make any page readable

Here is an example [from dotsies.org](http://dotsies.org/):

```html
<a onclick="alert('Drag this link to your browser\'s toolbar to create the bookmarklet.'); return false;" href="javascript:
            (function(){
            var existing = false;

            function remove_if_there(tagName) {
              elems = document.getElementsByTagName(tagName);
                for(var i=0; i<elems.length; ++i){
                  if (elems[i].className == 'dotsies-toggle') {
                  existing = true;
                  elems[i].parentNode.removeChild(elems[i]);
                }
              }
            }
            remove_if_there('style');
            remove_if_there('link');

            if(!existing){
              var l = document.createElement('link');
              document.getElementsByTagName('head')[0].appendChild(l);
              l.setAttribute('href', 'http://dotsies.org/dotsies.css');
              l.setAttribute('type', 'text/css');
              l.setAttribute('rel', 'stylesheet');
              l.setAttribute('class', 'dotsies-toggle');
              var s = document.createElement('style');
              s.setAttribute('class', 'dotsies-toggle');
              document.getElementsByTagName('head')[0].appendChild(s);
              s.innerHTML = ('body, p, div, span, a, li, textarea, input, font, blockquote {font-family: Dotsies !important}');
            }

            })()">dotsies</a>
```

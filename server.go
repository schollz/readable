package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(os.ExpandEnv("$MERCURY_API_KEY"))
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
<html><head></head><body>
<p>Drag this link to your browser's toolbar to create the bookmarklet:</p>
<a onclick="alert('Drag this link to your browser\'s toolbar to create the bookmarklet.'); return false;" href='javascript: !function(){var e=new XMLHttpRequest;e.open("POST","https://readable.schollz.com/",!0),e.setRequestHeader("Content-type","application/json"),e.onreadystatechange=function(){if(4===e.readyState&&200===e.status){var t=JSON.parse(e.responseText);document.write(t.html),document.close()}};var t=JSON.stringify({url:window.location.href});e.send(t)}();'>Make readable</a>
		
		</body></html>`))
	})
	router.POST("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		type Payload struct {
			URL string `json:"url" binding:"required"`
		}
		var json Payload
		if c.BindJSON(&json) == nil {
			c.JSON(http.StatusOK, gin.H{"html": generateHTML(json.URL)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Something happened"})
		}

	})
	router.OPTIONS("/*cors", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Returns nothing :)
	})

	router.Run(":8078")
}

func generateHTML(url string) (html string) {
	req, err := http.NewRequest("GET", "https://mercury.postlight.com/parser?url="+url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Api-Key", os.ExpandEnv("$MERCURY_API_KEY"))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	type Response struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	var target Response
	json.NewDecoder(resp.Body).Decode(&target)
	if len(target.Content) == 0 {
		target.Content = "Something went wrong, check API key"
	}
	html = `<html><head><style>b,em,i,strong{color:#222223}figure figcaption a,h1 a,h2 a,h3 a{text-decoration:none}blockquote,blockquote a{color:#666664}>h1,li ol,li ul,p{margin-top:0}a,code,tt{word-wrap:break-word}code,pre,tt{background-color:#f9f9f7}figure img,table{box-sizing:border-box;max-width:100%}figure img,hr,iframe{display:block}figure img,iframe,img,table{max-width:100%}.unstyled,table pre{margin:0;padding:0}pre code,pre tt,table pre{background:0 0;border:none}body{margin:1em auto;max-width:40em;padding:0 .62em}h1,h2,h3{line-height:1.2}@media print{body{max-width:none}}body{font:400 18px/1.62 -apple-system,BlinkMacSystemFont,"Segoe UI","Droid Sans","Helvetica Neue","PingFang SC","Hiragino Sans GB","Droid Sans Fallback","Microsoft YaHei",sans-serif;color:#444443}body ::-moz-selection{background-color:rgba(0,0,0,.2)}body ::selection{background-color:rgba(0,0,0,.2)}h1,h2,h3,h4,h5,h6{font-family:-apple-system,BlinkMacSystemFont,"Segoe UI","Droid Sans","Helvetica Neue","PingFang SC","Hiragino Sans GB","Droid Sans Fallback","Microsoft YaHei",sans-serif;color:#222223}h1{font-size:1.8em;margin:.67em 0}>h1{font-size:2em}h2{font-size:1.5em;margin:.83em 0}h3{font-size:1.17em;margin:1em 0}h4,h5,h6{font-size:1em;margin:1.6em 0 1em}h6{font-weight:500}p{margin-bottom:1.24em}a{color:#111;-webkit-text-decoration-color:rgba(0,0,0,.4);text-decoration-color:rgba(0,0,0,.4)}a:hover{color:#555;-webkit-text-decoration-color:rgba(0,0,0,.6);text-decoration-color:rgba(0,0,0,.6)}b,strong{font-weight:700}em,i{font-style:italic}img{height:auto;margin:.2em 0}a img{border:none}figure{position:relative;clear:both;outline:0;margin:10px 0 30px;padding:0;min-height:100px}figure img{margin:auto auto 4px}figure figcaption{position:relative;width:100%;text-align:center;left:0;margin-top:10px;font-weight:400;font-size:14px;color:#666665}code,pre,table,tt{font-size:.96em}figure figcaption a{color:#666665}hr{width:14%;margin:40px auto 34px;border:0;border-top:3px solid #dededc}blockquote{margin:0 0 1.64em;border-left:3px solid #dadada;padding-left:12px}ol,ul{margin:0 0 24px 6px;padding-left:16px}ul{list-style-type:square}ol{list-style-type:decimal}li{margin-bottom:.2em}li ol,li ul{margin-bottom:0;margin-left:14px}li ul{list-style-type:disc}li ul ul{list-style-type:circle}li p{margin:.4em 0 .6em}.unstyled{list-style-type:none}code,tt{color:grey;padding:1px 2px;border:1px solid #eee;border-radius:3px;font-family:Menlo,Monaco,Consolas,"Courier New",monospace}pre{margin:1.64em 0;padding:7px 7px 7px 10px;border:none;border-left:3px solid #dadada;overflow:auto;line-height:1.5;font-family:Menlo,Monaco,Consolas,"Courier New",monospace;color:#4c4c4c}pre code,pre tt{color:#4c4c4c;padding:0}table{width:100%;border-collapse:collapse;border-spacing:0;margin-bottom:1.5em}td,th{text-align:left;padding:4px 8px 4px 10px;border:1px solid #dadada}td{vertical-align:top}tr:nth-child(even){background-color:#efefee}iframe{margin-bottom:30px}figure iframe{margin:auto}@media (min-width:1100px){blockquote{margin-left:-24px;padding-left:20px;border-width:4px}blockquote blockquote{margin-left:0}}</style></head><body><h1>` + target.Title + `</h1>` + target.Content + `</body></html>`
	return
}

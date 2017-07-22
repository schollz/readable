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
	html = `<html><head><style>img {
   height: auto;
   width:100%;
}body{margin:1em auto;max-width:40em;padding:0 .62em;font:1.2em/1.62 sans-serif;}h1,h2,h3{line-height:1.2;}@media print{body{max-width:none}}</style></head><body>` + target.Content + `</body></html>`
	return
}

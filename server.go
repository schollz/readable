package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

var apiKey, port, externalServer string
var css = `b,em,i,strong{color:#222223}figure figcaption a,h1 a,h2 a,h3 a{text-decoration:none}blockquote,blockquote a{color:#666664}>h1,li ol,li ul,p{margin-top:0}a,code,tt{word-wrap:break-word}code,pre,tt{background-color:#f9f9f7}figure img,table{box-sizing:border-box;max-width:100%}figure img,hr,iframe{display:block}figure img,iframe,img,table{max-width:100%}.unstyled,table pre{margin:0;padding:0}pre code,pre tt,table pre{background:0 0;border:none}body{margin:1em auto;max-width:40em;padding:0 .62em}h1,h2,h3{line-height:1.2}@media print{body{max-width:none}}body{font:400 18px/1.62 -apple-system,BlinkMacSystemFont,"Segoe UI","Droid Sans","Helvetica Neue","PingFang SC","Hiragino Sans GB","Droid Sans Fallback","Microsoft YaHei",sans-serif;color:#444443}body ::-moz-selection{background-color:rgba(0,0,0,.2)}body ::selection{background-color:rgba(0,0,0,.2)}h1,h2,h3,h4,h5,h6{font-family:-apple-system,BlinkMacSystemFont,"Segoe UI","Droid Sans","Helvetica Neue","PingFang SC","Hiragino Sans GB","Droid Sans Fallback","Microsoft YaHei",sans-serif;color:#222223}h1{font-size:1.8em;margin:.67em 0}>h1{font-size:2em}h2{font-size:1.5em;margin:.83em 0}h3{font-size:1.17em;margin:1em 0}h4,h5,h6{font-size:1em;margin:1.6em 0 1em}h6{font-weight:500}p{margin-bottom:1.24em}a{color:#111;-webkit-text-decoration-color:rgba(0,0,0,.4);text-decoration-color:rgba(0,0,0,.4)}a:hover{color:#555;-webkit-text-decoration-color:rgba(0,0,0,.6);text-decoration-color:rgba(0,0,0,.6)}b,strong{font-weight:700}em,i{font-style:italic}img{height:auto;margin:.2em 0}a img{border:none}figure{position:relative;clear:both;outline:0;margin:10px 0 30px;padding:0;min-height:100px}figure img{margin:auto auto 4px}figure figcaption{position:relative;width:100%;text-align:center;left:0;margin-top:10px;font-weight:400;font-size:14px;color:#666665}code,pre,table,tt{font-size:.96em}figure figcaption a{color:#666665}hr{width:14%;margin:40px auto 34px;border:0;border-top:3px solid #dededc}blockquote{margin:0 0 1.64em;border-left:3px solid #dadada;padding-left:12px}ol,ul{margin:0 0 24px 6px;padding-left:16px}ul{list-style-type:square}ol{list-style-type:decimal}li{margin-bottom:.2em}li ol,li ul{margin-bottom:0;margin-left:14px}li ul{list-style-type:disc}li ul ul{list-style-type:circle}li p{margin:.4em 0 .6em}.unstyled{list-style-type:none}code,tt{color:grey;padding:1px 2px;border:1px solid #eee;border-radius:3px;font-family:Menlo,Monaco,Consolas,"Courier New",monospace}pre{margin:1.64em 0;padding:7px 7px 7px 10px;border:none;border-left:3px solid #dadada;overflow:auto;line-height:1.5;font-family:Menlo,Monaco,Consolas,"Courier New",monospace;color:#4c4c4c}pre code,pre tt{color:#4c4c4c;padding:0}table{width:100%;border-collapse:collapse;border-spacing:0;margin-bottom:1.5em}td,th{text-align:left;padding:4px 8px 4px 10px;border:1px solid #dadada}td{vertical-align:top}tr:nth-child(even){background-color:#efefee}iframe{margin-bottom:30px}figure iframe{margin:auto}@media (min-width:1100px){blockquote{margin-left:-24px;padding-left:20px;border-width:4px}blockquote blockquote{margin-left:0}}`

type Response struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func init() {
	os.MkdirAll(filepath.Join(".", "cache"), 0644)
}

func main() {
	flag.StringVar(&apiKey, "key", "", "Mercury API key (default: use docker)")
	flag.StringVar(&port, "port", "8078", "internal port to use")
	flag.StringVar(&externalServer, "server", "", "external server name to use (optional)")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(location.Default())

	router.GET("/*url", func(c *gin.Context) {
		u := location.Get(c)
		serverName := u.Scheme + "://" + u.Host
		if externalServer != "" {
			serverName = externalServer
		}
		url := c.Param("url")

		if url == "/" {
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
<html><head><title>Make readable</title><style>`+css+`</style></head><body>
<center>
<p>Drag this link to your browser's toolbar to create the bookmarklet:</p>
<strong>
<a onclick="alert('Drag this link to your browser\'s toolbar to create the bookmarklet.'); return false;" href='javascript: !function(){var url="`+serverName+`/" + window.location.href; window.location.href=url;}();'>Make readable</a></strong>
<p>When you are on a news site, click the link to make it readable.</p>
<p>For more information, or to run your own server, check out <a href="https://github.com/schollz/readability-bookmarklet">the source</a>.</p>
</center>
</body></html>`))
			return
		} else {
			c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(generateHTML(url[1:])))
		}

	})

	fmt.Println("Running! Open http://localhost:" + port + " to get your bookmarklet.")
	router.Run(":" + port)
}

func generateHTMLMercury(url string) (target Response) {
	req, err := http.NewRequest("GET", "https://mercury.postlight.com/parser?url="+url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("X-Api-Key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&target)
	if len(target.Content) == 0 {
		target.Content = "Something went wrong, check API key"
	}
	return
}

func generateHTMLSelf(url string) (target Response) {
	dir, _ := filepath.Abs(filepath.Join(".", "cache"))
	fmt.Println(dir)
	log.Println("Attempting docker")
	log.Println("docker", "run", "-v", dir+":/data", "-t", "readable", url)
	_, err := exec.Command("docker", "run", "-v", dir+":/data", "-t", "readable", url).Output()
	if err != nil {
		log.Print(err)
	}
	data, _ := ioutil.ReadFile(filepath.Join(".", "cache", GetMD5Hash(url)+".json"))
	json.Unmarshal(data, &target)
	return
}

func generateHTML(url string) (html string) {
	log.Println(url)
	cacheFile := filepath.Join(".", "cache", GetMD5Hash(url)+".json")
	log.Println(cacheFile)

	var target Response

	data, err := ioutil.ReadFile(cacheFile)
	if err == nil {
		json.Unmarshal(data, &target)
	} else {
		if apiKey != "" {
			log.Println("Using Mercury API")
			target = generateHTMLMercury(url)
		} else {
			log.Println("Using self-hosted")
			target = generateHTMLSelf(url)
		}
		jbytes, err := json.Marshal(target)
		if err != nil {
			log.Print(err)
		}
		err = ioutil.WriteFile(cacheFile, jbytes, 0644)
		if err != nil {
			log.Print(err)
		}
	}
	html = `<html><head><title>` + target.Title + `</title><style>` + css + `</style></head><body><h1>` + target.Title + `</h1>` + target.Content + `</body></html>`
	return html
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

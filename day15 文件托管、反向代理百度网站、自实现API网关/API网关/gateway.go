package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

/*
域名对域名的一对一代理，没有考虑path聚合不同域名的方式
*/
func gateway(targetUrl string, ssl bool) gin.HandlerFunc {
	targetURL, err := url.Parse(targetUrl)
	if err != nil {
		log.Fatalln(err)
	}
	host := targetURL.Host
	if ssl {
		targetUrl = "https://" + host
	} else {
		targetUrl = "http://" + host
	}
	replaceUrl := func(body io.Reader, contentType string, contentLength int64) (io.Reader, int64) {
		if strings.HasPrefix(contentType, "text/") {
			build := &strings.Builder{}
			io.Copy(build, body)
			html := strings.ReplaceAll(build.String(), "http://"+host, "")
			html = strings.ReplaceAll(html, "https://"+host, "")
			html = strings.ReplaceAll(html, "//"+host, "")
			return strings.NewReader(html), int64(len([]byte(html)))
		} else {
			return body, contentLength
		}

	}

	return func(c *gin.Context) {
		requestUrl := c.Request.URL.RequestURI()
		response, err := http.Get(targetUrl + requestUrl)
		if err != nil {
			c.AbortWithError(http.StatusServiceUnavailable, err)
			return
		}
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		reader, contentLength := replaceUrl(response.Body, contentType, contentLength)
		c.DataFromReader(response.StatusCode, contentLength, contentType, reader, map[string]string{})
	}

}

func main() {
	router := gin.Default()
	router.GET("/*filepath", gateway("http://www.baidu.com", false))
	router.Run(":8080")
}

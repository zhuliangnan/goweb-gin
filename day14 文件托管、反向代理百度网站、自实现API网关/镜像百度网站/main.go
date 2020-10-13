package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

func main() {
	router := gin.Default()

	router.GET("/*filepath", func(c *gin.Context) {
		requestUrl := c.Request.URL.RequestURI()
		response, err := http.Get("http://www.baidu.com" + requestUrl)
		if err != nil {
			c.AbortWithError(http.StatusServiceUnavailable, err)
			return
		}
		//reader := response.Body
		reader := replaceUrl(response)
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		c.DataFromReader(response.StatusCode, contentLength, contentType, reader, map[string]string{})
	})
	router.Run(":8080")
}
func replaceUrl(resp *http.Response) io.Reader {
	contentType := resp.Header.Get("Content-Type")
	fmt.Println(contentType)
	if strings.HasPrefix(contentType, "text/") {
		build := &strings.Builder{}
		io.Copy(build, resp.Body)
		html := strings.ReplaceAll(build.String(), "http://www.baidu.com", "")
		html = strings.ReplaceAll(html, "https://www.baidu.com", "")
		html = strings.ReplaceAll(html, "//www.baidu.com", "")
		return strings.NewReader(html)
	} else {
		return resp.Body
	}

}

package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func proxy(c *gin.Context) {
	remote, err := url.Parse("https://webhook.site/d98a8160-50a8-4f26-b68a-8e53989f3fd8")
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.Director = func(req *http.Request) {
		req.Header = c.Request.Header
		req.Host = remote.Host
		req.URL.Scheme = remote.Scheme
		req.URL.Host = remote.Host
		req.URL.Path = remote.Path + c.Param("proxyPath")
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {

	r := gin.Default()

	r.Any("/*proxyPath", proxy)

	r.Run(":8080")
}

// func main() {
// 	eng := gin.Default()
// 	eng.GET("/", ping)
// 	eng.Run(":8999")
// 	fmt.Println("HEllo world")
// }
// func ping(g *gin.Context) {
// 	g.JSON(http.StatusAccepted, gin.H{"Status": "hello"})

// }

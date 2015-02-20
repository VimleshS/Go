package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.String(200, "msg")
	//})
	//r.Run(":8080")

	v1 := r.Group("/v1")
	{
		v1.GET("/login", postHandleFunc)
		v1.GET("/submit", postHandleFunc)
		v1.GET("/read", postHandleFunc)
	}
	r.Run(":8080")

}

func postHandleFunc(c *gin.Context) {
	c.String(200, "Post Executed")
}

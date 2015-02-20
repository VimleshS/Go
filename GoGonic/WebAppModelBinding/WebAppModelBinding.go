package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	//"github.com/gin-gonic/gin/binding"
	"fmt"
)

// Binding from JSON
type LoginJSON struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Binding from form values
type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	/**/

	// Example for binding JSON ({"user": "manu", "password": "123"})
	r.POST("/login", func(c *gin.Context) {
		var json LoginJSON
		c.Bind(&json)

		fmt.Println(json)
		if json.User == "manu" && json.Password == "123" {
			c.JSON(200, gin.H{"status": "You r logged in"})
			response, _ := http.Get("http://www.google.co.in/robots.txt")
			body, _ := ioutil.ReadAll(response.Body)
			c.JSON(200, gin.H{"Data": string(body)})
		} else {
			c.JSON(401, gin.H{"status": "You r unauthorized"})
		}
	})

	// Example for binding a HTLM form (user=manu&password=123)
	//r.POST("/login", func(c *gin.Context) {
	//	var form LoginForm

	//	c.BindWith(&form, binding.Form) // You can also specify which binder to use. We support binding.Form, binding.JSON and binding.XML.
	//	if form.User == "manu" && form.Password == "123" {
	//		c.JSON(200, gin.H{"status": "you are logged in"})
	//	} else {
	//		c.JSON(401, gin.H{"status": "unauthorized"})
	//	}
	//})

	r.Run(":8080")
}

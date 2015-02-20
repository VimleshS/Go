package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

//http://stackoverflow.com/questions/21815520/golang-gorilla-session
//http://golang.org/ref/spec#Type_assertions

func main() {
	r := gin.Default()

	var store = sessions.NewCookieStore([]byte("something-very-secret"))

	r.GET("/session", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session-name")

		_, ok := session.Values["Counter"]
		if !ok {
			session.Values["Counter"] = 1
			session.Options = &sessions.Options{
				Path:     "/session",
				MaxAge:   10,
				HttpOnly: true}

		} else {
			temp := session.Values["Counter"]
			//http://golang.org/ref/spec#Type_assertions
			temp = temp.(int) + 1
			session.Values["Counter"] = temp
		}

		err := session.Save(c.Request, c.Writer)
		if err != nil {
			c.String(400, fmt.Sprintf("Error : %v", err))
			return
		}
		c.String(200, fmt.Sprintf("Counter variable is %d", session.Values["Counter"]))

	})
	r.GET("/delete", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session-name")
		session.Options = &sessions.Options{
			MaxAge: -1,
		}

		session.Save(c.Request, c.Writer)
	})
	r.Run(":8080")
}

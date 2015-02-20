/*
https://github.com/gin-gonic/gin/issues/27
https://gist.github.com/alexandernyquist/1ee331db9f5bf0d577d6
*/
package main

import (
	"fmt"
	//	"github.com/alexandernyquist/gin"
	"github.com/gin-gonic/gin"
	"html/template"
)

type ViewData struct {
	UserId string
}

func main() {
	r := gin.Default()

	r.GET("", func(g *gin.Context) {
		//r.HTMLTemplates = template.Must(template.ParseFiles("layout.html", "home.html"))
		temp := template.Must(template.ParseFiles("layout.html", "home.html"))
		r.SetHTMLTemplate(temp)
		fmt.Println(temp)
		//r.LoadHTMLFiles("layout.html", "home.html")
		g.HTML(200, "base", nil)
	})

	r.GET("/user/:id", func(g *gin.Context) {
		//r.HTMLTemplates = template.Must(template.ParseFiles("layout.html", "view.html"))
		temp := template.Must(template.ParseFiles("layout.html", "view.html"))
		r.SetHTMLTemplate(temp)
		//r.LoadHTMLFiles("layout.html", "view.html")
		g.HTML(200, "base", ViewData{g.Params.ByName("id")})
	})

	r.Run(":8082")
}

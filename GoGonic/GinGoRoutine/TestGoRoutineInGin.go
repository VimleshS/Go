package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	r := gin.Default()

	r.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		c_cp := c.Copy()
		//		str := make(chan string)
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note than you are using the copied context "c_cp", IMPORTANT
			//			str <- "Done! in path " + c_cp.Request.URL.Path
			log.Println("Done! in path " + c_cp.Request.URL.Path)

		}()

		//		c_cp.String(200, <-str)
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		log.Println("Done! in path " + c.Request.URL.Path)
		c.String(200, "long_sync OK")
	})

	// Listen and server on 0.0.0.0:8080
	r.Run(":8080")
}

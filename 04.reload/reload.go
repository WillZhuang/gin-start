package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := endless.ListenAndServe("127.0.0.1:4242", r)
	if err != nil {
		log.Println(err)
	}
	log.Println("Server on 4242 stopped")

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

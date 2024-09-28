package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	r := gin.Default()

	r.GET("/hello/:name", func(c *gin.Context) {

		name := c.Param("name")

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello %s", name),
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

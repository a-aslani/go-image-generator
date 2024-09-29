package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {

		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var fileSize int64 = 5

		if file.Size > fileSize<<20 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Errorf("file size exceeds maximum limit %d MB", fileSize).Error(),
			})
			return
		}

		if err = c.SaveUploadedFile(file, fmt.Sprintf("./tmp/%s", file.Filename)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message": fmt.Sprintf("File has been uploaded, %s", file.Filename),
		})
	})

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}

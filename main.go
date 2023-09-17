package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/upload", func(c *gin.Context) {
		// Get the file
		file, err := c.FormFile("image")

		if err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": "Failed to upload image",
			})
			return
		}

		filePath := "assets/uploads/" + file.Filename

		// Upload the file to specific folder.
		err = c.SaveUploadedFile(file, filePath)

		if err != nil {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"error": "Failed to upload image",
			})
			return
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"image": "/" + filePath,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

package main

import (

"github.com/gin-gonic/gin"
"net/http"
"fmt"
"log"
)

func main(){

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
	   c.String(http.StatusOK, "Hello World")
	   c.String(http.StatusOK, "fuck")
	})
	router.POST("/upload", uploadHandler)
	router.Run(":8000")
}

func uploadHandler(c *gin.Context) {
	fmt.Println("come to uploadHandler")
	form, _ := c.MultipartForm()
	files := form.File["upload-key"]

	for _, file := range files {
		log.Println(file.Filename)
		c.SaveUploadedFile(file,file.Filename)
		c.String(http.StatusOK, "Uploaded success")
	}
	c.String(http.StatusOK, "Uploaded finish")
}
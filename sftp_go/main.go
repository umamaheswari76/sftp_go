package main

import (
	"net/http"
	"sftp_go/config"
	"sftp_go/services"

	"github.com/gin-gonic/gin"
)

var remotePath = "sftp/testint"

func main() {
	client := config.SftpClient()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"message": "hello"})
	})

	router.GET("/list", func(c *gin.Context) {
		
	})

	//creating a remote file and copying content from local to it
	//services.Put(client)

	//get remote file into local directory
	// services.Get(client)

	services.List(client, remotePath)
	router.Run(":8080")
}

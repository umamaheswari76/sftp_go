package main

import (
	"sftp_go/config"
	"sftp_go/services"
)

func main() {
	// router := gin.Default()

	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK,gin.H{"message": "hello"})
	// })

	// router.GET("/list", controllers.ListControllers)

	client := config.SftpClient()
	// creating a remote file and copying content from local to it
	services.Put(client)

	// get remote file into local directory
	services.Get(client)

	services.List(client, remotePath)
	// router.Run(":8080")
}

package controllers

import (
	"sftp_go/config"
	"sftp_go/constants"
	"sftp_go/services"

	"github.com/gin-gonic/gin"
)

func ListControllers(c *gin.Context){
	services.List(c, config.SftpClient(),constants.RemotePath)
}

// func List(client *sftp.Client, remotePath string){

// 	files, err := client.ReadDir(remotePath)
// 	if err != nil{
// 		log.Fatal("Unable to list remote dir: ",err)
// 	}

// 	for _,f := range files{
// 		var name string
// 		name = f.Name()

// 		if f.IsDir(){
// 			name = name + "/"
// 		}
// 		fmt.Println(name)
// 	}
// }


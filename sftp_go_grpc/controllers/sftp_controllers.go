package controllers

import (
	"sftp_go_grpc/config"
	"sftp_go_grpc/constants"
	"sftp_go_grpc/services"

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


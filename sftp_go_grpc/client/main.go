package main

import (
	"context"
	"fmt"
	"log"

	pb "sftp_go_grpc/proto"

	"google.golang.org/grpc"
)

var RPath = "sftp/testint"

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewSftpServiceClient(conn)
	res, err := client.List(context.Background(), &pb.RemotePathRequest{
		RemotePath: RPath,
	})
	if err != nil {
		fmt.Println("Cant get the files")
	}

	fmt.Println(res.Files)
}



	// r := gin.Default()
	// fmt.Println("0")

	// r.GET("/list", func(ctx *gin.Context){
	// 	fmt.Println("1")
	// 	// remotepath := ctx.Param("remotepath")
	// 	res, err := client.List(ctx, &pb.RemotePathRequest{
	// 		RemotePath: RPath,
	// 	})
	// 	fmt.Println("2")
	// 	if err != nil {
	// 		ctx.JSON(http.StatusNotFound, gin.H{
	// 			"message": err.Error(),
	// 		})
	// 		return}
	// 	ctx.JSON(http.StatusOK,gin.H{"Files":res.Files})
	// })

	// response, err := client.List(context.Background(), &pb.RemotePathRequest{
	// 	RemotePath: RPath,
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to call List: %v", err)
	// }
	// fmt.Printf("Response: %v\n", response.Files)


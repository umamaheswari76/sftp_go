package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sftp_go_grpc/config"
	pb "sftp_go_grpc/proto"

	"google.golang.org/grpc"
)

//getting the sftp connection
//create grpc server
//implement rpc methods

var sftpClient = config.SftpClient()

type server struct {
	pb.UnimplementedSftpServiceServer
}

func (s *server) List(ctx context.Context,req *pb.RemotePathRequest) (*pb.Response, error) {
	files, err := sftpClient.ReadDir(req.RemotePath)
	if err != nil {
		log.Fatal("Unable to list remote dir: ", err)
	}

	var files_lst []string
	var response *pb.Response

	for _, f := range files {
		var name string
		name = f.Name()

		if f.IsDir() {
			name = name + "/"
		}
		files_lst = append(files_lst, name)
	}
	response.Files = files_lst
	return response,nil
}

func main() {
	fmt.Println("gRPC server listening on")
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSftpServiceServer(s, &server{})
	fmt.Println("Server listening on :8000")

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}

}

package main

import (

	"fmt"
	pb "grpc_rw/proto"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

//getting the sftp connection
//create grpc server
//implement rpc methods

type server struct {
	pb.UnimplementedGrpcRwServiceServer
}


func (s *server) Upload(stream pb.GrpcRwService_UploadServer) error {
	OutputFile, _ := os.Create("/home/vasenth/Documents/umamaheswari/sftp_go/test3.txt")
	var (
		fileName string
		fileSize uint32
	)
	fileSize = 0
	for {
		fmt.Println("1")
		req, err := stream.Recv()
		if err != nil {
			fmt.Println(err)
		}
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("err: ", err)
		}

		fileName = req.GetFileName()
		chunk := req.GetChunk()
		fileSize += uint32(len(chunk))
		fmt.Println("Received chunk with size: ", fileSize)
		_, err1 := OutputFile.Write(chunk)
		if err1 != nil {
			log.Fatal("error writing chunk: ", err1)
		}
	}

	fmt.Println("Received file size ", fileSize)
	return stream.SendAndClose(&pb.FileUploadResponse{
		FileName: fileName,
		Size:     fileSize,
	})
}

func main() {
	fmt.Println("gRPC server listening on")
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGrpcRwServiceServer(s, &server{})
	fmt.Println("Server listening on: 8000")

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}

package main

import (
	"context"
	pb "grpc_rw/proto"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	// defer conn.Close()
	client := pb.NewGrpcRwServiceClient(conn)

	// ctx, _ := context.WithCancel(context.Background())

	// err1 := upload(ctx, client)
	// if err1 != nil {
	// 	log.Fatal("error client service: ", err1)
	// }
	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatal("error uploading to stream: ", err)
	}

	file, err := os.Open("/home/vasenth/Documents/umamaheswari/sftp_go/grpc_rw/test.txt")
	if err != nil {
		log.Fatal("error opening file: ", err)
	}

	buffer := make([]byte, 14) 
	batchNumber := 1
	for {
		num, err := file.Read(buffer)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("error reading file: ", err)
		}
		chunk := buffer[:num]
		if err := stream.Send(&pb.FileUploadRequest{FileName: "test.txt", Chunk: chunk}); err != nil {
			log.Fatal("error sending file: ", err)
		}
		log.Printf("Sent - batch #%v - size - %v\n", batchNumber, len(chunk))
		batchNumber += 1
	}

	
}

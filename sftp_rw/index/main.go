package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"sftp_rw/config"
)

func main() {
	client, _ := config.SftpConnection()

	remotefile, err := client.Open("sftp/testint/test.txt")
	if err != nil {
		log.Fatal("Failed to open remote file ", err)
	}
	defer remotefile.Close()

	localPath, err1 := os.Create("F:/desktop/Go/sftp_go/test_files/test.txt")
	if err1 != nil {
		log.Fatal("Failed to create local file ", err1)
	}
	defer localPath.Close()

	// _, err2 := io.Copy(localPath, remotefile)
	// if err2 != nil{
	// 	log.Fatal("Failed to put file ", err)
	// }
	// fmt.Println("File written in local successfully")
	buffer := make([]byte, 10)

	for{
		n, err := remotefile.Read(buffer)

		// if (err!=io.EOF) && (err!=nil){
		// 	_, err := localPath.Write(buffer)
		// 	if err!=nil{
		// 		log.Fatal("Failed to write ", err)
		// 	}

		// 	fmt.Println(string(buffer[:n]))
		// 	fmt.Println("100")
		// }

		if err == io.EOF{
			_, err := localPath.Write(buffer)
			if err != nil{
				log.Fatal(err)
			}
			fmt.Println(string(buffer[:n]))
		}

		if err != nil{
			log.Fatal("")
		}
		// if err == io.EOF{
		// 	tot, err := localPath.Write(buffer)
		// 	if err != nil{
		// 		log.Fatal(err)
		// 	}
		// 	if tot != len(buffer[:n]){
		// 		log.Fatal("failed to write")
		// 	}
		// 	fmt.Println("file written succesfully ",string(buffer[:n]))
		// }
	}
}

// func Get(client *sftp.Client) {

//     remoteFile, err := client.Open("/sftp/testint/customer.controller.go")
//     if err != nil {
//         log.Fatal(err)
//     }
//     log.Println("Opened")

//     localFile, err := os.Create("/home/balaji/go/src/testint/test.go")
//     if err != nil {
//         log.Fatal(err)
//     }
//     log.Println("created")

//     // _, err = io.Copy(localFile, remoteFile)
//     // if err != nil {
//     //  log.Fatal(err)
//     // }
//     // log.Println("file uploaded")
//     // var buffer []byte
//     buffer := make([]byte, 1024)
//     for {
//         n, err := remoteFile.Read(buffer)
//         if err == io.EOF {
//             val, err1 := localFile.Write(buffer)
//             if err1 != nil {
//                 log.Println(err)
//             }
//             if val != len(buffer[:n]) {
//                 fmt.Println("failed!")
//             }
//             break
//         }
//         fmt.Print(string(buffer[:n]))
//     }
// }

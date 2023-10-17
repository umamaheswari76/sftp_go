package services

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/pkg/sftp"
)

//put "localFilePath"
//put file from local to remote

func Put(client *sftp.Client) {
	localFile, err := os.Open("/home/vasenth/Documents/umamaheswari/sample.go.json")
	if err != nil {
		log.Fatal("Failed to open local file: ", err)
	}
	
	remotePath, err1 := client.Create("sftp/testint/sample.go.json")
	if err1 != nil {
		log.Fatal("Failed to create remote file: ", err1)
	}
	
	_, err2 := io.Copy(remotePath, localFile)
	if err != nil {
		log.Fatal("Failed to put file: ", err2)
	}
}


//get "filename" localpath
//get file from remote to local

func Get(client *sftp.Client){
	localPath,err := os.Create("/home/vasenth/Documents/umamaheswari/sample1.go.json")
	if err != nil {
		log.Fatal("Failed to create local file: ", err)
	}

	remoteFile,err1 := client.Open("sftp/testint/sample.json")
	if err1 != nil {
		log.Fatal("Failed to open remote file: ", err1)
	}

	_,err2 := io.Copy(localPath,remoteFile)
	if err2 != nil {
		log.Fatal("Failed to put file: ", err2)
	}

}

func List(client *sftp.Client, remotePath string){

	files, err := client.ReadDir(remotePath)
	if err != nil{
		log.Fatal("Unable to list remote dir: ",err)
	}

	for _,f := range files{
		var name string
		name = f.Name()

		if f.IsDir(){
			name = name + "/"
		}
		fmt.Println(name)
	}
}
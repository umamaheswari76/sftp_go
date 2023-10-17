package config

import (
	"fmt"
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func SftpClient()(*sftp.Client) {
	config := &ssh.ClientConfig{
		User: "username",
		Auth: []ssh.AuthMethod{
			ssh.Password("password"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	//client connection to ssh server with the given network address
	conn, err := ssh.Dial("tcp", "hostname", config)
	if err != nil {
		log.Fatal("Failed to dial: ", err)
	}

	client, err1 := sftp.NewClient(conn)
	if err != nil {
		log.Fatal("Failed to create SFTP client: ", err1)
	} else {
		fmt.Println("Client created ")
	}
	return client
}



package config

import (
	"fmt"
	"log"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func SftpConnection() (*sftp.Client, error) {
	config := &ssh.ClientConfig{
		User: "netsys",
		Auth: []ssh.AuthMethod{
			ssh.Password("Netsys@4321!"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	conn, err := ssh.Dial("tcp", "18.216.98.58:22", config)

	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal("Failed to connect")
	}
	fmt.Println("Client ceated")
	return client, nil
}

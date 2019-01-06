package burp

import (
	"../config"
	"os"
	"net"
	"log"
	"golang.org/x/crypto/ssh"
	"fmt"
	"io/ioutil"
	"testing"
)

func SSHServer(username string, password string, port string,id_rsafilepath string){

	config := ssh.ServerConfig{
		//PublicKeyCallback: keyAuth,
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if c.User() == username && string(pass) == password {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}

	privateBytes, err := ioutil.ReadFile(id_rsafilepath)
	if err != nil {
		log.Fatal("Failed to load private key ", id_rsafilepath)
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}

	config.AddHostKey(private)

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	socket, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := socket.Accept()
		if err != nil {
			panic(err)
		}

		// From a standard TCP connection to an encrypted SSH connection
		sshConn, _, _, err := ssh.NewServerConn(conn, &config)
		if err != nil {
			panic(err)
		}

		log.Println("Connection from", sshConn.RemoteAddr())
		sshConn.Close()
	}
}

func TestSSHBrust(t *testing.T) {
	// Start SSH Server
	go SSHServer("user","pass","2222", "../test/id_rsa")

	crackData := models.Boomb{"user","pass"}
	testdata := models.Try{"127.0.0.1", "2222","ssh", &crackData, false}

	SSHBrust(&testdata)
}
package burp

import (
	"../config"
	"golang.org/x/crypto/ssh"
	"fmt"
	"os"
)

func SSHBrust (try *models.Try) *models.Try {

	sshConfig := &ssh.ClientConfig{
		User: try.Data.Username,
		Auth: []ssh.AuthMethod{ssh.Password(try.Data.Password)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", try.Target+":"+try.Port, sshConfig)
	if err != nil {
		fmt.Println("[Error]: Host Can't Access")
		os.Exit(1)
		return nil
	}

	session, err := client.NewSession()
	if err != nil {
		_ = session
		try.Status = true
		client.Close()
		return try
	}

	return nil
}
package burp

import (
	"boomb/config"
	"fmt"
	"golang.org/x/crypto/ssh"
	"os"
	"strings"
)

func SSHBrust (try *models.Try) *models.Try {

	sshConfig := &ssh.ClientConfig{
		User: try.Data.Username,
		Auth: []ssh.AuthMethod{ssh.Password(try.Data.Password)},
	}
	sshConfig.HostKeyCallback = ssh.InsecureIgnoreHostKey()

	client, err := ssh.Dial("tcp", try.Target+":"+try.Port, sshConfig)

	if err == nil {
		_ = client
		try.Status = true
		client.Close()
		return try
	}

	if strings.Contains(err.Error(),"connection refused"){
		fmt.Println("[Error] target connection can't access")
		os.Exit(1)
	}

	return nil
}
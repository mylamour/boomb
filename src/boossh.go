package burp

import (
	"../config"
	"golang.org/x/crypto/ssh"
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

	return nil
}
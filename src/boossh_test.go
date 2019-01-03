package burp
//
//import (
//	"golang.org/x/crypto/ssh"
//	"net"
//	"strings"
//	"testing"
//)
//
//func Test(t *testing.T) {
//	//Start SSH Server
//	//SSHBrust()
//
//	listener, err := net.Listen("tcp", "0.0.0.0:"+com.ToStr(port))
//	if err != nil {
//		panic(err)
//	}
//	for {
//		// Once a ServerConfig has been configured, connections can be accepted.
//		conn, err := listener.Accept()
//		if err != nil {
//			// handle error
//			continue
//		}
//		// Before use, a handshake must be performed on the incoming net.Conn.
//		sConn, chans, reqs, err := ssh.NewServerConn(conn, config)
//		if err != nil {
//			// handle error
//			continue
//		}
//
//		// The incoming Request channel must be serviced.
//		go ssh.DiscardRequests(reqs)
//		go handleServerConn(sConn.Permissions.Extensions["key-id"], chans)
//	}
//}

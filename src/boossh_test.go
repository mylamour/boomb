package burp

import (
	"../config"
	"fmt"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"testing"
)

func SSHServer(username string, password string, port string){

	config := ssh.ServerConfig{
		//PublicKeyCallback: keyAuth,
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if c.User() == username && string(pass) == password {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}

	// privateBytes, err := ioutil.ReadFile(id_rsafilepath)
	// fmt.Println(privateBytes)
	// if err != nil {
	// 	log.Fatal("Failed to load private key ", id_rsafilepath)
	// }

	//Notice: generator from http://travistidwell.com/jsencrypt/demo/

	privateBytes := []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpgIBAAKCAQEA1brQS2A5jsMoO2fWuWCPgFSICgxBoxgTZiVutA1jpGbGT66d
myiMIpPKDnP+OyE0JHaUpSTWmLfbdxhFcgcHBikuluHq9rDHTe4xEef+jaXwBPV3
5AEAUmX7shKNTDTpL2iZK6GaSFAiao8ZTYH4Xpz4tWq2UOqXBd5eP9qs4hqHsbOt
YPJsJMpxmZ8+SVFjx7wgCxcZ4TzdxbIpiGDqRKZwqw8uErZtKmqPOYXKVQ90EjGk
qEnpteS2cXc3sutXjGcWUqselyZJJgdEx665U63m0lublFuK/w8xTiPwlf6RzpyA
a0EhrrHg4QsEHl+7DLzscVt9yBHUKtrxFPB7+QIDAQABAoIBAQCyu9VxE3lO0BmX
BFFnGT8oXIifjnxdfcbLR0Z6wC+xzljNTgNzO8dlKx5wU8QpgqzuSVQpftMPR8H6
qIt6vjG7jzjs6OTrhA5IN4dFqDMSFdDQ9M5cGBJC1hJUantLXXwcL2bYO/6ftpPx
dNX+IVCpHrBGVoU8ydGeEMXUtHSbZ3e3tz5uaWEXJj6WDmGXXep9gLJQMsWx2LiY
kpFw3p+d/KuwdjIceyP42wa/RbAPAUxvtIfJBz4aWGYAawRB9IhHqT6NL4DmQilT
qPRC2lq1ZIBimR8StwpI5WOyE84vm2jka6708o8dvR50oHQhoxcw23nVZVOY8R/G
C8KygvJhAoGBAP5BkuTEFRKqlRK8uEFefLoVR//uCyMmHOAh5o53e1S+2Q2+5H1K
MssS1F6c3aBNv1UWMLqXlEcuWFdeuHfVkyk0R1zYJazrFizbrI1ePL5vlBzM/lS1
nANEQUZG4YvoiDkDS6i7ZWFH+5pXR3vVeU839eThLb/zMlh4887ySwVlAoGBANcy
FUOjxPDrNB61rgzDOQZtzYBqmWvnwmGmUlPHjVig3WZy0tEOjhNRVaNjCuAVQ35C
Wp52UW+p+xKt9Vm+Ae8nXFWvVOI55fFFa6KgcsROgD6OzCZVIUfZX4BwKdm6Dqg9
Xr8TuiN7RdMP9T9Guh4Mr7ySD3MyWv2xfnYs0k0FAoGBAPmYxett0oIQk2dhfEAv
0aGbYqMLvEM1FwOwQ/V3tcrrz4h+7S9Rt6tDQNfghnDn4fxVXGDQjO+Tv4WFpvF8
6Ip6l9O9HL8zyQEgZbQh3t/UCUJ8uu/NmOpcBvhGfQQrsg7F/XAXBt8JGyXYqIQY
fI4YEXwu0gqtY68WjcGKebtBAoGBAJCRuO9kCl6/5Jfs/izcymajRbfL5Z40aMYp
co2ONksgZxz4buC7on8f2SSW2SuJnXopIa/rVPJrg50c+QN9Ptdn3gRCcgg5VV0t
gg4TNIblJgrg7A2ki2M1iY9DyxnWgRpNgGVl31sO6e3NfrkvzsF5TGamyXJelfpx
T8AchHvxAoGBAK8Yo74IB2vWq8NtXVNZkD2HpQa8zH6EMfFJ/YZW2SpPgl7AEJ2W
giWRxZ2AMru+pCLIy+zbf/7S+LCgELG4mLAlIhuArXeQHkSJku6jvc/0a6N3eHwo
LA93RBd7wkMeoXR+9PmrQDzbX1jxVNLFuwvWrkO1J5vmvic/CF5B3M7W
-----END RSA PRIVATE KEY-----`)

	private, _ := ssh.ParsePrivateKey(privateBytes)

	config.AddHostKey(private)

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
	go SSHServer("hhh","lll","2222")

	crackData := models.Boomb{"hhh","lll"}
	testdata := models.Try{"127.0.0.1", "2222","ssh", &crackData, false}

	res := SSHBrust(&testdata)

	if res != nil {
		t.Log("SUCESS")
	}else{
		t.Error("Not Right")
	}
}
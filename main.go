package main

import (
	"./config"
	"./src"
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Fire( fire func(*models.Try) *models.Try, trys []*models.Try) *models.Boomb {
	for _, bob := range trys {
		result := fire(bob)
		if result != nil && result.Status {
			return &models.Boomb{result.Data.Username,result.Data.Password}
		}
	}
	return nil
}

func ArrangeSlic( usernames *[]string, passwords *[]string, result *[]models.Boomb) *[]models.Boomb {

	for _, user := range *usernames {
		for _,pass := range *passwords {
			*result = append(*result, models.Boomb{user,pass})
		}
	}
	return result
}

func IsFileExists(filename string) bool {

	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return  false
		}else {
			return true
		}
	}

	return true
}

func ReadDictFile(user string, authtickes *[]string) *[]string{

	username, err := os.Open(user)
	if err != nil {
		log.Fatal(err)
	}

	defer username.Close()

	scanner := bufio.NewScanner(username)

	for scanner.Scan() {
		*authtickes = append(*authtickes,scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return authtickes
}



func ParserTarget(target string) models.Try {
	var info models.Try
	parser, err := url.Parse(target)

	if err != nil {
		panic(err)
	}
	host, port, _ := net.SplitHostPort(parser.Host)

	if host == ""{
		fmt.Println("Target Host is None")
		os.Exit(1)
	}


	info.Target = host
	info.Protocal = parser.Scheme

	if port == "" {
		if parser.Scheme == "http" {
			info.Port = "80"
		}
		if parser.Scheme == "https" {
			info.Port = "443"
		}
		if parser.Scheme == "ssh" {
			info.Port = "22"
		}
		if parser.Scheme == "ftp" {
			info.Port = "23"
		}
	}else{
		info.Port = port
	}

	return info
}

func main() {

	cmd := os.Args[1]

	brustype := flag.NewFlagSet("brustype",flag.ContinueOnError)

	//targethost := brustype.String("host", "", "Your target host")
	//targetport := brustype.String("port", "", "Your target port")

	target := brustype.String("target", "", "your target")

	userdict := brustype.String("user", "", "Your username filepath")
	passwddict := brustype.String("pass", "", "Your password filepath")

	targetinfo := ParserTarget(*target)

	var user []string
	var pass []string
	var trys []models.Try
	var boomb []models.Boomb

	usernames := &user
	passwords := &pass
	boombs := &boomb
	boombtrys := &trys

	usernames = ReadDictFile(*userdict, usernames)
	passwords = ReadDictFile(*passwddict, passwords)
	boombs = ArrangeSlic(usernames, passwords, boombs)

	for

	in := models.Boomb{*userdict,*passwddict}

	crackdata := models.Try{"127.0.0.1", "8080","http", &in, false}

	switch cmd {

	case "ssh":
		if err := brustype.Parse(os.Args[2:]); err == nil {
			if !IsFileExists(*userdict) || !IsFileExists(*passwddict) {
				fmt.Println("Please make sure your dict was exists")
				os.Exit(0)
			}
			Fire(burp.SSHBrust, &crackdata)
			//fmt.Println("ssh brust",*userdict, *passwddict)
		}
	case "http":
		if err := brustype.Parse(os.Args[2:]); err == nil {
			fmt.Println("http brust",*userdict, *passwddict )

		}

	default:
		brustype.Usage()
		fmt.Println("example: ")
		fmt.Println("\tboomb ssh --user userdict --pass passdict")
		fmt.Println("\tboomb http --user userdict --pass passdict")

	}
}

func test() {

	var user []string
	var pass []string
	var boomb []models.Boomb
	usernames := &user
	passwords := &pass
	boombs := &boomb

	usernames = ReadDictFile("test/dict/user.list", usernames)
	passwords = ReadDictFile("test/dict/pass.list", passwords)

	boombs = ArrangeSlic(usernames, passwords, boombs)

	for _, v := range *boombs {
		crackData := models.Boomb{v.Username, v.Password}
		testdata := models.Try{"baidu.com", "80","http", &crackData, false}

		res := Fire(burp.SSHBrust, &testdata)

		if res != nil {
			fmt.Println("Cracked: ", res.Username, res.Password)
		}
	}
}
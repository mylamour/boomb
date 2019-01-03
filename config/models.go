package models

//type Burst interface {
//	Fire()
//}

type Try struct {
	Target string	// ip or hostname
	Port string
	Protocal string	// http ssh and what ever
	Data *Boomb		// burp force auth ticket for try
	Status bool		// sucessful or not
}

type Boomb struct {
	Username string
	Password string
}


var userDictPath = ""
var passwdDictPath = ""
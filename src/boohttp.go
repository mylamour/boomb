package burp

import (
	Models "../config"
	"net/http"
	"log"
	"fmt"
)

func HTTPBrust(try *Models.Try) *Models.Try{
	//Basic Auth Brust
	client := &http.Client{}
	req, err := http.NewRequest("GET", try.Protocal + "://" + try.Target + ":" + try.Port , nil)
	req.SetBasicAuth(try.Data.Username, try.Data.Password)
	resp, err := client.Do(req)
	if err != nil{
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		try.Status = true
		fmt.Println("CRACKED: ",try.Data.Username, try.Data.Password)
		return try
	}

	return nil
}

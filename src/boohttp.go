package burp

import (
	"boomb/config"
	"net/http"
	"log"
)

func HTTPBrust(try *models.Try) *models.Try{
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
		return try
	}

	return nil
}

func HTTPSBrust(try *models.Try) *models.Try {
	//Todo Add Https basic auth brust
	return nil
}
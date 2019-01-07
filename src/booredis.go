package burp

import (
	"boomb/config"
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strings"
)
func REDISBrust(try *models.Try) *models.Try{

	client := redis.NewClient(&redis.Options{
		Addr:     try.Target+":"+try.Port,
		Password: try.Data.Password, // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err == nil{
		_ = pong
		try.Status = true
		return try
	}

	switch {

	case strings.Contains(err.Error(),"connect: connection refused"):
		fmt.Println("[Error] target connection can't access")
		os.Exit(1)
	case strings.Contains(err.Error(),"no password is set"):
		fmt.Println("[Info] target password is none")
		os.Exit(0)
	default:
		fmt.Println(err.Error())

	}

	return nil
}

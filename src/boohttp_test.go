package burp

import (
	"../config"
	"encoding/base64"
	"net/http"
	"strings"
	"testing"
)

const (
	username = "user"
	password = "pass"
)

func checkAuth(w http.ResponseWriter, r *http.Request) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 { return false }

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil { return false }

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 { return false }

	return pair[0] == username && pair[1] == password
}

func Server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if checkAuth(w, r) {
			w.Write([]byte("Cracked"))
			return
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="MY REALM"`)
		w.WriteHeader(401)
		w.Write([]byte("401 Unauthorized\n"))
	})

	http.ListenAndServe(":8080", nil)
}

func TestHTTPBrust(t *testing.T) {
	// Start HTTP Server
	Server()
	// Begin To Burp
	crackData := models.Boomb{"user","pass"}
	testdata := models.Try{"127.0.0.1", "8080","http", &crackData, false}

	res := HTTPBrust(&testdata)

	if res != nil {
		t.Log("SUCESS")
	}else{
		t.Error("Not Right")
	}

}
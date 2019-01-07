package burp

import (
	"boomb/config"
	"testing"
)

func TestREDISBrust(t *testing.T) {

	// Before your test, you should setting your redis server with password
	// CONFIG SET REQUIREPASS whoami

	crackData := models.Boomb{"","whoami"}
	testdata := models.Try{"127.0.0.1", "6379","ssh", &crackData, false}

	res := REDISBrust(&testdata)

	if res != nil {
		t.Log("SUCESS")
	}else{
		t.Error("Not Right")
	}
}
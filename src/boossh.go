package burp

import "../config"

func SSHBrust (try *models.Try) *models.Try {
	if try.Data.Username == "ssss" && try.Data.Password == "bbbb" {
		try.Status = true
		return try
	}
	return nil
}
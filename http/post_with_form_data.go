package main

import (
	"encoding/base64"
	"github.com/theplant/qortexapi"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	data := url.Values{
		"UserId":      {"1231321"},
		"Message":     {"Hello World"},
		"AlertNumber": {"11"},
		"ItemId":      {"2222"},
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", "http://localhost:3000/push", strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth("admin", "guessme")
	client.Do(req)
}

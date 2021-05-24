package auth

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"
	"time"
)
var ACCESS_TOKEN = ""
var LAST_REFRESH = time.Now()
var REFRESH_RATE = -3000.0
var CLIENT_ID = "a00eef974890430c94e90dfda54fab28"
var CLIENT_SECRET = "8d170068cd6d436db439f785b7bcdac7"

func ValidateToken(){
	if LAST_REFRESH.Sub(time.Now()).Minutes() < REFRESH_RATE || ACCESS_TOKEN == "" {
		url := "https://accounts.spotify.com/api/token"

		data := url2.Values{}
		data.Set("grant_type", "client_credentials")

		enc:= base64.StdEncoding.EncodeToString([]byte(CLIENT_ID+":"+CLIENT_SECRET))

		auth:= "Basic " + enc

		req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
		req.Header.Add("Authorization", auth)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error on response.\n[ERROR] -", err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Error while reading the response bytes:", err)
		}

		var res map[string]interface{}
		json.Unmarshal(body, &res)

		ACCESS_TOKEN = res["access_token"].(string)
	}
}
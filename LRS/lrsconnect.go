package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"time"
)

const followRedirectsHeader = "LRS-FollowRedirects"

var client *Client

type Client struct {
	HTTPClient *http.Client
}

//Init initializes the underlying client that communicates with Stormpath
func Init() {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{},
		DisableCompression: true,
	}

	cookieJar, _ := cookiejar.New(nil)

	httpClient := &http.Client{Transport: tr,
		Jar: cookieJar} //Added to send cookies.

	client = &Client{httpClient}
}

func (client *Client) newRequestWithoutRedirects(method string, urlStr string, body interface{}) *http.Request {
	req := client.newRequest(method, urlStr, body)
	req.Header.Add(followRedirectsHeader, "false")
	return req
}

func (client *Client) newRequest(method string, urlStr string, body interface{}) *http.Request {
	//Works : Passing JSON String in bytes. []byte{<JSONString>} -> that reads to NewRequest.
	//	req, _ := http.NewRequest(method, urlStr, bytes.NewBuffer(body.([]byte)))

	//Works : Passing Structure :: struct -> []byte -> thats reads to NewRequest
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(method, urlStr, bytes.NewReader(jsonBody))

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Host", "devconnect.lrsus.com")
	req.Header.Set("Accept-Encoding", "gzip, deflate, sdch")
	req.Header.Set("Origin", "https://devconnect.lrsus.com")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	return req
}

//Impleemnting new Date.getTime() from JAVASCRIPT
func JSgetTime() int64 {
	test := time.Date(1970, 01, 01, 0, 0, 0, 0, time.UTC)
	//test, err := time.Parse("01/02/2006", "01/01/1970")

	ms1 := test.UnixNano() / int64(time.Millisecond)
	ms2 := time.Now().UnixNano() / int64(time.Millisecond)
	res := ms2 - ms1
	return res
}

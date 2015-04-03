package main

import (
	"encoding/json"
	"fmt"
	"log"
	//	"net/http"
	//	"os"
	//	"bytes"
	//	"io/ioutil"
)

var (
	STORMPATH_API_KEY_ID     = "21KQV3Z4NL77XVISQNWTFADN1"
	STORMPATH_API_KEY_SECRET = "80Twrr3oDxCJ90SvYrMzy5kNBhvGM0S4O/tyWZ+/TT8"
	TEST_PREFIX              string
	CLIENT                   *Client
)

func main() {
	fmt.Println("Program starts....")
	//	fmt.Println(STORMPATH_API_KEY_SECRET)
	client, err := NewClient(&ApiKeyPair{
		Id:     STORMPATH_API_KEY_ID,
		Secret: STORMPATH_API_KEY_SECRET,
	})
	if err != nil {
		log.Fatal("Couldn't create a Stormpath client.")
	}
	CLIENT = client

	if client.Tenant.Href == "" {
		log.Fatal("Couldn't create a Stormpath client.")
	}

	//	fmt.Printf("Tenent %v \n", client.Tenant.Href)
	//	Test_Request()

	TestLRSLogin()

}

type stormpathError struct {
	Status           int
	Code             int
	Message          string
	DeveloperMessage string
	MoreInfo         string
}

type Tenant1 struct {
	Href         string `json:"href"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	Applications link   `json:"applications"`
	Directories  link   `json:"directories"`
}
type link struct {
	Href string `json:"href"`
}

func emptyPayload() []byte {
	return []byte{}
}

func Test_Request() {
	resp, err := CLIENT.Request("GET", "/tenants/current", emptyPayload())

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 204 && resp.StatusCode != 201 && resp.StatusCode != 302 {
		spError := &stormpathError{}
		err := json.NewDecoder(resp.Body).Decode(spError)
		if err != nil {
			fmt.Printf("%s [%s]", spError.Message, resp.Request.URL.String())
		}
	}

	fmt.Printf("First Response. %v \n", resp)

	//	response, newerr := CLIENT.VDo("GET", resp.Header.Get("Location"), emptyPayload())
	response, newerr := CLIENT.Request("GET", resp.Header.Get("Location"), emptyPayload())

	if newerr != nil {
		log.Fatal(newerr)
	}
	defer response.Body.Close()

	fmt.Printf("Second  Response. %v \n", response)

	ten := &Tenant1{}
	err = json.NewDecoder(response.Body).Decode(ten)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ten.Name)

}

var JSONSTR = `{"password": "password1"
              "rememberMe": false
              "username": "admin@lrs.com"}`

func TestLRSLogin() {
	resp, err := CLIENT.LRSCall("POST", "https://devconnect.lrsus.com/web/v3/login", []byte(JSONSTR))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
}

//http://stackoverflow.com/questions/19228514/how-to-decompress-a-byte-content-in-gzip-format-that-gives-an-error-when-unmar
//http://stackoverflow.com/questions/21268000/unmarshaling-nested-json-objects-in-golang

//Further REad
//http://attilaolah.eu/2013/11/29/json-decoding-in-go/

package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

type InitLoad struct {
	Password   string `json:"password"`
	RememberMe bool   `json:"rememberMe"`
	Username   string `json:"username"`
}

type GenricObject struct {
	ReturnObject interface{} `json:"returnObject"`
}

type OathReturnObject struct {
	Expires_in   int    `json:"expires_in"`
	Token_type   string `json:"token_type"`
	Access_token string `json:"access_token"`
}

type OathResult struct {
	ResponseType string `json:"responseType"`
	ErrorDetail  string `json:"errorDetail"`
	ReturnObject string `json:"returnObject"`
}

const lrsLogin = "https://devconnect.lrsus.com/web/v3/login"
const lrsOath = "https://devconnect.lrsus.com/rest/v3/oauth/token"
const lrsGetCurrentBrand = "https://devconnect.lrsus.com/web/brand/getCurrentBrand?time=%d"
const lrsV3Authorise = "https://devconnect.lrsus.com/web/v3/authorize/eaa6a07a-d4d1-44bb-b826-b0a6548187e0?userAccountPlanProductGuid=58a2ab9a-1e13-11e4-8d3d-22000a4981fc&time=%d"

//const lrsReportCall = "https://devconnect.lrsus.com/web/v3/reports?queryTypes=AVERAGE_DELIVERY_TIME_OVER_TIME&dateFrom=2015-03-17&dateTo=2015-03-23&locationIds=all&reportSetIds=&goal=90&time=%d"
const lrsReportCall = "https://devconnect.lrsus.com/web/v3/authorize/eaa6a07a-d4d1-44bb-b826-b0a6548187e0?userAccountPlanProductGuid=58a70d37-1e13-11e4-8d3d-22000a4981fc&time=1427100690526"

func main() {
	fmt.Println("Program Starts...")

	Init()
	initLoad := InitLoad{"Password1", false, "admin@lrs.com"}
	//When we pass JSOn in this format it is visible in Request header.
	//var jsonStr = []byte(`{ "password": "Password1","rememberMe": false,"username": "admin@lrs.com" }`)
	//request := client.newRequest("POST", lrsLogin, jsonStr)

	request := client.newRequest("POST", lrsLogin, initLoad)
	response, err := client.HTTPClient.Transport.RoundTrip(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	//	fmt.Println("Sucessfull..")
	//	fmt.Printf("%v %v\n", response.Status, response.Cookies()[0])
	//	fmt.Println("---------------------------------------------------------")

	request = client.newRequest("GET", lrsOath, initLoad)
	for _, c := range response.Cookies() {
		request.AddCookie(c)
	}

	fmt.Println(request.Cookies())
	response, err = client.HTTPClient.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Oauth %v %v\n", response.Status, response)

	//	oathRes := &OathResult{}
	//	err = json.NewDecoder(response.Body).Decode(oathRes)
	//	if err != nil {
	//		fmt.Printf("Error in decoding JSON %v", err.Error())
	//	}
	//	fmt.Printf(" Oasth Reponse %v \n", oathRes)

	fmt.Println("===============================================")

	//content = body
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	// decompress the content into an io.Reader
	buf := bytes.NewBuffer(content)
	reader, err := gzip.NewReader(buf)
	if err != nil {
		panic(err)
	}

	var data OathResult
	dec := json.NewDecoder(reader)
	err = dec.Decode(&data)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Println(data.ReturnObject)

	var oathRetObj OathReturnObject
	json.Unmarshal([]byte(data.ReturnObject), &oathRetObj)
	fmt.Printf("Data :| %d %s \n %s\n", oathRetObj.Expires_in, oathRetObj.Token_type, oathRetObj.Access_token)

	response.Body.Close()

	fmt.Println("============REPORT CALL ==========================")
	//	lrsGetCurrentBrand_ := fmt.Sprintf(lrsGetCurrentBrand, JSgetTime())
	//	fmt.Println(lrsGetCurrentBrand_)
	//	request = client.newRequest("GET", lrsGetCurrentBrand_, initLoad)

	lrsV3Authorise_ := fmt.Sprintf(lrsV3Authorise, JSgetTime())
	fmt.Println(lrsV3Authorise)
	request = client.newRequest("GET", lrsV3Authorise_, initLoad)

	//	lrsReportCall_ := fmt.Sprintf(lrsReportCall, JSgetTime())
	//	fmt.Println(lrsReportCall)
	//	request = client.newRequest("GET", lrsReportCall_, initLoad)

	request = client.newRequest("GET", lrsReportCall, initLoad)

	authorizationStr := fmt.Sprintf("%s %s", oathRetObj.Token_type, oathRetObj.Access_token)
	request.Header.Set("Authorization", authorizationStr)
	request.Header.Set("Referer", "https://devconnect.lrsus.com/")
	//	req.Header.Set("Content-Type", "application/json")
	for _, c := range response.Cookies() {
		request.AddCookie(c)
	}

	//	fmt.Printf("-------------> %v \n", request.Cookies())

	response, err = client.HTTPClient.Transport.RoundTrip(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf(" Request ===> %s \n", request)

	fmt.Printf("Response Returned %s Content Length  %d\n Response location %v\n", response.Status, response.ContentLength, response.Header.Get("Location"))
	//	d, e := ioutil.ReadAll(response.Body)
	//	if e != nil {
	//		fmt.Println(e)
	//	}
	//	fmt.Println(string(d))

	//	//content = body
	//	content, err = ioutil.ReadAll(response.Body)
	//	if err != nil {
	//		panic(err)
	//	}

	//	var gdata interface{}
	//	err = json.Unmarshal(content, &gdata)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//	fmt.Println("Data from response", gdata)

	//	// decompress the content into an io.Reader
	//	buf = bytes.NewBuffer(content)
	//	reader, err = gzip.NewReader(buf)
	//	if err != nil {
	//		panic(err)
	//	}
	//	//	fmt.Println(reader)
	//	var gdata GenricObject
	//	dec = json.NewDecoder(reader)
	//	err = dec.Decode(&gdata)
	//	if err != nil && err != io.EOF {
	//		panic(err)
	//	}
	//	fmt.Println(gdata.ReturnObject)

	response.Body.Close()
}

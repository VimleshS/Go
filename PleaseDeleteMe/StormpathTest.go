package main

import (
	"fmt"
	"os"
)

func main() {

	//This would look for env variables first STORMPATH_API_KEY_ID and STORMPATH_API_KEY_SECRET if empty
	//then it would look for os.Getenv("HOME") + "/.config/stormpath/apiKey.properties" for the credentials
	credentials, _ := NewDefaultCredentials()
	fmt.Println(credentials)

	//Init Whithout cache
	Init(credentials, nil)

	//Get the current tenant

	tenant, err := CurrentTenant()
	fmt.Println()
	if err != nil {
		fmt.Println("Error 1")
		os.Exit(1)
	}
	fmt.Println(tenant.Name)
}

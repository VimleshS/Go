package main

import (
	"fmt"
	"github.com/jarias/stormpath-sdk-go"
	"os"
)

func main() {

	//This would look for env variables first STORMPATH_API_KEY_ID and STORMPATH_API_KEY_SECRET if empty
	//then it would look for os.Getenv("HOME") + "/.config/stormpath/apiKey.properties" for the credentials
	credentials, _ := stormpath.NewDefaultCredentials()
	fmt.Println(credentials)

	//Init Whithout cache
	stormpath.Init(credentials, nil)

	//Get the current tenant
	fmt.Println("----Before stormpath.CurrentTenant---")
	tenant, err := stormpath.CurrentTenant()
	fmt.Println()
	fmt.Println("############################")
	if err != nil {
		fmt.Println("Error 1")
		os.Exit(1)
	}

	//Get the tenat applications

	apps, erra := tenant.GetApplications(stormpath.NewDefaultPageRequest(), stormpath.NewEmptyFilter())
	if erra != nil {
		fmt.Printf("Error 2 %v", erra)
		os.Exit(1)
	}
	//	fmt.Println("\n----ListOfApps----")
	//	fmt.Println(apps)
	//	fmt.Println("----ListOfApps----")
	//Get the first application
	app := apps.Items[1]

	//Authenticate a user against the app
	fmt.Println("----------------------")
	fmt.Println(app.Href)
	fmt.Println("----------------------")
	accountRef, errn := app.AuthenticateAccount("Vimlesh.Sharma@synerzip.com", "newPass1s")
	if errn != nil {
		fmt.Printf("Error 3 %v", errn)
		os.Exit(1)
	}

	//Print the account information
	account, _ := accountRef.GetAccount()

	fmt.Println(account)
}

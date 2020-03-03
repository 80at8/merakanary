package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"strings"
)

type orgList struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	Networks []networkList
	//Inventory []inventoryList
}

// globals

var organization []orgList
var apiBase string = "https://api.meraki.com"

// insert your key here
var apiKey string = ""

// DISPLAYFLAG -- SORT FLAG -- DEBUG (Options)
var DISPLAYFLAG = flag.String("display", "CLIENTS", "display: CLIENTS|NETWORKS|ORG")
var SORTFLAG = flag.String("sort", "IP", "sort by: IP|MAC|OS|MFG|STATUS|FIRSTSEEN|LASTSEEN")
var DEBUG = flag.Bool("debug", false, "enable debug output: TRUE|FALSE")

func getOrganizations() error {

	request, err := getAPIRequest("organizations")

	if err != nil {
		return err
	}

	err = json.Unmarshal(request, &organization)

	if err != nil {
		return err
	}

	return nil
}

func main() {

	fmt.Println("Merakanary")

	flag.Parse()

	getOrganizations()

	switch strings.ToUpper(*DISPLAYFLAG) {

	case "CLIENTS":
		displayClients()
	case "NETWORKS":
		displayNetworks()
	case "ORG":
		displayOrg()

	default:
		displayClients()
	}

}

package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/landoop/tableprinter"
// )

// type inventoryList struct {
// 	Mac       string `json:"mac"`
// 	Serial    string `json:"serial"`
// 	NetworkID string `json:"networkId"`
// 	Model     string `json:"model"`
// 	ClaimedAt string `json:"claimedAt"`
// 	PublicIP  string `json:"publicIp"`
// 	Name      string `json:"name"`
// }

// func getInventory(organizationID string, index int) error {

// 	request, err := http.NewRequest(http.MethodGet, "https://api.meraki.com/api/v0/organizations/"+organizationID+"/inventory", nil)

// 	if err != nil {

// 		return err

// 	}

// 	request.Header.Set("X-Cisco-Meraki-API-Key", apiKey)

// 	client := http.Client{

// 		Timeout: time.Duration(5 * time.Second),
// 	}

// 	response, err := client.Do(request)

// 	if err != nil {

// 		return err

// 	}

// 	fmt.Printf("getInventory() Response Code: %d\n", response.StatusCode)

// 	body, err := ioutil.ReadAll(response.Body)

// 	if err != nil {

// 		return err

// 	}

// 	err = json.Unmarshal(body, &organization[index].Inventory)

// 	if err != nil {

// 		return err

// 	}

// 	return nil

// }

// func showInventory() {

// 	printer := tableprinter.New(os.Stdout)

// 	switch strings.ToUpper(*OUTPUTFLAG) {

// 	case "CSV":

// 		printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = false, false, false, false

// 		printer.ColumnSeparator = ","

// 		printer.RowSeparator = ""

// 	case "TABLE":

// 		printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true

// 		printer.CenterSeparator = "│"

// 		printer.ColumnSeparator = "│"

// 		printer.RowSeparator = "─"

// 	}

// 	for i := range organization {

// 		for j := 0; j < len(organization[i].Inventory); j++ {

// 			printer.Print(organization[i].Inventory[j])

// 		}

// 	}

// }

// func displayInventory() {
// 	for i := range organization {

// 		for j := 0; j < len(organization[i].Networks); j++ {

// 			getInventory(organization[i].[j].NetworkID, i, j)

// 		}

// 	}

// 	showInventory()
// }

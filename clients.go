package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/landoop/tableprinter"
)

type clientList struct {
	ID                 string `json:"id" header:"ID"`
	Mac                string `json:"mac" header:"MAC"`
	Description        string `json:"description" header:"DESC"`
	IP                 string `json:"ip" header:"IP"`
	IP6                string `json:"ip6" header:"IP6"`
	User               string `json:"user" header:"USER"`
	FirstSeen          string `json:"firstSeen" header:"FIRST-SEEN"`
	LastSeen           string `json:"lastSeen" header:"LAST-SEEN"`
	Manufacturer       string `json:"manufacturer" header:"MFG"`
	Os                 string `json:"os" header:"OS"`
	RecentDeviceSerial string `json:"recentDeviceSerial"`
	RecentDeviceName   string `json:"recentDeviceName"`
	RecentDeviceMac    string `json:"recentDeviceMac"`
	Ssid               string `json:"ssid" header:"SSID"`
	Vlan               int    `json:"vlan" header:"VLAN"`
	Switchport         string `json:"switchport" header:"SWITCHPORT"`
	Usage              struct {
		Sent int `json:"sent" header:"TX"`
		Recv int `json:"recv" header:"RX"`
	} `json:"usage" header:"USAGE"`
	Status string `json:"status" header:"STATUS"`
}

// Sort functions -- these satisfy sort interface

type clientSort []clientList

func (a clientSort) Len() int {
	return len(a)
}

func (a clientSort) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a clientSort) Less(i, j int) bool {
	switch strings.ToUpper(*SORTFLAG) {
	case "FIRSTSEEN":
		return a[i].FirstSeen < a[j].FirstSeen
	case "LASTSEEN":
		return a[i].LastSeen < a[j].LastSeen
	case "MAC":
		return a[i].Mac < a[j].Mac
	case "MFG":
		return a[i].Manufacturer < a[j].Manufacturer
	case "OS":
		return a[i].Os < a[j].Os
	case "STATUS":
		return a[i].Status < a[j].Status
	case "USER":
		return a[i].User < a[j].User
	default:
		return a[i].IP < a[j].IP
	}
}

func getNetworkClients(networkID string, orgIndex int, netIndex int) error {
	request, err := getAPIRequest("networks", networkID, "clients?perPage=1000")

	if err != nil {
		return err
	}

	err = json.Unmarshal(request, &organization[orgIndex].Networks[netIndex].Clients)

	if err != nil {
		return err
	}

	sort.Sort(clientSort(organization[orgIndex].Networks[netIndex].Clients))
	return nil
}

func showClients() {
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"

	for i := range organization {
		for j := 0; j < len(organization[i].Networks); j++ {
			if len(organization[i].Networks[j].Clients) > 0 {
				str := "| " + organization[i].Networks[j].Name + " | " + organization[i].Networks[j].Clients[0].RecentDeviceName + " | " + organization[i].Networks[j].Clients[0].RecentDeviceMac + " | " + organization[i].Networks[j].Clients[0].RecentDeviceSerial + "|"

				div := func() string {
					tmp := ""
					for x := 0; x < len(str); x++ {
						tmp += "─"
					}
					return tmp
				}()

				fmt.Println(div)
				fmt.Println(str)
				fmt.Println(div)
			}

			printer.Print(organization[i].Networks[j].Clients)
		}
	}
}

func displayClients() {

	for i, org := range organization {

		getNetworks(org.ID, i)

		for j := 0; j < len(organization[i].Networks); j++ {
			getNetworkClients(organization[i].Networks[j].NetworkID, i, j)
		}

	}

	showClients()

}

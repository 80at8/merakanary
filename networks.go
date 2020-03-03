package main

import (
	"encoding/json"
	"os"

	"github.com/landoop/tableprinter"
)

type networkList struct {
	NetworkID      string `json:"id" header:"NETWORK-ID"`
	Name           string `json:"name" header:"NAME"`
	OrganizationID string `json:"organization_id" header:"ORGANIZATION"`
	Type           string `json:"type" header:"TYPE"`
	TimeZone       string `json:"timeZone" header:"TIMEZONE"`
	Tags           string `json:"tags" header:"TAGS"`
	Clients        []clientList
}

func getNetworks(organizationID string, index int) error {

	request, err := getAPIRequest("organizations", organizationID, "networks")

	if err != nil {
		return err
	}

	err = json.Unmarshal(request, &organization[index].Networks)

	if err != nil {
		return err
	}

	return nil

}

func showNetworks() {

	printer := tableprinter.New(os.Stdout)

	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"

	for i := range organization {
		printer.Print(organization[i].Networks)
	}

}

func displayNetworks() {

	for i := range organization {
		getNetworks(organization[i].ID, i)
	}

	showNetworks()
}

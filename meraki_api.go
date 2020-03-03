package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func getAPIRequest(resources ...string) ([]byte, error) {

	tmp := "/api/v0"

	for _, r := range resources {
		tmp += "/" + r
	}

	u, err := url.Parse(tmp)

	if err != nil {
		return nil, err
	}

	b, err := url.Parse(apiBase)

	if err != nil {
		return nil, err
	}

	resourcePath := b.ResolveReference(u).String()

	if *DEBUG {
		fmt.Printf("getAPIRequest(): resourcePath is %v\n", resourcePath)
	}

	request, err := http.NewRequest(http.MethodGet, resourcePath, nil)

	if err != nil {
		return nil, err
	}

	request.Header.Set("X-Cisco-Meraki-API-Key", apiKey)

	client := http.Client{
		Timeout: time.Duration(5 * time.Second),
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	if *DEBUG == true {
		fmt.Printf("getAPIRequest(): X-Cisco-Meraki-API-Key : %v", apiKey)
		fmt.Printf("getAPIRequest(): response code: %d\n", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

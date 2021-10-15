package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
)

func main() {
	timeout := 10000 * time.Millisecond

	client := httpclient.NewClient(httpclient.WithHTTPTimeout(timeout))

	res, err := client.Get("https://cl-docs-addresses.web.app/addresses.json", nil)
	if err != nil {
		fmt.Println("Could not fetch chainlink addresses", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Could not read response body ", err)
	}

	ioutil.WriteFile("addresses.json", body, 0644)
}

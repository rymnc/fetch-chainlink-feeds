package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/gojek/heimdall/v7/httpclient"
	"github.com/joho/godotenv"
)

func getABI(url string, client *httpclient.Client, result chan []interface{}) {
	v, err := client.Get(url, nil)
	if err != nil {
		fmt.Println("Could not fetch contract ABI ", err)
	}
	abiBody, err := ioutil.ReadAll(v.Body)
	if err != nil {
		fmt.Println("Could not read response body ", err)
	}
	var f map[string]interface{}
	if err := json.Unmarshal([]byte(abiBody), &f); err != nil {
		panic(err)
	}
	var abiFormat []interface{}
	z := f["result"].(string)
	if err := json.Unmarshal([]byte(z), &abiFormat); err != nil {
		fmt.Println("Could not unmarshal ABI ", err)
	}
	result <- abiFormat
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
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

	f := make(map[string]Response)
	if err := json.Unmarshal([]byte(body), &f); err != nil {
		panic(err)
	}

	for _, element := range f {
		for _, network := range element.Networks {
			if network.Name == os.Getenv("NETWORK") {
				MakeNetworkDir(network.Name)
				var ProxyJson Proxy
				for _, proxy := range network.Proxies {
					if strings.Contains(network.Url, "etherscan") {
						baseUrl := strings.Split(network.Url, "/address")[0]
						var apiUrl string
						if strings.Contains(baseUrl, "https://etherscan.io") {
							apiUrl = strings.ReplaceAll(
								baseUrl,
								"https://",
								"https://api.",
							)
						} else {
							apiUrl = strings.ReplaceAll(
								baseUrl,
								"https://",
								"https://api-",
							)
						}

						formattedUrl := apiUrl + "/api?module=contract&action=getabi&address=" + proxy.Proxy + "&apikey=" + os.Getenv("ETHERSCAN_API_KEY")
						result := make(chan []interface{})
						go getABI(formattedUrl, client, result)
						value := <-result
						ProxyJson = Proxy{
							Pair:               proxy.Pair,
							DeviationThreshold: proxy.DeviationThreshold,
							Heartbeat:          proxy.Heartbeat,
							Decimals:           proxy.Decimals,
							Proxy:              proxy.Proxy,
							ABI:                value,
						}
					} else {
						fmt.Printf("Network not supported: %s\n", network.Name)
					}
					UpdateNetworkFile(network.Name, ProxyJson)
				}
			}
		}
	}
}

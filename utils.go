package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func MakeDir(dir string) {
	// make a directory but don't fail if it already exists
	err := os.MkdirAll(dir, 0755)
	if err != nil {
	}
}

func GetNetworkDir(network string) string {
	return fmt.Sprintf("output/%s", network)
}

func MakeNetworkDir(network string) {
	MakeDir(GetNetworkDir(network))
}

func GetNetworkFile(network string) string {
	return fmt.Sprintf("%s/addresses.json", GetNetworkDir(network))
}

func WriteNetworkToFile(network string, proxy map[string]Proxy) {
	v, err := json.MarshalIndent(proxy, "", "\t")
	if err != nil {
		fmt.Println("Could not marshal response ", err)
	}
	err = ioutil.WriteFile(GetNetworkFile(network), v, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

// UpdateNetworkFile Read addresses.json from the network and combine with given network and write back to file
func UpdateNetworkFile(network string, proxy Proxy) {
	var oldProxy map[string]Proxy
	// read the file
	f, err := ioutil.ReadFile(GetNetworkFile(network))
	// if file does not exist, create it
	if err != nil {
		p := make(map[string]Proxy)
		p[proxy.Pair] = proxy
		WriteNetworkToFile(network, p)
		return
	}

	// unmarshal the json
	err = json.Unmarshal(f, &oldProxy)
	if err != nil {
		fmt.Println("Could not unmarshal json ", err)
	}
	// combine old proxy and new proxy json
	oldProxy[proxy.Pair] = proxy

	// marshal the combined json
	v, err := json.MarshalIndent(oldProxy, "", "\t")
	if err != nil {
		fmt.Println("Could not marshal response ", err)
	}
	// write the combined json to file
	err = ioutil.WriteFile(GetNetworkFile(network), v, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

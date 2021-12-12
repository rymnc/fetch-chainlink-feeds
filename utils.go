package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func MakeDir(dir string) {
	// make a directory but don't fail if it already exists
	os.MkdirAll(dir, 0755)
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

func WriteNetworkToFile(network string, proxy []interface{}) {
	v, err := json.MarshalIndent(proxy, "", "\t")
	if err != nil {
		fmt.Println("Could not marshal response ", err)
	}
	ioutil.WriteFile(GetNetworkFile(network), v, 0644)
}

// Read addresses.json from the network and combine with given network and write back to file
func UpdateNetworkFile(network string, proxy Proxy) {
	var oldProxy []Proxy
	// read the file
	f, err := ioutil.ReadFile(GetNetworkFile(network))
	// if file does not exist, create it
	if err != nil {
		WriteNetworkToFile(network, make([]interface{}, 0))
	}

	// unmarshal the json
	err = json.Unmarshal(f, &oldProxy)
	if err != nil {
		fmt.Println("Could not unmarshal json ", err)
	}
	// combine the two
	newProxy := append(oldProxy, proxy)

	// marshal the combined json
	v, err := json.MarshalIndent(newProxy, "", "\t")
	if err != nil {
		fmt.Println("Could not marshal response ", err)
	}
	// write the combined json to file
	ioutil.WriteFile(GetNetworkFile(network), v, 0644)
}

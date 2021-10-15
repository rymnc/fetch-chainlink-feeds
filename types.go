package main

type Proxy struct {
	Pair               string        `json:"pair"`
	DeviationThreshold float32       `json:"deviationThreshold"`
	Heartbeat          string        `json:"heartbeat"`
	Decimals           int           `json:"decimals"`
	Proxy              string        `json:"proxy"`
	ABI                []interface{} `json:"abi"`
}

type Network struct {
	Name    string  `json:"name"`
	Url     string  `json:"url"`
	Proxies []Proxy `json:"proxies"`
}

type Response struct {
	Title    string    `json:"title"`
	Networks []Network `json:"networks"`
}

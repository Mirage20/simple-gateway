package config

import (
	"encoding/json"
	"io/ioutil"
)

type Route struct {
	Match       Match       `json:"match"`
	Destination Destination `json:"destination"`
}

type Match struct {
	Exact  []string `json:"exact"`
	Prefix []string `json:"prefix"`
	Regex  []string `json:"regex"`
}

type Destination struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

//func Save(routes []Route) {
//	jsonB, err := json.Marshal(routes)
//	if err != nil {
//		log.Fatal(err)
//	}
//	err = ioutil.WriteFile("/home/miraj/WSO2/go/src/github.com/mirage20/simple-gateway/output.json", jsonB, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//}

func Load(file string) ([]Route, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var routes []Route
	err = json.Unmarshal(data, &routes)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	person := make(map[string]string)
	var name string
	var addr string
	fmt.Scan(&name)
	fmt.Scan(&addr)
	person["name"] = name
	person["address"] = addr
	barr, err := json.Marshal(person)
	if err != nil {
		return
	}
	fmt.Println(string(barr))
}

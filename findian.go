package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	tmp := strings.ToLower(s)
	if tmp[0] == 'i' && strings.Contains(tmp, "a") && tmp[len(tmp)-1] == 'n' {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}

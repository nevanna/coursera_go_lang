package main

import "fmt"

func main() {
	var nb float64
	b, err := fmt.Scan(&nb)
	if b == 0 || err != nil {
		return
	}
	var n int = int(nb)
	fmt.Printf("%d\n", n)
}

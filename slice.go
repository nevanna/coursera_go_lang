package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var el int
	var buf string

	sli := make([]int, 3)
	var i int = 0
	tmp := make([]int, 3)
	for {
		fmt.Printf("Enter number-> ")
		fmt.Scanln(&buf)
		if buf == "X" {
			return
		} else {
			el, _ = strconv.Atoi(buf)
		}
		if i < 3 {
			tmp[i] = el
			copy(sli, tmp)
		} else {
			sli = append(sli, el)
		}
		i++
		sort.SliceStable(sli, func(i, j int) bool { return sli[i] < sli[j] })
		fmt.Println(sli)
	}
}

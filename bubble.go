package main

import "fmt"

func Swap(sli []int, i int) {
	tmp := sli[i]
	sli[i] = sli[i+1]
	sli[i+1] = tmp
}

func BubbleSort(sli []int) {
	l := len(sli)
	for j := 0; j < l; j++ {
		for i := 0; i < l-1; i++ {
			if sli[i] > sli[i+1] {
				Swap(sli, i)
			}
		}
	}
}

func main() {
	sli := make([]int, 0)
	for {
		var nb int
		_, err := fmt.Scanln(&nb)
		if err != nil {
			break
		}
		sli = append(sli, nb)
	}
	BubbleSort(sli)
	fmt.Println(sli)
}

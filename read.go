package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Persons struct {
	fname string
	lname string
}

func main() {
	var file_name string
	sli := make([]Persons, 0)
	fmt.Scan(&file_name)
	fd, err := os.Open(file_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		tmp := scanner.Text()
		s := strings.Split(tmp, " ")
		var p Persons
		p.fname = s[0]
		p.lname = s[1]
		sli = append(sli, p)
	}
	for _, s := range sli {
		fmt.Printf("%s %s\n", s.fname, s.lname)
	}
}

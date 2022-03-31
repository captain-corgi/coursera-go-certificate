package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type (
	// Person struct
	Person struct {
		fname string
		lname string
	}
)

func main1() {
	// Define a slice of Person
	var people []Person

	// Read file
	file, err := os.Open("cmd\\week4\\test.txt")
	if err != nil {
		panic(err)
	}
	br := bufio.NewReader(file)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				data := strings.Split(line, " ")
				people = append(people, Person{data[0], data[1]})
				break
			}
		}
		data := strings.Split(line, " ")
		people = append(people, Person{data[0], data[1]})
	}

	// Print slice
	for _, person := range people {
		fmt.Println(person.fname, person.lname)
	}
}

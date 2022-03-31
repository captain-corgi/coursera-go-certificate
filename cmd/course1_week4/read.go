package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {
	var fileName string
	var list = make([]Name, 0)

	fmt.Printf("Enter a filename: ")
	fmt.Scan(&fileName)
	fmt.Println("Opening ", fileName)
	f, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		// Read the file until it encounters a newline character
		row, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		// Split line on space to get the two names per row (first and last)
		names := strings.Split(string(row), " ")
		firstName := truncate(names[0])
		lastName := truncate(names[1])
		currentName := Name{Fname: firstName, Lname: lastName}
		list = append(list, currentName)
	}
	for _, person := range list {
		fmt.Printf("%s %s\n", person.Fname, person.Lname)
	}
}

// Limit fields to 20 characters in length
func truncate(s string) string {
	r := ""
	for i, c := range s {
		if i < 20 {
			if string(c) != "\n" {
				r = r + string(c)
			}
		}
	}
	return r
}

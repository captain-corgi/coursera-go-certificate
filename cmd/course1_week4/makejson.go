// Write a program which prompts the user to first enter a name,
// and then enter an address. Your program should create a map
// and add the name and address to the map using the keys “name” and “address”,
// respectively. Your program should use Marshal() to create a JSON object from the map,
// and then your program should print the JSON object.
package main

import (
	"encoding/json"
	"fmt"
)

func mainMe() {
	var (
		name    string
		address string
		myMap   map[string]string
	)
	myMap = make(map[string]string)

	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	fmt.Println("Enter your address: ")
	fmt.Scanln(&address)

	myMap["name"] = name
	myMap["address"] = address

	jsonData, err := json.Marshal(myMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonData))
}

package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//declare variables for user to enter
	var name string
	var address string

	fmt.Println("Enter your name")
	_, _ = fmt.Scan(&name) //give name's addresss so Scan can store it

	fmt.Println("Enter your address")
	_, _ = fmt.Scan(&address) //give address's addresss so Scan can store it

	//create hash map
	var myMap map[string]string
	myMap = make(map[string]string)
	myMap["name"] = name
	myMap["address"] = address

	// convert to json
	barr, _ := json.Marshal(myMap)
	fmt.Println(string(barr))

}

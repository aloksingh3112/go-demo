package main

import "fmt"

type contacts map[string]string

func main() {
	//var person map[string]string
	// contact := make(map[int]string)
	colors := map[string]string{
		"red":   "gg",
		"blue":  "ss",
		"gress": "ff",
	}

	a := contacts{
		"name": "alok",
		"age":  "10",
	}

	a.printL()

	print(colors)
	//person["name"] = "alok"
	// contact[10] = "998877"
	// fmt.Println(contact)

	// fmt.Println(colors, contact)
}

func print(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

func (c contacts) printL() {
	for key, value := range c {
		fmt.Println(key, value)
	}
}

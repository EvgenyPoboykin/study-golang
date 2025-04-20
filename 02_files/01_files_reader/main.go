package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("./db.json")
	if err != nil {
		fmt.Println("Could read the file")
	}
	fmt.Println(string(data))

	var items Items
	JsonFileReader("./db.json", &items)

	fmt.Println(items)
	fmt.Println(items.Items[0])
	fmt.Println(items.Items[0].ClientName)
}

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

	fmt.Println("before", items)

	newItem := Subject{
		MasterId:        "123456",
		AgreementNumber: "55/66",
		ActivationDate:  "2025-04-22T11:30:12",
		Inn:             "456789123",
		ClientName:      "More then I",
	}

	newItems := append(items.Items, newItem)

	items.Items = newItems

	fmt.Println("after", items)

	JsonFileWriter("./new-db.json", items)
}

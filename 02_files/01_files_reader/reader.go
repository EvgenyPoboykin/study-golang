package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func JsonFileReader[O any](filePath string, contentStruct O) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Could read the file")
	}

	fmt.Println(string(data))

	json.Unmarshal(data, contentStruct)
}

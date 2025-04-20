package main

import "fmt"

type Item struct {
	Label string
	Value string
}

func main() {
	arr := []Item{
		{"Work", "outside_home"},
		{"Home", "inside_home"},
	}

	fmt.Println(arr)

	obj1 := Item{}

	for _, v := range arr {
		if v.Value == "inside_home" {
			obj1.Label = v.Label
			obj1.Value = v.Value
			break
		}
	}

	fmt.Println(obj1)

	newArr := append(arr, Item{"Serbia", "new_home"})

	fmt.Println(newArr)
}

package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr)

	fmt.Println(arr[2:3])

	arrString := [3]string{"My", "Name", "Is"}

	fmt.Println(strings.Join(arrString[:], " "))

	usizeArray := [...]int{1, 2, 3}

	fmt.Println(usizeArray)
}

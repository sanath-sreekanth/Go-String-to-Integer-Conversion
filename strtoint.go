package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Parsing in base 10 (decimal)
	str1 := "123"
	int1, err := strconv.ParseInt(str1, 10, 64)
	if err != nil {
		fmt.Println("Error parsing str1:", err)
	} else {
		fmt.Println("Parsed int1:", int1)
	}

	// Parsing in base 2 (binary)
	str2 := "1101"
	int2, err := strconv.ParseInt(str2, 2, 64)
	if err != nil {
		fmt.Println("Error parsing str2:", err)
	} else {
		fmt.Println("Parsed int2:", int2)
	}
}

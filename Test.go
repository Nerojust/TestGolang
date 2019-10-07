package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	println("Hello this is a test project")
	fmt.Print("hey there")
	const pi float64 = 6.5500005005
	name := "card"
	fmt.Print(name + "\n")

	x := 16
	y := 3
	sum := x + y

	if sum > 5 {
		fmt.Print("above five")
	} else if sum < 5 {
		fmt.Print("less than five")
	}

	dd := []int{1, 4, 5, 6, 3, 7}
	//dd[4]=8

	fmt.Print(dd)
	myMap := make(map[string]int)

	myMap["triangle"] = 3
	myMap["circle"] = 5
	myMap["sphere"] = 8
	delete(myMap, "sphere")
	fmt.Print(myMap)
	//for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//while loop
	i := 0

	for i < 8 {
		fmt.Println(i)
		i++

	}
	arr := make(map[string]string)
	//arr:= []string{"a","b","c"}
	arr["car"] = "bmw"
	arr["dar"] = "crv"

	for key, value := range arr {
		fmt.Println("key: ", key, "value: ", value)
	}
	squareResult, error := squareRoot(81)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(squareResult)
	}
	result := addValues(84, 98)
	fmt.Println(result)
}
func addValues(a int, b int) int {
	return a + b
}
func squareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("undefined for negative numbers")
	}
	return math.Sqrt(a), nil
}

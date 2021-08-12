package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"rsc.io/quote"
)

type person struct {
	name string
	age  int
}

func main() {
	y := 7
	var x int = 5
	var array5 [5]int
	array5init := [5]int{5, 4, 3, 4, 2}
	mapInt := createMap()
	mapString := createMapString()
	var arrSlice = arraySlice()
	p := person{name: "Lucas", age: 27}

	fmt.Println(quote.Hello())

	forLoop()
	enumerateArray(arrSlice)
	enumerateMap(mapString)
	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(p)
	fmt.Println(sum(1, 4))
	fmt.Println(mapInt)
	fmt.Println(arrSlice)
	fmt.Println(checkOdd(y))
	fmt.Println(array5)
	fmt.Println(array5init)
	fmt.Println(x + y)
}

func getRandomNumber() int {
	return rand.Intn(10)
}

func sum(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
}

func checkOdd(x int) bool {
	return x%2 == 0
}

func arraySlice() []int {
	return []int{1, 1, 1}
}

func createMap() map[string]int {
	vertices := make(map[string]int)
	vertices["triange"] = 2
	vertices["square"] = 3
	return vertices
}

func createMapString() map[string]string {
	alphabet := make(map[string]string)
	alphabet["a"] = "alpha"
	alphabet["b"] = "beta"

	return alphabet
}

func forLoop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func enumerateArray(arr []int) {
	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
	}
}

func enumerateMap(mapIt map[string]string) {
	for key, value := range mapIt {
		fmt.Println("key:", key, "value:", value)
	}
}

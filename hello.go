package main

import (
	"fmt"

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
	result, err := mathutil.sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(p)
	fmt.Println(mathutil.sum(2, 3))
	fmt.Println(mapInt)
	fmt.Println(arrSlice)
	fmt.Println(mathutil.checkOdd(y))
	fmt.Println(array5)
	fmt.Println(array5init)
	fmt.Println(x + y)
	fmt.Println(strutil.reverse("hello"))
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

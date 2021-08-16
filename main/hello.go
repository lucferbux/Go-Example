package main

import (
	"fmt"
	"log"

	"example.com/error"
	"example.com/mathutil"
	"example.com/strutil"
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
	result, err := mathutil.Sqrt(16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	fmt.Println(p)
	fmt.Println(mathutil.Sum(2, 3))
	fmt.Println(mapInt)
	fmt.Println(arrSlice)
	fmt.Println(mathutil.CheckOdd(y))
	fmt.Println(array5)
	fmt.Println(array5init)
	fmt.Println(x + y)
	fmt.Println(strutil.Reverse("hello"))

	fmt.Println("+++++++++++++")

	message, err := error.Hello("Lucas")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	fmt.Println(message)
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

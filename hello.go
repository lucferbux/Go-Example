package main

import (
	"fmt"
)

func main() {
	y := 7
	var x int = 5
	var array5 [5]int
	array5init := [5]int{5, 4, 3, 4, 2}

	var arrSlice = arraySlice()

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for index, value := range arrSlice {
		fmt.Println("index:", index, "value:", value)
	}

	fmt.Println(arrSlice)
	fmt.Println(checkOdd(y))
	fmt.Println(array5)
	fmt.Println(array5init)
	fmt.Println(x + y)
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

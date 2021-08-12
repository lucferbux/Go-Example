package mathutil

import (
	"math/rand"
)

func GetRandomNumber() int {
	return rand.Intn(10)
}

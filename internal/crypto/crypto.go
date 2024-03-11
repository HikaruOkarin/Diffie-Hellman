package crypto

import (
	"fmt"
	"math"
)

func Ceaser(word string, key int) []int {
	alphanumeric, digits := Digitize(word)
	encryption := []int{}
	for _, val := range word {
		encryption = append(encryption, (alphanumeric[val]+key)%26)
	}

	fmt.Println("digits:", digits)
	return encryption
}

func Digitize(word string) (map[rune]int, []int) {
	alphanumeric := make(map[rune]int)
	count := 0
	for i := 'a'; i <= 'z'; i++ {
		alphanumeric[i] = count
		count++
	}
	Digits := []int{}
	for _, val := range word {
		Digits = append(Digits, alphanumeric[val])
	}
	return alphanumeric, Digits
}

func ToDecimal(word string) int {
	alphanumeric := make(map[rune]int)
	count := 0
	for i := 'a'; i <= 'z'; i++ {
		alphanumeric[i] = count
		count++
	}
	n := 26
	var result int = 0
	degree := len(word) - 1
	for _, val := range word {

		result += alphanumeric[val] * int(math.Pow(float64(n), float64(degree)))
		degree--
	}
	return result
}

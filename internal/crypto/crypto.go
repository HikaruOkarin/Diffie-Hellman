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
	letter := 'a'
	alpha := make(map[int]rune)
	for i := 0; i < 26; i++ {
		alpha[i] = letter
		letter++
	}
	enc_word := ""
	for _, val := range encryption {
		enc_word += string(alpha[val])
	}
	fmt.Println("encryption:", enc_word)
	dec_word := ""
	for _, val := range encryption {
		sum := (val - key) % 26

		if sum < 0 {
			sum += 26
		}
		dec_word += string(alpha[(sum)])
	}
	fmt.Println("decryption:", dec_word)

	fmt.Println("digits:", digits)
	return encryption
}

func Affine(word string, a, b int) []int {
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

func Gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return Gcd(b, a%b)
	}
}

package crypto

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func Ceaser(word string, key int) []int {
	alphanumeric, _ := Digitize(word)
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
	fmt.Print("encryption number:")
	return encryption
}

func Affine(word string, a, b int) ([]int, error) {
	if Gcd(a, 26) != 1 {
		return []int{}, fmt.Errorf("Some error happend")
	}
	alphanumeric, _ := Digitize(word)
	encryption := []int{}
	for _, val := range word {
		encryption = append(encryption, ((alphanumeric[val])*a+b)%26)
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
	modinverse := ModInverse(a, 26)
	dec_word := ""
	for _, val := range encryption {
		sum := val - b
		if sum < 0 {
			sum += 26
		}
		dec_word += string(alpha[(modinverse*sum)%26])

	}
	fmt.Println("decryption:", dec_word)
	fmt.Print("encryption number:")
	return encryption, nil
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

func ModInverse(a, m int) int {
	m0 := m
	t, q := 0, 0
	x0, x1 := 0, 1
	if m == 1 {
		return 0
	}
	for a > 1 {
		q = a / m
		t = m
		m = a % m
		a = t
		t = x0
		x0 = x1 - q*x0
		x1 = t
	}
	if x1 < 0 {
		x1 += m0
	}
	return x1
}

func Cipher() {
	fmt.Println("type cipher:")
	fmt.Println("1.Ceaser")
	fmt.Println("2.Affine")
	fmt.Println("3.Diffie-Hellman")
	var cipher int
	fmt.Scan(&cipher)
	switch cipher {
	case 1:
		var word string
		var key int
		fmt.Print("word:")
		fmt.Scan(&word)
		fmt.Print("key:")
		fmt.Scan(&key)
		fmt.Println(Ceaser(word, key))
	case 2:
		var word string
		var a, b int
		fmt.Print("word:")
		fmt.Scan(&word)
		for Gcd(a, 26) != 1 {
			a = rand.Intn(25)
		}
		fmt.Println("a:", a)
		b = rand.Intn(15)
		fmt.Println("b:", b)
		crypto, _ := (Affine(word, a, b))
		fmt.Println(crypto)
	case 3:
		Diffie()
	}
}

func Diffie() {
	fmt.Println("Alice:")
	fmt.Println("Public keys:")
	var p, g, a, b int
	p = FindPrime()

	fmt.Println("p:", p)
	rand.Seed(time.Now().UnixNano())
	g = rand.Intn(19) + 2
	fmt.Println("g:", g)
	fmt.Print("a:")
	fmt.Scan(&a)
	fmt.Println("private key")
	A := modExp(g, a, p)
	fmt.Println("key A transfer for Bob->", A)
	fmt.Println("Bob:")
	fmt.Println("private key")
	fmt.Print("b:")
	fmt.Scan(&b)

	B := modExp(g, b, p)
	fmt.Println("key B transfer for Alice->", B)

	k_a := modExp(B, a, p)
	k_b := modExp(A, b, p)
	fmt.Println("Shared secret key computed by Alice:", k_a)
	fmt.Println("Shared secret key computed by Bob:", k_b)
}

func modExp(base, exponent, modulus int) int {
	result := 1
	base = base % modulus
	for exponent > 0 {
		if exponent%2 == 1 {
			result = (result * base) % modulus
		}
		exponent = exponent >> 1
		base = (base * base) % modulus
	}
	return result
}

func Isprime(nb int) bool {
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindPrime() int {
	num := rand.Intn(1000) + 3
	for num > 2 {
		if Isprime(num) {
			return num
		}
		num--
	}
	return 2
}

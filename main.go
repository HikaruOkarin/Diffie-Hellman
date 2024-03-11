package main

import (
	"Diffie-Hellman/internal/crypto"
	"fmt"
)

func main() {
	info := crypto.Ceaser("didar", 3)
	fmt.Println(info)
}

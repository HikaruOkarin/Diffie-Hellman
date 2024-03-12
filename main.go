package main

import (
	"Diffie-Hellman/internal/crypto"
	"Diffie-Hellman/internal/server"
)

func main() {
	crypto.Cipher()
	server.Runserver()
}

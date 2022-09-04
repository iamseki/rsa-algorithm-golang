package main

import "fmt"

func main() {
	message := []byte("The information security is of significant importance to ensure the privacy of communications")

	rsa := newRSAFromScratch(2048)
	encrypted := rsa.Encrypt(message)
	fmt.Printf("encrypted message => %s \n", encrypted)

	decrypted := rsa.Decrypt()
	fmt.Printf("decrypted message => %s \n", decrypted)
}

package main

import (
	"bytes"
	"testing"
)

func Test64BitRsaFromScratch(t *testing.T) {
	shortMsg := []byte("Hello World")

	rsa := newRSAFromScratch(64)

	rsa.Encrypt(shortMsg)
	decrypted := rsa.Decrypt()

	if !bytes.Equal(shortMsg, decrypted) {
		t.Errorf("64 bit RSA algorithm error, expect decrypted message to be: %v but got %v", shortMsg, decrypted)
	}
}

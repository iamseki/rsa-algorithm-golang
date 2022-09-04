package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"sync"
)

type RSAFromScratch struct {
	p          *big.Int
	q          *big.Int
	n          *big.Int
	m          *big.Int
	e          *big.Int
	d          *big.Int
	PublicKey  *big.Int
	PrivateKey *big.Int
}

// Encrypt, encrypts the message using rsa algorithm
// returning encrypted "public key" string representation
func (rsa *RSAFromScratch) Encrypt(message []byte) string {
	// get integer message representation "P"
	P := big.NewInt(0).SetBytes(message)

	// public key => P^e MOD n , aka RSA Encrypt
	rsa.PublicKey = big.NewInt(0).Exp(P, rsa.e, rsa.n)

	return rsa.PublicKey.String()
}

// Decrypt, decrypts the message using rsa algorithm
// returning its original data
func (rsa *RSAFromScratch) Decrypt() []byte {
	// private key => encryptedInt^d MOD n , aka RSA Decrypt
	rsa.PrivateKey = big.NewInt(0).Exp(rsa.PublicKey, rsa.d, rsa.n)
	return rsa.PrivateKey.Bytes()
}

// generate2048BitsPrimeNumbers, returns two 2048 bits random prime numbers
// this execution is made with concurrency possibly in parallel
func generatePrimeNumbers(bitPrimeNumberSize int) (*big.Int, *big.Int) {
	var wg sync.WaitGroup
	var p, q *big.Int

	wg.Add(2)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		n, err := rand.Prime(rand.Reader, bitPrimeNumberSize)
		if err != nil {
			log.Fatalln("err on genereate first random prime number: ", err)
		}

		p = n
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		n, err := rand.Prime(rand.Reader, bitPrimeNumberSize)
		if err != nil {
			log.Fatalln("err on genereate second random prime number: ", err)
		}

		q = n
	}(&wg)

	wg.Wait()

	return p, q
}

// calculateTotiente, phi(n) = (p-1) * (q-1)
func calculateTotiente(p, q *big.Int) *big.Int {
	// this is (p-1)
	n1 := big.NewInt(0).Sub(p, big.NewInt(1))
	// this is (q-1)
	n2 := big.NewInt(0).Sub(q, big.NewInt(1))

	return big.NewInt(0).Mul(n1, n2)
}

func newRSAFromScratch(bitPrimeNumberSize int) *RSAFromScratch {
	// generate random prime numbers, p and q with size bits = bitPrimeNumberSize
	p, q := generatePrimeNumbers(bitPrimeNumberSize)

	// compute n = p * q
	n := big.NewInt(0).Mul(p, q)

	// phi(n) = (p-1) * (q-1)
	m := calculateTotiente(p, q)

	// choose e => 1 < e < phi(n) | e and phi(n) prime numbers between them selfs
	e := big.NewInt(3)
	// Cmp returning 1 means x > y, so it will sum on "e" until gcd < 1
	for big.NewInt(0).GCD(nil, nil, m, e).Cmp(big.NewInt(1)) == 1 {
		e = big.NewInt(0).Add(e, big.NewInt(2))
	}

	// d must be inverse multiplier of "e"
	d := big.NewInt(0).ModInverse(e, m)

	return &RSAFromScratch{
		p: p, q: q, n: n, m: m, e: e, d: d,
	}
}

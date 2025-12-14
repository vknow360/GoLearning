package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"
	// hash := sha256.Sum256([]byte(password))
	// fmt.Println(hash)
	// fmt.Printf("SHA-256 hash hex val: %x\n", hash)

	// hash512 := sha512.Sum512([]byte(password))
	// fmt.Println(hash512)
	// fmt.Printf("SHA-512 hash hex val: %x\n", hash512)

	salt, err := generateSalt()
	if err != nil {
		fmt.Println(err)
		return
	}

	signUpHash := hashPassword(password, salt)
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println(saltStr)
	fmt.Println(signUpHash)

	// verify
	// retrieve saltstr and decode it
	decodedSalt, _ := base64.StdEncoding.DecodeString(saltStr)
	loginHash := hashPassword("password123", decodedSalt)

	// compare
	if signUpHash == loginHash {
		fmt.Println("Correct password")
	} else {
		fmt.Println("Login failed")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}

package main

import (
	"crypto/rand"
	"crypto/sha256"
	// "crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	fmt.Println("----- HASHING -----")

	password := "password1234"
	
	// hash := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))

	// fmt.Println(hash)
	// fmt.Println(hash512)
	// //convert hash to hex
	// fmt.Printf("hash converted to hex value: %x\n", hash)
	// fmt.Printf("hash converted to hex value: %x\n", hash512)

	//Salting: adding some string just before hashing a password.
	salt, err := generateSalt()
	fmt.Println("printing generated salt:", salt)
	if err != nil {
		fmt.Println("error occured:", err)
		return
	}
	//hash the password with salt.
	signupHash := hashPassword(password, salt)

	//store the salt and password in database, just printing as of now.
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Salt:", saltStr)
	fmt.Println("hash:", signupHash)

	//verify
	loginPassword := "password134"
	//retrieve the saltStr and decode it.
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error while decoding salt:", err)
		return
	}
	fmt.Println("Decoded Salt:", decodedSalt)
	loginHash := hashPassword(loginPassword, decodedSalt)
	if loginHash == signupHash {
		fmt.Println("Password is correct, You are logged in!")
	}else {
		fmt.Println("Incorrect Password, Try Again!")
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

//function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}

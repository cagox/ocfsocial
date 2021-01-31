package crypto

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

//HashPassword returns a hashed version of the password.
func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		//TODO: Do proper error handling
		log.Fatal(err)
	}
	return string(hash)
}

//ComparePassword tells us if the password matches
func ComparePassword(password string, databaseHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(databaseHash), []byte(password))
	if err != nil {
		//fmt.Println(err.Error())
		return false
	}
	return true
}

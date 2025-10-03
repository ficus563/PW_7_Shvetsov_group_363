package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type user struct {
	username string
	email    string
	password []byte
}

func (u *user) SetPassword(password string) {
	h := sha256.Sum256([]byte(password))
	u.password = h[:]
}

func (u *user) verify_passwd(password string) bool {
	var input_hash = sha256.Sum256([]byte(password))
	return hex.EncodeToString(input_hash[:]) == hex.EncodeToString(u.password)
}

func main() {
	testUser := user{"test_userok", "testik@example.com", nil}
	testUser.SetPassword("123")
	if testUser.verify_passwd("wrong_password") {
		fmt.Println("Пароль верный.")
	} else {
		fmt.Println("Неверный пароль.")
	}

	if testUser.verify_passwd("password123") {
		fmt.Println("Пароль правильный!")
	} else {
		fmt.Println("Ошибка верификации.")
	}
}

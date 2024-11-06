package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	email    string
	username string
	age      int
}

func (u User) Email() string {
	return u.email
}

func Email(u User) string {
	return u.email
}

func (u *User) updateEmail(email string) {
	u.email = email

	fmt.Println("Email updated to", u.email)
}

func UpdateEmail(u *User, email string) {
	u.email = email
}

func main() {
	user := User{
		email: "abc@foo.com",
	}

	user.updateEmail("asdas#@sd")
	UpdateEmail(&user, "dasfsa.asdaf")

	fmt.Println(user.Email(), unsafe.Sizeof(user))
}

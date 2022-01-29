package authentication

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

var UserList []*user

func CreateUser(userId, password string) {
	salt := createSalt()
	newPassword := password + salt
	hashPassword := createHash(newPassword)
	newUser := NewUser(userId, hashPassword, salt)
	UserList = append(UserList, newUser)
	fmt.Println("Your account has been created user: ", userId, "\nPassword stored as: ", hashPassword)

}

func LoginUser() string {
	var userId, password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&userId)
	fmt.Print("Enter Password: ")
	fmt.Scanln(&password)
	for _, singleUser := range UserList {
		if userId == singleUser.id {
			salt := singleUser.salt
			pass := password + salt
			hashPass := createHash(pass)
			status := checkPassword(hashPass, singleUser.password)
			return status

		}
	}
	return "Invalid username and password"

}

func checkPassword(enteredPass, storedPass string) string {
	if enteredPass == storedPass {
		return "Login Successfull"
	}
	return "Incorrect Password"
}

func createSalt() string {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	var numberRunes = []rune("1234567890")
	b := make([]rune, 5)
	for i := range b {
		b[i] = numberRunes[rand.Intn(len(numberRunes))]
	}
	return string(b)

}

func createHash(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd)) //storing key in slice of bytes
	return hex.EncodeToString(hasher.Sum(nil))
}

func NewUser(userId string, password string, salt string) *user {
	var user = &user{
		id:       userId,
		password: password,
		salt:     salt,
	}
	return user
}

type user struct {
	id       string
	password string
	salt     string //random data generated that will be used for salting
}

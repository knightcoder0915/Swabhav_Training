package authentication

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var UserList []*user

func CreateUser(userId, password string) {
	salt := createSalt()
	saltHash := createHash(salt)
	newPassword := password + saltHash
	hashPassword := createHash(newPassword)
	newUser := NewUser(userId, hashPassword, salt)
	UserList = append(UserList, newUser)
	fmt.Println("Your account has been created user: ", userId, "\nPassword stored as: ", hashPassword)

}

func LoginUser() string {
	var userId, password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&userId)
	for _, singleUser := range UserList {
		if userId == singleUser.id {
			fmt.Print("Enter Password: ")
			fmt.Scanln(&password)
			salt := singleUser.salt
			saltHash := createHash(salt)
			pass := password + saltHash
			hashPass := createHash(pass)
			status := checkPassword(hashPass, singleUser.password)
			return status

		}
	}
	return "Invalid username and password"

}

func ListAllUsers() {
	for _, singleUser := range UserList {
		fmt.Println("User name: ", singleUser.id)
	}
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

func CheckIfUserExists(username string) error {
	if UserList == nil {
		return nil

	} else {
		for _, singleUser := range UserList {
			if username != singleUser.id {
				return nil
			}
		}
	}

	return errors.New("Username exists")

}

func ReadData() {
	file, err := os.Open("userList.txt")
	if err != nil {
		//fmt.Println("File reading error", err)
		return
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text, data []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	for _, eachline := range text {
		data = append(data, eachline)
	}
	if data == nil {

	}
	mySlice := make([]string, 0)

	for _, detail := range data {
		mySlice = strings.Split(string(detail), " ") //creating a slice inorder to access each element
		username := mySlice[0]
		password := mySlice[1]
		randomNo := mySlice[3]
		newUser := NewUser(username, password, randomNo)
		UserList = append(UserList, newUser)

	}

}

func WriteData() {
	f, _ := os.Create("userList.txt") //it will return an os pointer and error

	defer f.Close()
	for _, singlePerson := range UserList {
		_, e := f.WriteString(singlePerson.id + " " + singlePerson.password + " " + singlePerson.salt + "\n")
		if e != nil {
			fmt.Println("Error occured")
			return
		}
	}

	fmt.Println("Done Writing")
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

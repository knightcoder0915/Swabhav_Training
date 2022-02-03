package authentication

import (
	"bufio"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var UserList []*User

func CreateUsers(fullName, password string, levelBiba int, levelBell int) {
	id := CreateRandomId()
	salt := createSalt()
	saltHash := createHash(salt)
	newPassword := password + saltHash
	hashPassword := createHash(newPassword)
	newUser := NewUser(fullName, id, hashPassword, salt, levelBiba, levelBell)
	UserList = append(UserList, newUser)
	fmt.Println("Your account has been created user: ", id, "\nPassword stored as: ", hashPassword)

}

func LoginUser() (string, *User) {
	var userId, password string
	fmt.Print("Enter id: ")
	fmt.Scanln(&userId)
	for _, singleUser := range UserList {
		if userId == singleUser.Id {
			fmt.Print("Enter Password: ")
			fmt.Scanln(&password)
			salt := singleUser.salt
			saltHash := createHash(salt)
			pass := password + saltHash
			hashPass := createHash(pass)
			status := checkPassword(hashPass, singleUser.password)
			return status, singleUser

		}
	}
	return "Invalid username and password", nil

}

func ListAllUsers() {
	for _, singleUser := range UserList {
		fmt.Println("User name: ", singleUser.Id)
	}
}

func checkPassword(enteredPass, storedPass string) string {
	if enteredPass == storedPass {
		return "Login Successfull"
	}
	return "Incorrect Password"
}

func createSalt() string {
	RandomCrypto, _ := rand.Prime(rand.Reader, 15)
	id := RandomCrypto.String()
	return id

}

func createHash(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd)) //storing key in slice of bytes
	return hex.EncodeToString(hasher.Sum(nil))
}

func CheckIfUserExists(fullName string) error {
	if UserList == nil {
		return nil

	} else {
		for _, singleUser := range UserList {
			if fullName != singleUser.fullName {
				return nil
			}
		}
	}

	return errors.New("Username exists")

}

func CreateRandomId() string {
	RandomCrypto, _ := rand.Prime(rand.Reader, 10)
	id := RandomCrypto.String()
	return id

}

func ReadUserData() {
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
		fullName := mySlice[0]
		id := mySlice[1]
		password := mySlice[2]
		randomNo := mySlice[3]
		levelBiba, _ := strconv.Atoi(mySlice[4])
		levelBellaLaPadula, _ := strconv.Atoi(mySlice[5])
		newUser := NewUser(fullName, id, password, randomNo, levelBiba, levelBellaLaPadula)
		UserList = append(UserList, newUser)

	}

}

func WriteUserData() {
	f, _ := os.Create("userList.txt") //it will return an os pointer and error

	defer f.Close()
	for _, singlePerson := range UserList {
		level1 := strconv.Itoa(singlePerson.LevelBiba)
		level2 := strconv.Itoa(singlePerson.LevelBellLaPadula)
		_, e := f.WriteString(singlePerson.fullName + " " + singlePerson.Id + " " + singlePerson.password + " " + singlePerson.salt + " " + level1 + " " + level2 + "\n")
		if e != nil {
			fmt.Println("Error occured")
			return
		}
	}

	fmt.Println("Done Writing")
}

func NewUser(fullName, userId string, password string, salt string, levelBiba int, levelBellaLaPadula int) *User {
	var user = &User{
		fullName:          fullName,
		Id:                userId,
		password:          password,
		salt:              salt,
		LevelBiba:         levelBiba,
		LevelBellLaPadula: levelBellaLaPadula,
	}
	return user
}

type User struct {
	fullName          string
	Id                string
	password          string
	salt              string //random data generated that will be used for salting
	LevelBiba         int
	LevelBellLaPadula int
}

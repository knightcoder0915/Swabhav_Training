package main

import (
	"bufio"
	"example/Day6/Task/aescipher"
	"example/Day6/Task/authentication"
	"example/Day6/Task/file"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Files []*file.File

func main() {
	var choice, level1, level2 int
	var fName, lName, password string
	authentication.ReadUserData()
	file1 := file.NewFile("data/roles.txt", 3, 1) //first is biba and second is bell
	FileAdd(file1)
	file2 := file.NewFile("data/sales.txt", 3, 2)
	FileAdd(file2)
	file3 := file.NewFile("data/dev.txt", 3, 1)
	FileAdd(file3)
	// fmt.Println("Display File List in System")
	// for _, singleFile := range Files {
	// 	fmt.Println(*&singleFile.Name)
	// }

	for {
		fmt.Println("Welcome!1-->Register New User,2-->Login,3-->List Users,4-->Exit App \nEnter your Choice:")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter First Name")
			fmt.Scanln(&fName)
			fmt.Println("Enter Last Name")
			fmt.Scanln(&lName)
			err := authentication.CheckIfUserExists(fName + lName)
			if err != nil {
				fmt.Println(err)
				break
			}
		begin:
			fmt.Println("Enter password")
			fmt.Scanln(&password)
			i := len(password)
			if i == 8 {
				fmt.Print("Enter Biba Level: ")
				fmt.Scanln(&level1)
				fmt.Print("Enter Bell La Padula Level: ")
				fmt.Scanln(&level2)
				fullName := fName + lName
				authentication.CreateUsers(fullName, password, level1, level2)
				break
			}
			fmt.Println("Password should be of length 8")
			goto begin

		case 2:
			status, user := authentication.LoginUser()
			var ch int
			if status == "Login Successfull" {
			start:
				fmt.Println("Choose Mode 1-->Read,2-->Write 3-->Logout")
				fmt.Scanln(&ch)
				if ch == 1 {
					mode := "Read"
					ImplementAAA(user, mode)
					goto start

				} else if ch == 2 {
					mode := "Write"
					ImplementAAA(user, mode)
					goto start
				} else {
					break
				}

			} else {
				fmt.Println(status)
			}
		case 3:
			authentication.ListAllUsers()

		}
		if choice == 4 {
			authentication.WriteUserData()
			break
		}
	}
}

func FileAdd(file *file.File) []*file.File {
	Files = append(Files, file)
	return Files
}

func ImplementAAA(user *authentication.User, mode string) {
	var ch1 int
	var count int = 0
	userLevelBiba := *&user.LevelBiba
	userLevelBell := *&user.LevelBellLaPadula
	// fmt.Println("Enter which mode you want to use: 1-->Read 2-->Write")
	// fmt.Scanln(&ch)
	if mode == "Read" {
		count = 0
		fmt.Println("These are the files you can Read")
		for i, singleFile := range Files {
			if singleFile.AccessLevelBiba >= userLevelBiba && singleFile.AccessLevelBellLaPadula <= userLevelBell {
				fmt.Println(i, "->", singleFile.Name)
				count++
			}
		}
		if count == 0 {
			fmt.Println("Sorry you don't have access to any files")
			return
		}

		fmt.Println("Choose File to Read")
		fmt.Scanln(&ch1)
		for i, singleFile := range Files {
			if i == ch1 {
				Read(singleFile.Name)
				break
			}
		}
		//fmt.Println("Invalid choice")

	} else {
		count = 0
		fmt.Println("These are the files you can Write")
		for i, singleFile := range Files {
			if singleFile.AccessLevelBiba <= userLevelBiba && singleFile.AccessLevelBellLaPadula >= userLevelBell {
				fmt.Println(i, "->", singleFile.Name)
				count++
			}
		}
		if count == 0 {
			fmt.Println("Sorry you don't have access to any files")
			return
		}
		fmt.Println("Choose File to Write")
		fmt.Scanln(&ch1)
		for i, singleFile := range Files {
			if i == ch1 {
				Write(singleFile.Name)
				break

			}
			//fmt.Println("Invalid choice")
		}

	}

}

func Write(fileName string) {
	//var data string
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Enter data to be written")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	cipherText := aescipher.Encrypt([]byte(input), "hello")
	//fmt.Println("Cipher", cipherText)
	if _, err = f.WriteString(string(cipherText) + "line"); err != nil {
		panic(err)
	}
	fmt.Println("Successfully Written")
}

func Read(fileName string) {
	f, err := os.Open(fileName)
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := ioutil.ReadFile(fileName)
	fmt.Println("Contents of File")
	res1 := strings.Split(string(data), "line")
	for i := 0; i < len(res1)-1; i++ {
		fmt.Print(string(aescipher.Decrypt([]byte(res1[i]), "hello")))
	}
	fmt.Scanln()

}

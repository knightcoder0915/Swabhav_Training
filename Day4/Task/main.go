package main

import (
	"example/Day4/Task/authentication"
	"fmt"
)

func main() {
	var choice int
	var userName, password string
	authentication.ReadData()
	for {
		fmt.Println("Welcome!1-->Register New User,2-->Login,3-->List Users,4-->Exit App \nEnter your Choice:")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter username")
			fmt.Scanln(&userName)
			err := authentication.CheckIfUserExists(userName)
			if err != nil {
				fmt.Println(err)
				break
			}
		begin:
			fmt.Println("Enter password")
			fmt.Scanln(&password)
			i := len(password)
			if i == 8 {
				authentication.CreateUser(userName, password)
				break
			}
			fmt.Println("Password should be of length 8")
			goto begin

		case 2:
			status := authentication.LoginUser()
			if status == "Login Successfull" {
				userDisplay()

			} else {
				fmt.Println(status)
			}
		case 3:
			authentication.ListAllUsers()

		}
		if choice == 4 {
			authentication.WriteData()
			break
		}
	}
}

func userDisplay() {
	var choice int
	for {
		fmt.Println("Welcome to Tanvi World....Press 1 to Logout")
		fmt.Scanln(&choice)

		if choice == 1 {
			break
		}
		fmt.Println("Please Choose correct option")
	}
}

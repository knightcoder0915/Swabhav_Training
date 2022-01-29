package main

import (
	"example/Day4/Task/authentication"
	"fmt"
)

func main() {
	var choice int
	var userName, password string
	for {
		fmt.Println("Welcome!1-->Register New User,2-->Login,3-->Exit App \nEnter your Choice:")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Enter username")
			fmt.Scanln(&userName)
			fmt.Println("Enter password")
			fmt.Scanln(&password)
			authentication.CreateUser(userName, password)
		case 2:
			status := authentication.LoginUser()
			if status == "Login Successfull" {
				userDisplay()

			} else {
				fmt.Println(status)
			}

		}
		if choice == 3 {
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

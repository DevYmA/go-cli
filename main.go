package main

import (
	"fmt"
	"strings"
)

func main() {

	var applicationName string = "Go CLI"
	const messageCount int = 30

	var allGiftCount uint = 80
	var givenGiftCount uint = 10

	var winnerList []string

	fmt.Printf("Welcome to %v 🌝 \n", applicationName)

	for givenGiftCount != allGiftCount {

		var remainingGift uint = allGiftCount - givenGiftCount
		fmt.Println("-------------------------------------------------------------")

		fmt.Printf("We have %v availble gift. Hurry up \n \n", remainingGift)

		fmt.Println("Before rollout to hand shake enter your details 🚀")

		var firstName string
		var lastName string
		var email string
		var age int

		isValidFullName := len(firstName) > 4 && len(lastName) > 4
		isValidEmail := strings.Contains(email, "@")

		if !isValidFullName && !isValidEmail {
			fmt.Println("Your input data is invalid, please try again")
			continue
		}

		if age < 20 {
			fmt.Println("Seems like your are adult 🧒 ")
		} else {
			fmt.Println("Seems like your are adult 🧑 ")
			fmt.Print("Sorry, Your are not eligible for this giveaway")
			break
		}

		fmt.Println("Please enter your First Name :")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your Last Name :")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your Email :")
		fmt.Scan(&email)

		fmt.Println("Please enter your age ")
		fmt.Scan(&age)

		fmt.Println("Entered details are following ")
		fmt.Printf("Your name is %v %v , email  %v and Your age is %v \n", firstName, lastName, email, age)

		givenGiftCount = givenGiftCount + 1
		winnerList = append(winnerList, email)

		fmt.Printf("We have %v gift available right now. One of them are yours", remainingGift)
		fmt.Printf("Gift Winners Emails %v \n ", winnerList)
	}

	fmt.Println("Currently all available gift are 0. Please join with us at next round ")
}

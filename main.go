package main

import (
	"fmt"
	"go-cli/regional"
)

const messageCount int = 30

var applicationName string = "Go CLI"
var allGiftCount uint = 80
var givenGiftCount uint = 10
var winnerList = make([]UserDetails, 0)

type UserDetails struct {
	firstName string
	lastName  string
	email     string
	age       uint
	region    string
}

func main() {

	fmt.Printf("Welcome to %v 🌝 \n", applicationName)

	for givenGiftCount != allGiftCount {

		var firstName string
		var lastName string
		var email string
		var age uint
		var region string

		var remainingGift uint = allGiftCount - givenGiftCount
		fmt.Println("-------------------------------------------------------------")

		fmt.Printf("We have %v availble gift. Hurry up \n \n", remainingGift)

		fmt.Println("Before rollout to hand shake enter your details 🚀")

		fmt.Println("Please enter your First Name :")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your Last Name :")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your Email :")
		fmt.Scan(&email)

		fmt.Println("Please enter your age :")
		fmt.Scan(&age)

		fmt.Println("Please enter your region :")
		fmt.Scan(&region)

		isValidFullName, isValidEmail := validateInputs(firstName, lastName, email)

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

		regionalMsg := regional.GetRegionalGiftMessage(region)
		fmt.Println(regionalMsg)

		fmt.Println("Entered details are following ")
		fmt.Printf("Your name is %v %v , email  %v and Your age is %v \n", firstName, lastName, email, age)

		userData := UserDetails{
			firstName: firstName,
			lastName:  lastName,
			email:     email,
			age:       age,
			region:    region,
		}

		// var userData = make(map[string]string)
		// userData["firstName"] = firstName
		// userData["lastName"] = lastName
		// userData["ëmail"] = email
		// userData["region"] = region
		// userData["age"] = strconv.FormatInt(int64(age), 10)

		givenGiftCount = givenGiftCount + 1
		winnerList = append(winnerList, userData)

		fmt.Printf("We have %v gift available right now. One of them are yours", remainingGift)
		fmt.Printf("Gift Winners Emails %v \n ", winnerList)
		go sendEmailToWinner(firstName, email)
	}

	fmt.Println("Currently all available gift are 0. Please join with us at next round ")
}

func sendEmailToWinner(firstName string, email string) {
	fmt.Println("Generating email")
	details := fmt.Sprintf("%v user name %v email address", firstName, email)
	fmt.Println(details)
	fmt.Println("Email sent")
}

package main

import "fmt"

func main() {

	var applicationName string = "Go CLI"
	const messageCount int = 30

	var allGiftCount uint = 80
	var givenGiftCount uint = 10

	fmt.Printf("Welcome to %v 🌝 \n", applicationName)
	fmt.Println("Before rollout to hand shake enter your name 🚀")

	var userName string
	var userAge int

	fmt.Println("Please enter your good name here ")
	fmt.Scan(&userName)

	fmt.Println("Please enter your age ")
	fmt.Scan(&userAge)

	fmt.Printf("Your name is %v and Your age is %v \n", userName, userAge)
	fmt.Println("We are looking details you have entered. ")

	var remainingGift uint = allGiftCount - givenGiftCount

	fmt.Printf("We have %v gift available right now. One of them are yours", remainingGift)
}

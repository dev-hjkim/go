package main

import "fmt"

func main() {
	var a int = 10
	var msg string = "Hello Variable"

	a = 20
	msg = "Good Morning"
	fmt.Println(msg, a)

	var minimumAge int = 10
	var workingHour int = 20

	var income int = minimumAge * workingHour
	
	fmt.Println(minimumAge, workingHour, income)
}
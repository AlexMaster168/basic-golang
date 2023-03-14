package main

import "log"

func pointer() {
	var myString string = "Green"
	log.Println("Value myString is ", myString)
	changeUsingPointer(&myString)
	log.Println("after func call myString is set to", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is ", s)
	newValue := "Red"
	*s = newValue
}

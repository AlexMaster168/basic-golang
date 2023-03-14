package main

import "fmt"

func variables() {
	var hello string = "Hello world"
	var a int = 12
	fmt.Println(hello, "number is ", a)
	whatWasSaid, elseWasSaid := saySomething()
	fmt.Println("function returned", whatWasSaid, elseWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}

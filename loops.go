package main

import (
	"log"
)

func loops() {
	//for i := 0; i <= 5; i++ {
	//	log.Println(i)
	//}
	//animals := []string{"dog", "cat", "bird", "fish", "elephant"}
	//for _, animal := range animals {
	//	log.Println(animal)
	//}
	//str := "Hello World!"
	//for i, l := range str {
	//	log.Println(i, l)
	//}
	animals := map[string]string{
		"dog":      "Rover",
		"cat":      "Whiskers",
		"bird":     "Tweety",
		"fish":     "Nemo",
		"elephant": "Dumbo",
	}
	for animalType, animalName := range animals {
		log.Println(animalType, animalName)
	}
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@smith.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@smith.com", 17})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}

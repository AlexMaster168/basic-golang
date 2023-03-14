package main

import (
	"log"
	"sort"
)

type UserTwo struct {
	FirstName string
	LastName  string
}

func dataStructures() {
	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
	mySlice := []int{1, 5, 9, 10}
	//var mySlice []string
	mySlice = append(mySlice, 8, 2)
	sort.Ints(mySlice)
	//sort.Slice(mySlice, func(i, j int) bool {
	//	return mySlice[i] > mySlice[j]
	//})
	log.Println(mySlice)
	mySliced := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(mySliced[2:7])
	myMap := make(map[string]UserTwo)
	me := UserTwo{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}
	myMap["me"] = me
	//delete(myMap, "me")
	log.Println(myMap["me"])
	//myMap := map[string]string{"hie": "Hello", "bye": "Goodbye"}
	//myMap["key1"] = "Value"
	//myMap["key2"] = "Hi"
	//log.Println()
	////_, prs := myMap["key"]
	////log.Println(prs) // false
}

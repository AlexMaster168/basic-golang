package main

import (
	"errors"
	"log"
)

func funcForTesting() {
	result, err := divide(1, 5)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result of devision is ", result)
}

func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divide by 0")
	}
	result = x / y
	return result, nil
}

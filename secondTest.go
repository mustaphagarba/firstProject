package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"‎ایند مالس",
	"Привет, мир",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(helloList))
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}
	// Print our message to the console
	fmt.Println(msg)

}
func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		// Create an error, convert the int type to a string
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}

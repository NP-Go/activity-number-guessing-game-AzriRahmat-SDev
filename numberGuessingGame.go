package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int64
	var i int = 0
	//This code creates random number and place as variable
	rand.Seed(time.Now().UnixNano())
	hiddenNumber := int64(rand.Intn(100))
	//Insert code here

	for i = 0; i <= 5; i++ {
		fmt.Println("Please select a number from 0 - 100: ")
		fmt.Scanln(&userInput)

		if userInput < hiddenNumber {
			fmt.Println("The number you have selected is too low please try again!")
		} else if userInput > hiddenNumber {
			fmt.Println("The number you have selected is too high please try again!")
		} else {
			fmt.Println("well done you have selected the correct number")
			break
		}

		if i == 5 {
			fmt.Println("End of guessing game ", hiddenNumber)
			break
		}
	}

}

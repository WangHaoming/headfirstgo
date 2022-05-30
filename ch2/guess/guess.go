package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	result := rand.Intn(100)
	rightAnswer := false
	fmt.Println("Please enter a number between 0 and 100")
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 10; i++ {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > result {
			fmt.Println("Your number is bigger!")
			continue
		} else if guess < result {
			fmt.Println("Your number is smaller!")
			continue
		} else {
			rightAnswer = true
			fmt.Println("you got it!")
			break
		}

	}
	if !rightAnswer {
		fmt.Println("Failed")
	}
	fmt.Println("The resut is ", result)
}

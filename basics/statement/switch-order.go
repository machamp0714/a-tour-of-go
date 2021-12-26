package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Whens Saturday?")
	today := time.Now().Weekday()
	fmt.Printf("Today is %v\n", today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tommorow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far way.")
	}
}

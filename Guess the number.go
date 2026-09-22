package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ANSI Colors
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Cyan   = "\033[36m"
	Blue   = "\033[34m"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var choice int
	var playAgain string

	for {

		// Welcome Screen
		fmt.Println(Cyan + "╔══════════════════════════════════════╗" + Reset)
		fmt.Println(Cyan + "║       🎯 NUMBER GUESSING GAME       ║" + Reset)
		fmt.Println(Cyan + "╚══════════════════════════════════════╝" + Reset)

		fmt.Println()
		fmt.Println(Yellow + "Select Difficulty:" + Reset)
		fmt.Println("1. Easy   (1 - 50)")
		fmt.Println("2. Medium (1 - 100)")
		fmt.Println("3. Hard   (1 - 500)")

		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)

		var maxNumber int
		var attempts int

		switch choice {
		case 1:
			maxNumber = 50
			attempts = 10
		case 2:
			maxNumber = 100
			attempts = 7
		case 3:
			maxNumber = 500
			attempts = 10
		default:
			fmt.Println(Red + "Invalid choice!" + Reset)
			continue
		}

		
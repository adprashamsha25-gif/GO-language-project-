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
		// Generate random number
		secretNumber := rand.Intn(maxNumber) + 1

		score := 100
		won := false

		fmt.Println()
		fmt.Println(Green + "🎮 Game Started!" + Reset)
		fmt.Printf("Guess a number between 1 and %d\n", maxNumber)
		fmt.Printf("You have %d attempts.\n\n", attempts)

		for i := 1; i <= attempts; i++ {

			var guess int

			fmt.Printf("Attempt %d/%d - Enter your guess: ", i, attempts)
			fmt.Scan(&guess)

			if guess < 1 || guess > maxNumber {
				fmt.Println(Red + "⚠️ Please enter a valid number!" + Reset)
				i--
				continue
			}

			if guess == secretNumber {
				won = true

				// Increase/decrease score according to attempts
				score = score - (i-1)*10

				if score < 10 {
					score = 10
				}

				fmt.Println()
				fmt.Println(Green + "🎉 CORRECT! YOU WON! 🎉" + Reset)
				fmt.Println()
				fmt.Println("╔════════════════════════════╗")
				fmt.Println("║       🏆 GAME RESULT       ║")
				fmt.Println("╠════════════════════════════╣")
				fmt.Printf("║ Number: %-17d ║\n", secretNumber)
				fmt.Printf("║ Attempts: %-15d ║\n", i)
				fmt.Printf("║ Score: %-18d ║\n", score)
				fmt.Println("╚════════════════════════════╝")

				break

			} else if guess < secretNumber {
				fmt.Println(Yellow + "⬆️ Too Low! Try a HIGHER number." + Reset)
			} else {
				fmt.Println(Yellow + "⬇️ Too High! Try a LOWER number." + Reset)
			}

			// Hint
			if i == attempts-2 {
				if secretNumber%2 == 0 {
					fmt.Println(Cyan + "💡 Hint: The number is EVEN." + Reset)
				} else {
					fmt.Println(Cyan + "💡 Hint: The number is ODD." + Reset)
				}
			}
		}

		if !won {
			fmt.Println()
			fmt.Println(Red + "😢 GAME OVER!" + Reset)
			fmt.Printf("The correct number was: %d\n", secretNumber)
			fmt.Println("Better luck next time! 🍀")
		}

		// Play Again
		fmt.Print("\nDo you want to play again? (y/n): ")
		fmt.Scan(&playAgain)

		if playAgain != "y" && playAgain != "Y" {
			fmt.Println()
			fmt.Println(Cyan + "Thanks for playing! 👋" + Reset)
			break
		}

		fmt.Println("\n--------------------------------------\n")
	}
}

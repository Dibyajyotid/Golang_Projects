package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	LoadEnv()
	reader := bufio.NewReader(os.Stdin) // Create one reader to use throughout

	for {
		ShowMenu()
		fmt.Print("Choose an option: ")

		// Read the entire line and convert to number
		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid option - please enter a number")
			continue
		}

		switch choice {
		case 1:
			ShowAllEnv()

		case 2:
			AddEnv(reader) // Pass the reader to AddEnv

		case 3:
			UpdateEnv(reader) // Similarly for other functions that need input

		case 4:
			SaveEnv()

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid Option")
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var customAddedVars = make(map[string]bool)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("❌ No .env file found or failed to load")
	}
}

func ShowMenu() {
	fmt.Println("\n-- ENV MANAGER --")
	fmt.Println("1. View Variables")
	fmt.Println("2. Add Variable")
	fmt.Println("3. Update Variable")
	fmt.Println("4. Save Changes")
	fmt.Println("5. Exit")
}

func ShowAllEnv() {
	envMap, _ := godotenv.Read()
	fmt.Println("\nCurrent Environment Variables:")

	for key, val := range envMap {
		fmt.Printf("%s=%s\n", key, val)
	}
}

func AddEnv(reader *bufio.Reader) {
	// reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter new key: ")
	key, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("❌ Error reading key:", err)
		return
	}
	key = strings.TrimSpace(key)
	if key == "" {
		fmt.Println("Key cannot be empty.")
		return
	}

	if os.Getenv(key) != "" {
		fmt.Println("Variable already exists. Use update option instead.")
		return
	}

	fmt.Print("Enter value: ")
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("❌ Error reading value:", err)
		return
	}
	value = strings.TrimSpace(value)
	if value == "" {
		fmt.Println("Value cannot be empty.")
		return
	}

	err = os.Setenv(key, value)
	if err != nil {
		fmt.Println("❌ Error setting environment variable:", err)
		return
	}

	customAddedVars[key] = true

	fmt.Printf("✅ Added %s=%s\n", key, value)
}

func UpdateEnv(reader *bufio.Reader) {
	fmt.Print("Enter existing key to update: ")
	key, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("❌ Error reading key:", err)
		return
	}
	key = strings.TrimSpace(key)

	// Check if variable exists (either in environment or in customAddedVars)
	if os.Getenv(key) == "" && !customAddedVars[key] {
		fmt.Println("❌ Variable does not exist. Use add option instead.")
		return
	}

	fmt.Print("Enter new value: ")
	value, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("❌ Error reading value:", err)
		return
	}
	value = strings.TrimSpace(value)

	if value == "" {
		fmt.Println("❌ Value cannot be empty")
		return
	}

	// Update the environment variable
	os.Setenv(key, value)
	customAddedVars[key] = true // Mark as managed variable
	fmt.Printf("✅ Updated %s=%s\n", key, value)
}

func SaveEnv() {
	envMap := map[string]string{}

	//Merging already existing .env if it exists
	if oldEnv, err := godotenv.Read(); err == nil {
		maps.Copy(envMap, oldEnv)
	}

	// Add new/updated variables from current session
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			key := parts[0]
			// Only overwrite if it was in original .env or we added it
			if _, existsInEnv := envMap[key]; existsInEnv || customAddedVars[key] {
				envMap[key] = parts[1]
			}
		}
	}

	// Save to file
	err := godotenv.Write(envMap, ".env")
	if err != nil {
		fmt.Println("❌ Failed to save .env:", err)
	} else {
		fmt.Println("✅ .env saved successfully.")
	}
}

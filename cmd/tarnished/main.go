package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var rune_level int
	var weapon_of_choice string

	fmt.Print("Enter current rune level: ")
	fmt.Scan(&rune_level)

	fmt.Print("Enter weapon of choice: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	weapon_of_choice = scanner.Text()

	fmt.Println("You desire to use the", weapon_of_choice, "with a rune level of", rune_level, "...")
}

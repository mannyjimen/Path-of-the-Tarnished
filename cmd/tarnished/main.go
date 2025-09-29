package main

import (
	"Path-of-the-Tarnished/internal/coregame"
	"Path-of-the-Tarnished/internal/optimizer"
	"bufio"
	"fmt"
	"os"
)

func PrintStats(attrs coregame.Attributes) {
	fmt.Println("Vgr -", attrs.Vgr)
	fmt.Println("Mnd -", attrs.Mnd)
	fmt.Println("End -", attrs.End)
	fmt.Println("Str -", attrs.Str)
	fmt.Println("Dex -", attrs.Dex)
	fmt.Println("Int -", attrs.Int)
	fmt.Println("Fth -", attrs.Fth)
	fmt.Println("Arc -", attrs.Arc)
}

func main() {
	var runeLevel int
	var weaponStr string

	fmt.Print("Enter current rune level: ")
	fmt.Scan(&runeLevel)

	fmt.Print("Enter weapon of choice: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	weaponStr = scanner.Text()

	fmt.Println("You desire to use the", weaponStr, "with a rune level of", runeLevel, "...")

	//conv str to weapon
	var weapon coregame.Weapon
	var class coregame.Class

	optimalStats := optimizer.GetOptimizedStats(weapon, class, runeLevel)

	//printing optimal stats
	PrintStats(optimalStats)
}

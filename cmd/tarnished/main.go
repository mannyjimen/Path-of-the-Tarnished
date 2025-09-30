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
	var runeLevel uint16

	fmt.Print("Enter current rune level: ")
	fmt.Scan(&runeLevel)

	fmt.Print("Enter weapon of choice: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	weaponName := scanner.Text()

	fmt.Print("Enter base class: ")
	scanner.Scan()
	className := scanner.Text()

	fmt.Println("You desire to use the", weaponName, "with a rune level of", runeLevel, "and the class", className, "...")

	//conv str to weapon

	optimalStats, err := optimizer.GetOptimizedStats(weaponName, className, runeLevel)

	if err != nil {
		fmt.Println("could not get optimize stats")
		return
	}

	//printing optimal stats
	PrintStats(optimalStats)
}

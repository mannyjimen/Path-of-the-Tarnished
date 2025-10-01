package main

import (
	"Path-of-the-Tarnished/internal/api"
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

var AllWeapons = api.RetrieveWeapons()

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

	//checking for existence of weapon
	_, ok := AllWeapons[weaponName]
	if !ok {
		fmt.Println("Weapon does not exist")
		return
	}

	fmt.Println("Weapon exists!")

	optimalStats, err := optimizer.GetOptimizedStats(weaponName, className, runeLevel)

	if err != nil {
		fmt.Println("Could not optimize stats")
		return
	}

	//printing optimal stats
	PrintStats(optimalStats)
}

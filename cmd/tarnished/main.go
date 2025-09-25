package main

import (
	"Path-of-the-Tarnished/internal/api"
)

func main() {
	// Weapons := api.RetrieveWeapons()
	Talismans := api.RetrieveTalismans()

	// for _, currentWeapon := range Weapons {
	// 	api.PrettyPrintWeapon(currentWeapon)
	// }

	for _, currentTalisman := range Talismans {
		api.PrettyPrintTalisman(currentTalisman)
	}
}

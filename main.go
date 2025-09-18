package main

import "Path-of-the-Tarnished/api"

func main() {
	Weapons := api.RetrieveWeapons()

	for _, currentWeapon := range Weapons {
		api.PrettyPrintWeapon(currentWeapon)
	}
}

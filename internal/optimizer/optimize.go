package optimizer

import (
	"Path-of-the-Tarnished/internal/api"
	"Path-of-the-Tarnished/internal/coregame"
	"fmt"
	"log"
)

func GetOptimizedStats(weaponName string, className string, runeLvl uint16) (coregame.Attributes, error) {
	weapon := api.AllWeapons[weaponName]

	character, err := coregame.NewCharacter(className)

	newDamage := coregame.GetBaseDamage(weapon.Attack[0].Amount, 25)

	fmt.Println(newDamage)

	if err != nil {
		log.Fatalf("couldn't create character: %v", err)
	}
	// runesToUse := runeLvl - character.RuneLvl
	// will do optimized brute force loop here

	return character.Attrs, nil
}

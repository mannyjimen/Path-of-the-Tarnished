package optimizer

import (
	"Path-of-the-Tarnished/internal/coregame"
	"fmt"
	"log"
)

func GetMaxDamageThreeFactors(character *coregame.Character, weapon *coregame.Weapon, runesToUse uint16) float32 {
	return 0
}

// I DONT CARE ABOUT DAMAGE!
func GetMaxDamageTwoFactors(character *coregame.Character, weapon *coregame.Weapon, runesToUse uint16) float32 {
	var bothFactors []string

	for factor := range weapon.ScalingAttrs {
		bothFactors = append(bothFactors, factor)
	}

	//setting one factor to max
	firstFactor := bothFactors[0]
	secondFactor := bothFactors[1]
	character.AddToAttr(firstFactor, runesToUse)

	maxDamage := coregame.CalculateDamage(character, weapon)

	for range runesToUse {
		character.AddToAttr(secondFactor, 1)
		character.SubFromAttr(firstFactor, 1)

		currentDamage := coregame.CalculateDamage(character, weapon)

		maxDamage = max(maxDamage, currentDamage)
	}

	return maxDamage
}

func GetOptimizedStats(weaponName string, className string, runeLvl uint16) (coregame.Attributes, error) {
	weapon := coregame.GetFinalWeapon(weaponName)

	fmt.Println("weapon name:", weapon.Name)
	fmt.Println("scaling attrs:", weapon.ScalingAttrs)
	fmt.Println("base damage:", weapon.BaseDamage)

	character, err := coregame.NewCharacter(className)

	if err != nil {
		log.Fatalf("couldn't create character: %v", err)
	}

	// runesToUse := runeLvl - character.RuneLvl
	// will do optimized brute force loop here

	return character.Attrs, nil
}

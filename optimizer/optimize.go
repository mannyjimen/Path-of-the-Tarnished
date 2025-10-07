package optimizer

import (
	"log"

	"github.com/mannyjimen/Path-of-the-Tarnished/coregame"
)

/*

for every range of the first factor (runesToUse -> 0) i will delete this factor from character, and send to GetMaxTwoFactors
where my runes to use is runesToUse - currentRangeValue for first factor.

*/

func GetMaxDamageThreeFactors(character *coregame.Character, weapon *coregame.Weapon, runesToUse uint16) coregame.Attributes {
	var firstFactor string
	for factor := range weapon.ScalingAttrs {
		firstFactor = factor
		break
	}
	delete(weapon.ScalingAttrs, firstFactor)

	character.AddToAttr(firstFactor, runesToUse)
	initialFirstAttr := character.GetAttr(firstFactor)

	var maxDamage float32
	var tempDamage float32
	var currentOptimalAttrs coregame.Attributes
	var tempAttrs coregame.Attributes

	for i := 0; i <= int(runesToUse); i++ {
		character.SetAttr(firstFactor, initialFirstAttr-uint16(i))

		copyCharacter := *character
		tempAttrs = GetMaxDamageTwoFactors(&copyCharacter, weapon, uint16(i))

		if tempDamage = coregame.CalculateDamage(&copyCharacter.Attrs, weapon); tempDamage > maxDamage {
			maxDamage = tempDamage
			currentOptimalAttrs = tempAttrs
		}
	}

	return currentOptimalAttrs
}

func GetMaxDamageTwoFactors(character *coregame.Character, weapon *coregame.Weapon, runesToUse uint16) coregame.Attributes {
	var bothFactors []string

	for factor := range weapon.ScalingAttrs {
		bothFactors = append(bothFactors, factor)
	}

	//setting one factor to max
	firstFactor := bothFactors[0]
	secondFactor := bothFactors[1]

	character.AddToAttr(firstFactor, runesToUse)
	initialFirstAttr := character.GetAttr(firstFactor)
	initialSecondFactor := character.GetAttr(secondFactor)

	var maxDamage float32
	var currentDamage float32
	var currentOptimalAttrs coregame.Attributes

	for i := 0; i <= int(runesToUse); i++ {
		character.SetAttr(firstFactor, initialFirstAttr-uint16(i))
		character.SetAttr(secondFactor, initialSecondFactor+uint16(i))

		if currentDamage = coregame.CalculateDamage(&character.Attrs, weapon); currentDamage > maxDamage {
			maxDamage = currentDamage
			currentOptimalAttrs = character.Attrs
		}
	}

	return currentOptimalAttrs
}

func GetOptimizedStats(weaponName string, className string, runeLvl uint16) (coregame.Attributes, error) {
	weapon := coregame.GetFinalWeapon(weaponName)

	character, err := coregame.NewCharacter(className)

	if err != nil {
		log.Fatalf("couldn't create character: %v", err)
	}

	runesToUse := runeLvl - character.RuneLvl

	var optimalAttrs coregame.Attributes
	numScalingAttrs := len(weapon.ScalingAttrs)

	switch numScalingAttrs {
	case 2:
		optimalAttrs = GetMaxDamageTwoFactors(character, weapon, runesToUse)
	case 3:
		optimalAttrs = GetMaxDamageThreeFactors(character, weapon, runesToUse)
	}

	return optimalAttrs, nil
}

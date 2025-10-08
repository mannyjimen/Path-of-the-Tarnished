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
	allFactors := weapon.ScalingAttrs

	firstFactor := allFactors[0].ScaleAttr

	character.AddToAttr(firstFactor, runesToUse)
	initialFirstAttr := character.GetAttr(firstFactor)

	var maxDamage float32
	var tempDamage float32
	var currentOptimalAttrs coregame.Attributes
	var tempAttrs coregame.Attributes

	for i := 0; i <= int(runesToUse); i++ {
		character.SetAttr(firstFactor, initialFirstAttr-uint16(i))

		copyCharacter := *character
		tempAttrs = GetMaxDamageTwoFactors(&copyCharacter, weapon, uint16(i), true)

		copyCharacter.Attrs = tempAttrs
		tempDamage = coregame.CalculateDamage(&copyCharacter, weapon)

		if tempDamage > maxDamage {
			maxDamage = tempDamage
			currentOptimalAttrs = tempAttrs
		}
	}

	return currentOptimalAttrs
}

func GetMaxDamageTwoFactors(character *coregame.Character, weapon *coregame.Weapon, runesToUse uint16, fromThree bool) coregame.Attributes {

	allFactors := weapon.ScalingAttrs

	var firstFactor, secondFactor string
	if fromThree {
		firstFactor = allFactors[1].ScaleAttr
		secondFactor = allFactors[2].ScaleAttr
	} else {
		firstFactor = allFactors[0].ScaleAttr
		secondFactor = allFactors[1].ScaleAttr
	}

	character.AddToAttr(firstFactor, runesToUse)
	initialFirstAttr := character.GetAttr(firstFactor)
	initialSecondFactor := character.GetAttr(secondFactor)

	var maxDamage float32
	var currentDamage float32
	var currentOptimalAttrs coregame.Attributes

	for i := 0; i <= int(runesToUse); i++ {
		character.SetAttr(firstFactor, initialFirstAttr-uint16(i))
		character.SetAttr(secondFactor, initialSecondFactor+uint16(i))

		currentDamage = coregame.CalculateDamage(character, weapon)

		if currentDamage > maxDamage {
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
		optimalAttrs = GetMaxDamageTwoFactors(character, weapon, runesToUse, false)
	case 3:
		optimalAttrs = GetMaxDamageThreeFactors(character, weapon, runesToUse)
	}

	return optimalAttrs, nil
}

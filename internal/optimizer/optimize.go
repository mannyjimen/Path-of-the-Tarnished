package optimizer

import (
	"Path-of-the-Tarnished/internal/api"
	"Path-of-the-Tarnished/internal/coregame"
	"fmt"
	"log"
)

func GetScaling(scalingDetails []api.ScaleDetail) map[string]rune {
	var FormattedScalingDetails = make(map[string]rune)
	var attr string
	var scalefactor rune

	for _, currentDetail := range scalingDetails {
		attr = currentDetail.Name
		scalefactor = rune(currentDetail.Scaling[0])

		FormattedScalingDetails[attr] = scalefactor
	}
	return FormattedScalingDetails
}

func GetFinalWeapon(weaponName string) *coregame.Weapon {
	schemaWeapon := api.AllWeapons[weaponName]
	schemaScaling := GetScaling(schemaWeapon.ScalesWith)
	basePhyDamage := schemaWeapon.Attack[0].Amount

	realBaseDamage := coregame.GetBaseDamage(basePhyDamage, 25)

	weapon := coregame.Weapon{
		Name:         weaponName,
		ScalingAttrs: schemaScaling,
		BaseDamage:   realBaseDamage,
	}

	return &weapon
}

func GetOptimizedStats(weaponName string, className string, runeLvl uint16) (coregame.Attributes, error) {
	weapon := GetFinalWeapon(weaponName)

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

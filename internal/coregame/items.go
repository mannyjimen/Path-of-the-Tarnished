package coregame

import "Path-of-the-Tarnished/internal/api"

func getScaling(scalingDetails []api.ScaleDetail) map[string]float32 {
	var FormattedScalingDetails = make(map[string]float32)
	var attr string
	var scaleFactorLetter byte
	var scaleFactorRatio float32

	for _, currentDetail := range scalingDetails {
		attr = currentDetail.Name
		scaleFactorLetter = currentDetail.Scaling[0]
		scaleFactorRatio = scalingRatio[scaleFactorLetter]

		FormattedScalingDetails[attr] = scaleFactorRatio
	}
	return FormattedScalingDetails
}

func GetFinalWeapon(weaponName string) *Weapon {
	schemaWeapon := api.AllWeapons[weaponName]
	schemaScaling := getScaling(schemaWeapon.ScalesWith)
	basePhyDamage := schemaWeapon.Attack[0].Amount

	realBaseDamage := GetBaseDamage(basePhyDamage, 25)

	weapon := Weapon{
		Name:         weaponName,
		ScalingAttrs: schemaScaling,
		BaseDamage:   realBaseDamage,
	}

	return &weapon
}

type Weapon struct {
	Name         string
	ScalingAttrs map[string]float32
	BaseDamage   float32
}

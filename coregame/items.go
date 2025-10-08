package coregame

import "github.com/mannyjimen/Path-of-the-Tarnished/api"

func getScaling(scalingDetails []api.ScaleDetail) []scalePair {
	var FormattedScalingDetails []scalePair
	var attr string
	var scaleFactorLetter byte
	var scaleFactorRatio float32

	for _, currentDetail := range scalingDetails {
		attr = currentDetail.Name
		scaleFactorLetter = currentDetail.Scaling[0]
		scaleFactorRatio = scalingRatio[scaleFactorLetter]
		scalePairEntry := scalePair{
			ScaleAttr:  attr,
			ScaleRatio: scaleFactorRatio,
		}

		FormattedScalingDetails = append(FormattedScalingDetails, scalePairEntry)
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

type scalePair struct {
	ScaleAttr  string
	ScaleRatio float32
}

type Weapon struct {
	Name         string
	ScalingAttrs []scalePair
	BaseDamage   float32
}

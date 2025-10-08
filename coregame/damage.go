package coregame

//total attack rating = base dmg + bonus dmg
//attr spec bonus dmg = base dmg * attr scaling * attr bonus ratio

func CalculateDamage(character *Character, weapon *Weapon) float32 {
	//using current character attributes and current weapon, calc damage output
	baseDamage := weapon.BaseDamage
	var totalBonusDamage float32
	for _, currScalePair := range weapon.ScalingAttrs {
		attrStr := currScalePair.ScaleAttr
		letterRatio := currScalePair.ScaleRatio
		attrLevel := character.GetAttr(attrStr)
		attrBonus := GetAttrBonusRatio(attrLevel)
		attrSpecBonusDamage := baseDamage * letterRatio * attrBonus

		totalBonusDamage += attrSpecBonusDamage
	}

	totalAttackRating := weapon.BaseDamage + totalBonusDamage
	return totalAttackRating
}

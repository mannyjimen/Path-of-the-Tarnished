package coregame

var normalWeaponRatio = []float32{
	1, 1.058, 1.117, 1.175, 1.233, 1.292, 1.350, 1.408, 1.467,
	1.525, 1.583, 1.642, 1.700, 1.715, 1.773, 1.832, 1.890, 1.948,
	2.007, 2.065, 2.123, 2.182, 2.249, 2.298, 2.357, 2.447,
}

var scalingRatio = map[byte]float32{
	'E': 0.225,
	'D': 0.425,
	'C': 0.75,
	'B': 1.15,
	'A': 1.575,
	'S': 1.75,
}

func GetBaseDamage(zeroDamage float32, reinforcedLevel int) float32 {
	return float32(zeroDamage) * normalWeaponRatio[reinforcedLevel]
}

//99 is max level for attr
//assuming soft caps are 20, 55, 80
// level > 80 = 0.10 effective per level addition
// 80 >= level > 55 = 0.40 effective per level addition
// 55 >= level > 20 = 0.75 effective per level addition
// level <= 20 = 1.0 effective per level addition

func GetAttrBonusRatio(level uint16) float32 {
	var currentBonus float32 = 0

	if level > 80 {
		currentBonus += 0.1 * float32(level-80)
		level = 80
	}

	if level > 55 {
		currentBonus += 0.4 * float32(level-55)
		level = 55
	}

	if level > 20 {
		currentBonus += 0.75 * float32(level-20)
		level = 20
	}

	currentBonus += 1 * float32(level)

	//currentBonus needs to be divided my total possible Bonus (level 99)
	//total possible bonus is 58.15

	var totalPossibleBonus float32 = 58.15
	statBonusRatio := currentBonus / totalPossibleBonus

	return statBonusRatio
}

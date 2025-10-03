package coregame

var normalWeaponRatio = []float32{
	1, 1.058, 1.117, 1.175, 1.233, 1.292, 1.350, 1.408, 1.467,
	1.525, 1.583, 1.642, 1.700, 1.715, 1.773, 1.832, 1.890, 1.948,
	2.007, 2.065, 2.123, 2.182, 2.249, 2.298, 2.357, 2.447,
}

func GetBaseDamage(zeroDamage float32, reinforcedLevel int) float32 {

	return float32(zeroDamage) * normalWeaponRatio[reinforcedLevel]
}

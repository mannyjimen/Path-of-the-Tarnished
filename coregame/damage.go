package coregame

import "math/rand/v2"

func CalculateDamage(attrs *Attributes, weapon *Weapon) float32 {
	//using current character attributes and current weapon, calc damage output
	test_num := rand.IntN(1000)
	return float32(test_num)
}

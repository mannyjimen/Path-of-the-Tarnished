package optimizer

import (
	"Path-of-the-Tarnished/internal/coregame"
	"log"
)

func GetOptimizedStats(weaponName string, className string, runeLvl uint16) (coregame.Attributes, error) {

	character, err := coregame.NewCharacter(className)

	if err != nil {
		log.Fatalf("couldn't create character: %v", err)
	}
	// runesToUse := runeLvl - character.RuneLvl
	// will do optimized brute force loop here

	return character.Attrs, nil
}

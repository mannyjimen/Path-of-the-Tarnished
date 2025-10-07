package coregame

import (
	"fmt"
)

type Class int

const (
	Hero Class = iota
	Bandit
	Astrologer
	Warrior
	Prisoner
	Confessor
	Wretch
	Vagabond
	Prophet
	Samurai
)

type Character struct {
	RuneLvl   uint16
	Attrs     Attributes
	BaseClass Class
}

type Attributes struct {
	Vgr uint16
	Mnd uint16
	End uint16
	Str uint16
	Dex uint16
	Int uint16
	Fth uint16
	Arc uint16
}

// short var `:=` can only be used inside functions
var baseClassAttrs = map[string]Attributes{
	"Wretch":     {Vgr: 10, Mnd: 10, End: 10, Str: 10, Dex: 10, Int: 10, Fth: 10, Arc: 10},
	"Hero":       {Vgr: 14, Mnd: 9, End: 12, Str: 16, Dex: 9, Int: 7, Fth: 8, Arc: 11},
	"Bandit":     {Vgr: 10, Mnd: 11, End: 10, Str: 9, Dex: 13, Int: 9, Fth: 8, Arc: 14},
	"Astrologer": {Vgr: 9, Mnd: 15, End: 9, Str: 8, Dex: 12, Int: 16, Fth: 7, Arc: 9},
	"Warrior":    {Vgr: 11, Mnd: 12, End: 11, Str: 10, Dex: 16, Int: 10, Fth: 8, Arc: 9},
	"Prisoner":   {Vgr: 11, Mnd: 12, End: 11, Str: 11, Dex: 14, Int: 14, Fth: 6, Arc: 9},
	"Confessor":  {Vgr: 10, Mnd: 13, End: 10, Str: 12, Dex: 12, Int: 9, Fth: 14, Arc: 9},
	"Vagabond":   {Vgr: 15, Mnd: 10, End: 11, Str: 14, Dex: 13, Int: 9, Fth: 9, Arc: 7},
	"Prophet":    {Vgr: 10, Mnd: 14, End: 8, Str: 11, Dex: 10, Int: 7, Fth: 16, Arc: 10},
	"Samurai":    {Vgr: 12, Mnd: 11, End: 13, Str: 12, Dex: 15, Int: 9, Fth: 8, Arc: 8},
}

var baseClassRuneLevel = map[string]uint16{
	"Wretch":     1,
	"Hero":       7,
	"Bandit":     5,
	"Astrologer": 6,
	"Warrior":    8,
	"Prisoner":   9,
	"Confessor":  10,
	"Vagabond":   9,
	"Prophet":    7,
	"Samurai":    9,
}

var baseClassType = map[string]Class{
	"Wretch":     Wretch,
	"Hero":       Hero,
	"Bandit":     Bandit,
	"Astrologer": Astrologer,
	"Warrior":    Warrior,
	"Prisoner":   Prisoner,
	"Confessor":  Confessor,
	"Vagabond":   Vagabond,
	"Prophet":    Prophet,
	"Samurai":    Samurai,
}

// first method (pointer receiver)
func (character *Character) AddToAttr(attr string, levelsToAdd uint16) {
	switch attr {
	case "Vgr":
		character.Attrs.Vgr += levelsToAdd
	case "Mnd":
		character.Attrs.Mnd += levelsToAdd
	case "End":
		character.Attrs.End += levelsToAdd
	case "Str":
		character.Attrs.Str += levelsToAdd
	case "Dex":
		character.Attrs.Dex += levelsToAdd
	case "Int":
		character.Attrs.Int += levelsToAdd
	case "Fai":
		character.Attrs.Fth += levelsToAdd
	case "Arc":
		character.Attrs.Arc += levelsToAdd
	}
}

func (character *Character) SetAttr(attr string, runeLevel uint16) {
	switch attr {
	case "Vgr":
		character.Attrs.Vgr = runeLevel
	case "Mnd":
		character.Attrs.Mnd = runeLevel
	case "End":
		character.Attrs.End = runeLevel
	case "Str":
		character.Attrs.Str = runeLevel
	case "Dex":
		character.Attrs.Dex = runeLevel
	case "Int":
		character.Attrs.Int = runeLevel
	case "Fai":
		character.Attrs.Fth = runeLevel
	case "Arc":
		character.Attrs.Arc = runeLevel
	}
}

func (character Character) GetAttr(attr string) uint16 {
	switch attr {
	case "Vgr":
		return character.Attrs.Vgr
	case "Mnd":
		return character.Attrs.Mnd
	case "End":
		return character.Attrs.End
	case "Str":
		return character.Attrs.Str
	case "Dex":
		return character.Attrs.Dex
	case "Int":
		return character.Attrs.Int
	case "Fai":
		return character.Attrs.Fth
	case "Arc":
		return character.Attrs.Arc
	}
	return 0
}

// constructor function (using error!)
func NewCharacter(className string) (*Character, error) {
	attrs, ok := baseClassAttrs[className]

	if !ok {
		return nil, fmt.Errorf("base class %q does not exist", className)
	}

	runeLvl := baseClassRuneLevel[className]
	baseClass := baseClassType[className]

	//creating new Character instance
	character := &Character{RuneLvl: runeLvl, Attrs: attrs, BaseClass: baseClass}

	return character, nil
}

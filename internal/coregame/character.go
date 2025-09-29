package coregame

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
	Attrs     Attributes
	BaseClass Class
}

type Attributes struct {
	Vgr uint8
	Mnd uint8
	End uint8
	Str uint8
	Dex uint8
	Int uint8
	Fth uint8
	Arc uint8
}

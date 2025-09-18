package api

type WeaponJSON struct {
	Success bool     `json:"success"`
	Data    []Weapon `json:"data"`
}

type AttackDetail struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
}

type DefenceDetail struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
}

type RequiredAttribute struct {
	Name   string  `json:"name"`
	Amount float32 `json:"amount"`
}

type ScaleDetail struct {
	Name    string `json:"name"`
	Scaling string `json:"scaling"`
}

type Weapon struct {
	Id                 string              `json:"id"`
	Name               string              `json:"name"`
	Image              string              `json:"image"`
	Description        string              `json:"description"`
	Category           string              `json:"category"`
	Weight             float32             `json:"weight"`
	Attack             []AttackDetail      `json:"attack"`
	Defence            []DefenceDetail     `json:"defence"`
	RequiredAttributes []RequiredAttribute `json:"requiredAttributes"`
	ScalesWith         []ScaleDetail       `json:"scalesWith"`
}

package api

type WeaponsJSON struct {
	Success bool     `json:"success"`
	Count   float64  `json:"count"`
	Data    []Weapon `json:"data"`
}

type TalismansJSON struct {
	Success bool       `json:"success"`
	Data    []Talisman `json:"data"`
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

type Talisman struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Effects     string `json:"effect"`
}

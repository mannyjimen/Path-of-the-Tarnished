package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type generalDetails struct {
	Name   string
	Amount float32
}

type scaleDetails struct {
	Name    string
	Scaling string
}

type Weapon struct {
	Id                 string         `json:"id"`
	Name               string         `json:"name"`
	Image              string         `json:"image"`
	Description        string         `json:"description"`
	Category           string         `json:"category"`
	Weight             float32        `json:"weight"`
	Attack             generalDetails `json:"attack"`
	Defence            generalDetails `json:"defence"`
	RequiredAttributes generalDetails `json:"requiredAttributes"`
	ScalesWith         scaleDetails   `json:"scalesWith"`
}

func main() {

	resp, err := http.Get("https://eldenring.fanapis.com/api/weapons/17f69aeb629l0i1oifaza9fgo5b598")

	if err != nil {
		fmt.Println("failed to http GET")
		return
	}

	defer resp.Body.Close()

	byteBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("failed to read respone")
		return
	}

	var highlandAxe Weapon

	json.Unmarshal(byteBody, &highlandAxe)

	// fmt.Println(string(byteBody))

}

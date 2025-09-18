package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func PrettyPrintWeapon(currWeapon Weapon) {
	fmt.Println("Name:", currWeapon.Name)
	fmt.Println("Description:", currWeapon.Description)
}

func retrieveWeaponPage(url string) []Weapon {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to http GET")
		return nil
	}

	defer resp.Body.Close()

	byteBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("failed to read respone")
		return nil
	}

	var AllWeaponsJSON WeaponJSON
	json.Unmarshal(byteBody, &AllWeaponsJSON)

	return AllWeaponsJSON.Data
}

func RetrieveWeapons() []Weapon {

	var AllWeapons []Weapon
	current_url := "https://eldenring.fanapis.com/api/weapons?limit=30&page="

	for currentPage := range 11 {
		currentRangeOfWeapons := retrieveWeaponPage(current_url + strconv.Itoa(currentPage))
		AllWeapons = append(AllWeapons, currentRangeOfWeapons...)
	}

	// fmt.Println(AllWeapons)
	// for _, currWeapon := range AllWeapons {
	// 	prettyPrintWeapon(currWeapon)
	// }

	return AllWeapons
}

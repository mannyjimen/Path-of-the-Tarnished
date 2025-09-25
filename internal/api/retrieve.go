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

func PrettyPrintTalisman(currTalisman Talisman) {
	fmt.Println("Name:", currTalisman.Name)
	fmt.Println("Description:", currTalisman.Description)
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
		fmt.Println("failed to read weapon page")
		return nil
	}

	var AllWeaponsJSON WeaponsJSON
	json.Unmarshal(byteBody, &AllWeaponsJSON)

	return AllWeaponsJSON.Data
}

func retrieveTalismanPage(url string) []Talisman {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("failed to GET")
		return nil
	}

	defer resp.Body.Close()

	byteBody, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("failed to read talisman page")
		return nil
	}

	var AllTalismansJSON TalismansJSON
	json.Unmarshal(byteBody, &AllTalismansJSON)

	return AllTalismansJSON.Data
}

func RetrieveWeapons() []Weapon {

	var AllWeapons []Weapon
	current_url := "https://eldenring.fanapis.com/api/weapons?limit=30&page="

	//fix - what if more weapons are added to API (im hard coding num pages)
	for currentPage := range 11 {
		currentRangeOfWeapons := retrieveWeaponPage(current_url + strconv.Itoa(currentPage))
		AllWeapons = append(AllWeapons, currentRangeOfWeapons...)
	}

	return AllWeapons
}

func RetrieveTalismans() []Talisman {
	var AllTalismans []Talisman
	current_url := "https://eldenring.fanapis.com/api/talismans?limit=88"

	AllTalismans = retrieveTalismanPage(current_url)

	return AllTalismans
}

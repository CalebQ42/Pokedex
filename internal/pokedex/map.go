package pokedex

import (
	"encoding/json"
	"fmt"
)

type mapListRes struct {
	Next     string
	Previous string
	Results  []struct {
		Name string
		Url  string
	}
}

func (p *Pokedex) mapListNext(...string) {
	if p.nextMapURL == "" {
		fmt.Println("No more locations to get")
		return
	}
	dat, err := p.getURLData(p.nextMapURL)
	if err != nil {
		fmt.Println("Error while getting the next map list:", err)
		return
	}
	var res mapListRes
	err = json.Unmarshal(dat, &res)
	if err != nil {
		fmt.Println("Error while marshaling response:", err)
		return
	}
	p.nextMapURL = res.Next
	p.prevMapURL = res.Previous
	for _, l := range res.Results {
		fmt.Println(l.Name)
	}
}

func (p *Pokedex) mapListPrev(...string) {
	if p.prevMapURL == "" {
		fmt.Println("Please advance map further before going back")
		return
	}
	dat, err := p.getURLData(p.prevMapURL)
	if err != nil {
		fmt.Println("Error while reading the previous map list:", err)
		return
	}
	var res mapListRes
	err = json.Unmarshal(dat, &res)
	if err != nil {
		fmt.Println("Error while marshaling response:", err)
		return
	}
	p.nextMapURL = res.Next
	p.prevMapURL = res.Previous
	for _, l := range res.Results {
		fmt.Println(l.Name)
	}
}

type locationRes struct {
	Encounters []struct {
		Pokemon struct {
			Name string
		}
	} `json:"pokemon_encounters"`
}

func (p *Pokedex) explore(args ...string) {
	if len(args) == 0 {
		fmt.Println("Please provide a location name")
		return
	}
	url := "https://pokeapi.co/api/v2/location-area/" + args[0]
	dat, err := p.getURLData(url)
	if err != nil {
		fmt.Println("Error while getting location data:", err)
		return
	}
	if string(dat) == "Not Found" {
		fmt.Println(args[0], "not a valid location")
		return
	}
	var res locationRes
	err = json.Unmarshal(dat, &res)
	if err != nil {
		fmt.Println("Error while marshalling json:", err)
		return
	}
	fmt.Printf("Exploring %v...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, p := range res.Encounters {
		fmt.Printf(" - %v\n", p.Pokemon.Name)
	}
}

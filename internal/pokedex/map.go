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

func (p *Pokedex) mapListNext() {
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

func (p *Pokedex) mapListPrev() {
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

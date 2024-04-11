package pokedex

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"strings"
)

type pokemon struct {
	Name  string
	Stats []pokeStats
	Types []struct {
		Typ struct{ Name string } `json:"type"`
	}
	Height  int
	Weight  int
	BaseExp int `json:"base_experience"`
}

type pokeStats struct {
	Stat struct {
		Name string
	}
	BaseStat int `json:"base_stat"`
}

func (p *Pokedex) catch(args ...string) {
	if len(args) == 0 {
		fmt.Println("Please specify a pokemon you want to catch")
		return
	}
	pmn := strings.ToLower(args[0])
	if _, ok := p.curPokemon[pmn]; ok {
		fmt.Println("A", pmn, "has already been caught")
		return
	}
	dat, err := p.getURLData("https://pokeapi.co/api/v2/pokemon/" + pmn)
	if err != nil {
		fmt.Printf("Error while getting data for %v: %v\n", pmn, err)
		return
	}
	if string(dat) == "Not Found" {
		fmt.Println(pmn, "not a valid pokemon")
		return
	}
	var poke pokemon
	err = json.Unmarshal(dat, &poke)
	if err != nil {
		fmt.Println("Error while marshaling result:", err)
		return
	}
	fmt.Printf("Throwing a Pokeball at %v...\n", pmn)
	var r int
	// Blissey has such a high base exp, we handle it specially so that it doesn't make everything else trivial to catch
	if poke.BaseExp < 400 {
		r = rand.IntN(400)
	} else {
		// Blissey clause
		r = rand.IntN(609)
	}
	if r <= poke.BaseExp {
		fmt.Println(pmn, "escaped!")
		return
	}
	fmt.Println(pmn, "was caught!")
	p.curPokemon[pmn] = poke
}

func (p *Pokedex) inspect(args ...string) {
	if len(args) == 0 {
		fmt.Println("Please specify a caught pokemon to inspect")
		return
	}
	pmn := strings.ToLower(args[0])
	poke, ok := p.curPokemon[pmn]
	if !ok {
		fmt.Println("You must catch the pokemon before inspecting it")
		return
	}
	fmt.Println("Name:", poke.Name)
	fmt.Println("Height:", poke.Height)
	fmt.Println("Weight:", poke.Weight)
	fmt.Println("Stats:")
	for _, s := range poke.Stats {
		fmt.Printf(" -%v: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf(" - %v\n", t.Typ.Name)
	}
}

func (p *Pokedex) list(...string) {
	if len(p.curPokemon) == 0 {
		fmt.Println("You GOTTA CATCH THEM ALL (or at least one)")
		return
	}
	fmt.Println("Your Pokedex:")
	for i := range p.curPokemon {
		fmt.Println(" -", i)
	}
}

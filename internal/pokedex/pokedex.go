package pokedex

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/CalebQ42/Pokedex/internal/cache"
)

type cmd struct {
	fn   func(args ...string)
	name string
	desc string
}

type Pokedex struct {
	curPokemon map[string]pokemon
	commands   map[string]cmd
	cache      *cache.Cache
	nextMapURL string
	prevMapURL string
}

func NewPokedex(quit *bool) *Pokedex {
	p := new(Pokedex)
	p.curPokemon = make(map[string]pokemon)
	p.cache = cache.NewCache(time.Minute)
	p.nextMapURL = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	p.commands = map[string]cmd{
		"quit": {
			name: "quit",
			desc: "Close the Pokedex",
			fn: func(...string) {
				*quit = true
			},
		},
		"help": {
			name: "help",
			desc: "Print this usage message",
			fn:   p.help,
		},
		"map": {
			name: "map",
			desc: "List the next 20 map locations",
			fn:   p.mapListNext,
		},
		"mapb": {
			name: "mapb",
			desc: "List the previous 20 map locations",
			fn:   p.mapListPrev,
		},
		"explore": {
			name: "explore <map location>",
			desc: "Explore a map area (lists the Pokemon found in the area)",
			fn:   p.explore,
		},
		"catch": {
			name: "catch <pokemon>",
			desc: "Attempt to catch the specified Pokemon",
			fn:   p.catch,
		},
		"inspect": {
			name: "inspect <pokemon>",
			desc: "Inspect a caught Pokemon",
			fn:   p.inspect,
		},
		"pokedex": {
			name: "pokedex",
			desc: "List all your caught Pokemon",
			fn:   p.list,
		},
	}
	return p
}

func (p *Pokedex) Handle(command string) {
	command = strings.Trim(command, " ")
	if command == "" {
		return
	}
	spl := strings.Split(command, " ")
	c, ok := p.commands[spl[0]]
	if !ok {
		fmt.Println("Invalid Command")
		return
	}
	c.fn(spl[1:]...)
}

func (p *Pokedex) help(...string) {
	fmt.Println("\nUsage:")
	fmt.Println()
	for i := range p.commands {
		fmt.Printf("%v: %v\n", i, p.commands[i].desc)
	}
	fmt.Println()
}

func (p *Pokedex) getURLData(url string) ([]byte, error) {
	dat, ok := p.cache.Get(url)
	if ok {
		return dat, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	dat, err = io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	p.cache.Add(url, dat)
	return dat, nil
}

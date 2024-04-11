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
	desc string
}

type Pokedex struct {
	commands   map[string]cmd
	cache      *cache.Cache
	nextMapURL string
	prevMapURL string
}

func NewPokedex(quit *bool) *Pokedex {
	p := new(Pokedex)
	p.cache = cache.NewCache(time.Minute)
	p.nextMapURL = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	p.commands = map[string]cmd{
		"quit": {
			desc: "Close the Pokedex",
			fn: func(...string) {
				*quit = true
			},
		},
		"help": {
			desc: "Print this usage message",
			fn:   p.help,
		},
		"map": {
			desc: "List the next 20 map locations",
			fn:   p.mapListNext,
		},
		"mapb": {
			desc: "List the previous 20 map locations",
			fn:   p.mapListPrev,
		},
		"explore": {
			desc: "Explore a map area (lists the Pokemon found in the area)",
		},
	}
	return p
}

func (p *Pokedex) Handle(command string) {
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

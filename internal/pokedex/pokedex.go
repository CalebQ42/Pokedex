package pokedex

import "fmt"

type cmd struct {
	fn   func()
	name string
	desc string
}

func (c cmd) String() string {
	return fmt.Sprintf("%v: %v", c.name, c.desc)
}

type Pokedex struct {
	commands map[string]cmd
}

func NewPokedex(quit *bool) *Pokedex {
	p := new(Pokedex)
	p.commands = map[string]cmd{
		"quit": {
			name: "quit",
			desc: "Close the Pokedex",
			fn: func() {
				*quit = true
			},
		},
		"help": {
			name: "help",
			desc: "Print this usage message",
			fn:   p.help,
		},
	}
	return p
}

func (p *Pokedex) Handle(command string) {
	c, ok := p.commands[command]
	if !ok {
		fmt.Println("Invalid Command")
		return
	}
	c.fn()
}

func (p *Pokedex) help() {
	fmt.Println("\nUsage:")
	fmt.Println()
	for i := range p.commands {
		fmt.Println(p.commands[i])
	}
	fmt.Println()
}

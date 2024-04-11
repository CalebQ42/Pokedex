package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/CalebQ42/Pokedex/internal/pokedex"
)

var (
	quit = false
)

func main() {
	p := pokedex.NewPokedex(&quit)
	scan := bufio.NewScanner(os.Stdin)
	for !quit {
		fmt.Print("Pokedex > ")
		scan.Scan()
		p.Handle(strings.ToLower(scan.Text()))
	}
}

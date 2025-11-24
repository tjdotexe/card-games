package main

import (
	"fmt"
	"github.com/tjdotexe/card-games/internal/deck"
)

func main() {
	fmt.Println("initializing ninety-nine...")
	deck := deck.Initialize()
	deck.Shuffle()
}

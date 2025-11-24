package deck

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

type Card struct {
	Face string
	Suit string
}

type Deck struct {
	ID    uuid.UUID
	Cards []Card
}

func (d Deck) String() string {
	return fmt.Sprintf("Deck{ID: %s, Cards: %d}", d.ID, len(d.Cards))
}

func (c Card) String() string {
	return fmt.Sprintf("Card{%s, %s}", c.Face, c.Suit)
}

func Initialize() Deck {
	var deck Deck
	deck = Deck{uuid.New(), createDeck()}
	return deck
}

func createDeck() []Card {
	suits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	faces := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	var deck []Card
	for _, suit := range suits {
		for _, face := range faces {
			card := Card{face, suit}
			deck = append(deck, card)
		}
	}
	fmt.Println("deck created")
	return deck
}

func (d *Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
	fmt.Println("cards shuffled")
}

func deal() {

}

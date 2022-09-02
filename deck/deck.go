package deck

import (
	"math/rand"
	"sync/atomic"
	"time"

	"github.com/rishabh053/card-game/card"
	"github.com/rishabh053/card-game/player"
)

type Deck []card.Card

var CurrentCard = atomic.Int32{}

var Suits = []card.Suit{
	card.Suit("Heart(♥)"),
	card.Suit("Diamond(♦)"),
	card.Suit("Club(♣)"),
	card.Suit("Spade(♠)"),
}

var CardFaces = []card.CardFace{
	{Type: "Two", Rank: 2},
	{Type: "Three", Rank: 3},
	{Type: "Four", Rank: 4},
	{Type: "Five", Rank: 5},
	{Type: "Six", Rank: 6},
	{Type: "Seven", Rank: 7},
	{Type: "Eight", Rank: 8},
	{Type: "Nine", Rank: 9},
	{Type: "Ten", Rank: 10},
	{Type: "Jack", Rank: 11},
	{Type: "Queen", Rank: 12},
	{Type: "King", Rank: 13},
	{Type: "Ace", Rank: 14},
}

func NewDeck() Deck {
	var deck Deck
	for _, suit := range Suits {
		for _, cardface := range CardFaces {
			card := card.Card{
				Face: cardface,
				Suit: suit,
			}
			deck = append(deck, card)
		}
	}
	return deck
}

func (d Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range d {
		r := rand.Intn(i + 1)
		if i != r {
			d[r], d[i] = d[i], d[r]
		}
	}
}

func (d Deck) DistributeCard(playerList []*player.Player) {
	numofPlayer := len(playerList)
	for i, player := range playerList {
		player.Cards = append(player.Cards, d[i])
		player.Cards = append(player.Cards, d[i+numofPlayer])
		player.Cards = append(player.Cards, d[i+numofPlayer*2])
		CurrentCard.Add(3)
	}
}

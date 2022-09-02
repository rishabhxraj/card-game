package player

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rishabh053/card-game/card"
	"github.com/rishabh053/card-game/deck"
	"github.com/rishabh053/card-game/player"
)

func TestPlayer_IsCardTrail(t *testing.T) {
	p1 := getPlayer()
	tests := []struct {
		name string
		p    player.Player
		want bool
	}{
		{"Test1", p1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsCardTrail(); got != tt.want {
				t.Errorf("Player.IsCardTrail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_TopCard(t *testing.T) {
	p1 := getPlayer()

	tests := []struct {
		name string
		p    player.Player
		want card.Card
	}{
		{"Test1", p1, p1.Cards[2]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.TopCard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Player.TopCard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_GetPairSum(t *testing.T) {
	p1 := getPairedCardPlayer()

	tests := []struct {
		name string
		p    player.Player
		want int
	}{
		{"Test1", p1, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.p.Cards)
			if got := tt.p.GetPairSum(); got != tt.want {
				t.Errorf("Player.GetPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_GetSum(t *testing.T) {

	p1 := getPlayer()

	tests := []struct {
		name string
		p    player.Player
		want int
	}{
		{"Test1", p1, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.p.Cards)
			if got := tt.p.GetSum(); got != tt.want {
				t.Errorf("Player.GetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_IsCardSequence(t *testing.T) {

	p1 := getPlayer()

	tests := []struct {
		name string
		p    player.Player
		want bool
	}{
		{"Test1", p1, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.p.Cards)
			if got := tt.p.IsCardSequence(); got != tt.want {
				t.Errorf("Player.IsCardSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPlayer_IsCardPair(t *testing.T) {
	p1 := getPlayer()
	p2 := getPairedCardPlayer()

	tests := []struct {
		name string
		p    player.Player
		want bool
	}{
		{"Test1", p1, false},
		{"Test1", p2, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.IsCardPair(); got != tt.want {
				t.Errorf("Player.IsCardPair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getPlayer() player.Player {
	player := player.New("Test")
	player.Cards = getCards()
	return player
}

func getPairedCardPlayer() player.Player {
	player := player.New("TestPair")
	player.Cards = getPairCards()
	return player
}

func getPairCards() player.PlayerCardList {
	var cardlist player.PlayerCardList
	deck := deck.NewDeck()
	a := deck[0]
	var b card.Card
	for _, card := range deck {
		if card.Face.Rank == a.Face.Rank {
			b = card
		}
	}
	c := deck[len(deck)-1]
	cardlist = append(cardlist, a, b, c)
	return cardlist
}

func getCards() player.PlayerCardList {
	var cardlist player.PlayerCardList
	deck := deck.NewDeck()
	a := deck[0]
	b := deck[1]
	c := deck[2]
	cardlist = append(cardlist, a, b, c)
	return player.PlayerCardList(cardlist)
}

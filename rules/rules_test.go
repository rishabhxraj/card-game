package rules

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rishabh053/card-game/card"
	"github.com/rishabh053/card-game/deck"
	"github.com/rishabh053/card-game/player"
)

func TestCheckWinnerForHighestTrail(t *testing.T) {
	type args struct {
		player1 player.Player
		player2 player.Player
	}
	var p args
	p.player1 = getTrailingCardPlayer(0)
	p.player2 = getTrailingCardPlayer(12)

	tests := []struct {
		name string
		args args
		want player.Player
	}{
		{"Test#1", p, p.player2}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckWinnerForHighestTrail(tt.args.player1, tt.args.player2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckWinnerForHighestTrail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getTrailingCardPlayer(index int) player.Player {
	p := player.New(fmt.Sprintf("Test#%d", index))
	card := card.Card{
		Face: deck.CardFaces[index],
		Suit: deck.Suits[0],
	}
	p.Cards = append(p.Cards, card, card, card)
	return p
}

func TestCheckWinnerForHighestSequence(t *testing.T) {
	type args struct {
		player1 player.Player
		player2 player.Player
	}
	var p args
	p.player1 = getSequenceCardPlayer(0)
	p.player2 = getSequenceCardPlayer(3)

	tests := []struct {
		name string
		args args
		want player.Player
	}{
		{"Test#1", p, p.player2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckWinnerForHighestSequence(tt.args.player1, tt.args.player2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckWinnerForHighestSequence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getSequenceCardPlayer(index int) player.Player {
	p := player.New(fmt.Sprintf("Test#%d", index))
	for i := 0; i < 3; i++ {
		card := card.Card{
			Face: deck.CardFaces[index+i],
			Suit: deck.Suits[0],
		}
		p.Cards = append(p.Cards, card)
	}
	return p
}

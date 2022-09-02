package rules

import (
	"fmt"

	"github.com/rishabh053/card-game/card"
	"github.com/rishabh053/card-game/deck"
	"github.com/rishabh053/card-game/player"
)

const NumofPlayer = 4

type Priority int

const (
	HighestCard Priority = iota
	Pair
	Sequence
	Trail
)

func CheckWinnerForHighestTrail(player1, player2 player.Player) player.Player {
	if player1.GetSum() > player2.GetSum() {
		return player1
	}
	return player2
}

func CheckWinnerForHighestSequence(player1, player2 player.Player) player.Player {
	if player1.GetSum() > player2.GetSum() {
		return player1
	}
	return player2
}

func CheckWinnerForHighestPair(player1, player2 player.Player) player.Player {
	if player1.GetPairSum() > player2.GetPairSum() {
		return player1
	}
	return player2
}

func CheckWinnerForHighCard(deck deck.Deck, player1, player2 player.Player) player.Player {
	if player1.TopCard().Face.Rank > player2.TopCard().Face.Rank {
		return player1
	}
	return player2
}

func TopCardOfGame(playerlist []*player.Player) card.Card {
	topCard := playerlist[0].TopCard()
	for _, player := range playerlist {
		if topCard.Face.Rank < player.TopCard().Face.Rank {
			topCard = player.TopCard()
		}
	}
	return topCard
}

func BiggerOfPickedCards(p1, p2 player.Player) player.Player {
	if p1.Cards[len(p1.Cards)-1].Face.Rank > p2.Cards[len(p2.Cards)-1].Face.Rank {
		return p1
	}
	return p2
}

func CheckCardPriority(p player.Player) Priority {
	p.Cards.SortCards()
	if p.IsCardTrail() {
		return Trail
	}
	if p.IsCardSequence() {
		return Sequence
	}
	if p.IsCardPair() {
		return Pair
	}
	return HighestCard
}

func PickOneMore(d deck.Deck, player1, player2 *player.Player) {

	if !player1.Picked {
		fmt.Printf("\n%s picked %s of %s\n", player1.Name, d[deck.CurrentCard.Load()].Face.Type, d[deck.CurrentCard.Load()].Suit)
		player1.Cards = append(player1.Cards, d[deck.CurrentCard.Load()])
		deck.CurrentCard.Add(1)
		player1.Picked = true
	}

	if !player2.Picked {
		fmt.Printf("\n%s picked %s of %s\n", player2.Name, d[deck.CurrentCard.Load()].Face.Type, d[deck.CurrentCard.Load()].Suit)
		player2.Cards = append(player2.Cards, d[deck.CurrentCard.Load()])
		deck.CurrentCard.Add(1)
		player2.Picked = true

	}
}

package manager

import (
	"fmt"
	"reflect"

	"github.com/rishabh053/card-game/card"
	"github.com/rishabh053/card-game/deck"
	"github.com/rishabh053/card-game/player"
	"github.com/rishabh053/card-game/rules"
)

var colorGreen = "\033[32m"
var colorReset = "\033[0m"

func StartGame() {
	deck := deck.NewDeck()
	var playerList []*player.Player
	var playerName string
	for i := 0; i < rules.NumofPlayer; i++ {
		fmt.Printf("Enter Player %d Name: ", i+1)
		fmt.Scanln(&playerName)
		player := player.New(playerName)
		playerList = append(playerList, &player)
	}
	deck.Shuffle()
	deck.DistributeCard(playerList)
	startGamePlay(deck, playerList)
}

func startGamePlay(deck deck.Deck, playerlist []*player.Player) {
	print(playerlist)
	winner := show(deck, playerlist)
	fmt.Println("Winner is", string(colorGreen), winner.Name, string(colorReset))
	fmt.Println()
}

func print(playerlist []*player.Player) {
	for _, p := range playerlist {
		p.Cards.SortCards()
		fmt.Printf("\n%s's Card\n", p.Name)
		for _, card := range p.Cards {
			fmt.Printf(" %s of %s ", card.Face.Type, card.Suit)
		}
		fmt.Println()
	}
	fmt.Println()
}

func show(deck deck.Deck, playerlist []*player.Player) player.Player {
	var winner player.Player
	topCard := rules.TopCardOfGame(playerlist)
	for len(playerlist) != 1 {
		winner = winnerAmongTwo(topCard, deck, *playerlist[0], *playerlist[1])
		if reflect.DeepEqual(winner, *playerlist[0]) {
			playerlist = append(playerlist[:1], playerlist[2:]...)
		} else {
			playerlist = append(playerlist[:0], playerlist[1:]...)
		}
	}
	return winner
}

func winnerAmongTwo(topCard card.Card, deck deck.Deck, player1, player2 player.Player) player.Player {
	player1Priority := rules.CheckCardPriority(player1)
	player2Priority := rules.CheckCardPriority(player2)

	if player1Priority > player2Priority {
		return player1
	}

	if player1Priority == player2Priority {
		switch player1Priority {
		case rules.Trail:
			return rules.CheckWinnerForHighestTrail(player1, player2)
		case rules.Sequence:
			return rules.CheckWinnerForHighestSequence(player1, player2)
		case rules.Pair:
			return rules.CheckWinnerForHighestPair(player1, player2)
		case rules.HighestCard:
			if player1.TopCard().Face.Rank == topCard.Face.Rank && player2.TopCard().Face.Rank == topCard.Face.Rank &&
				!player1.Picked && !player2.Picked {
				rules.PickOneMore(deck, &player1, &player2)
				return rules.BiggerOfPickedCards(player1, player2)
			} else {
				return rules.CheckWinnerForHighCard(deck, player1, player2)
			}
		}
	}

	return player2
}

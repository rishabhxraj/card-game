package player

import (
	"sort"

	"github.com/rishabh053/card-game/card"
)

type Player struct {
	Name   string
	Cards  PlayerCardList
	Picked bool
}

type PlayerCardList []card.Card

func New(name string) Player {
	var p Player
	p.Name = name
	return p
}

func (p Player) GetSum() int {
	sum := 0
	for _, c := range p.Cards {
		sum += c.Face.Rank
	}
	return sum
}

func (p Player) GetPairSum() int {
	sum := 0
	for i := 0; i < len(p.Cards)-1; i++ {
		if p.Cards[i].Face.Rank == p.Cards[i+1].Face.Rank {
			sum = p.Cards[i].Face.Rank + p.Cards[i+1].Face.Rank
		}
	}
	return sum
}

func (p PlayerCardList) SortCards() {
	sort.Slice(p, func(i, j int) bool {
		return p[i].Face.Rank < p[j].Face.Rank
	})
}

func (p Player) IsCardTrail() bool {
	count := 0
	for i := 0; i < len(p.Cards)-1; i++ {
		if p.Cards[i].Face.Rank == p.Cards[i+1].Face.Rank {
			count++
		}
	}
	return count == len(p.Cards)-1
}

func (p Player) IsCardSequence() bool {
	count := 0
	for i := 0; i < len(p.Cards)-1; i++ {
		if p.Cards[i].Face.Rank+1 == p.Cards[i+1].Face.Rank {
			count++
		}
	}
	return count == len(p.Cards)-1
}

func (p Player) IsCardPair() bool {
	count := 0
	for i := 0; i < len(p.Cards)-1; i++ {
		if p.Cards[i].Face.Rank == p.Cards[i+1].Face.Rank {
			count++
		}
	}
	return count >= 1
}

func (p Player) TopCard() card.Card {
	p.Cards.SortCards()
	return p.Cards[len(p.Cards)-1]
}

package card

type Suit string

type CardFace struct {
	Type string
	Rank int
}

type Card struct {
	Face CardFace
	Suit Suit
}

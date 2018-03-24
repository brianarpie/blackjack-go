//go:generate stringer -type=Suit,Rank
package deck

import "fmt"

type Suit uint8

const (
  Spade Suit = iota
  Diamond
  Club
  Heart
)

type Rank uint8

const (
  _ Rank = iota // increments, starting with 0
  Ace
  Two
  Three
  Four
  Five
  Six
  Seven
  Eight
  Nine
  Ten
  Jack
  Queen
  King
)

type Card struct {
  Suit
  Rank
}

func (c Card) String() string {
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}
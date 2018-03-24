package deck

import (
  "fmt"
  "testing"
)

func ExampleCard() {
  fmt.Println(Card{Rank: Ace, Suit: Heart})
  fmt.Println(Card{Rank: Two, Suit: Spade})
  fmt.Println(Card{Rank: Three, Suit: Club})
  fmt.Println(Card{Rank: Four, Suit: Diamond})

  // Output:
  // Ace of Hearts
  // Two of Spades
  // Three of Clubs
  // Four of Diamonds
}

func TestNew(t *testing.T) {
  cards := New()
  // 13 ranks * 4 suits
  if len(cards) != 13*4 {
    t.Error("Wrong number of cards in a new deck.")
  }
}

func TestDefaultSort(t *testing.T) {
  cards := New(DefaultSort)
  exp := Card{Rank: Ace, Suit: Spade}
  if cards[0] != exp {
    t.Error("Expected Ace of Spades as first card. Received:", cards[0])
  }
}

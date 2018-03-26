package deck

import (
  "fmt"
  "testing"
  "math/rand"
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

func TestSort(t *testing.T) {
  cards := New(Sort(Less))
  exp := Card{Rank: Ace, Suit: Spade}
  if cards[0] != exp {
    t.Error("Expected Ace of Spades as first card. Received:", cards[0])
  }
}

func TestDeck(t *testing.T) {
  cards := New(Deck(3))
  // 13 ranks * 4 suits * 3 decks
  if len(cards) != 13*4*3 {
    t.Error("Expected %d cards, received %d", 13*4*3, len(cards))
  }
}

func TestShuffle(t *testing.T) {
  // make shuffleRand deterministic
  shuffleRand = rand.New(rand.NewSource(0))

  orig := New()
  cards := New(Shuffle)
  first := orig[40]
  second := orig[35]

  if cards[0] != first {
    t.Errorf("Expected the first card to be %s, received %s", first, cards[0])
  } 
  if cards[1] != second {
    t.Errorf("Expected the second card to be %s, received %s", second, cards[1])
  } 
}

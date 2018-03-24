package deck

import "fmt"

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

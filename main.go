package main

import (
  "github.com/brianarpie/blackjack-go/deck"
  "fmt"
  "strings"
)

type Hand []deck.Card

func (h Hand) String() string {
  strs := make([]string, len(h))
  for i := range h {
    strs[i] = h[i].String()
  }
  return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
  return h[0].String() + ", **HIDDEN **"
}

func (h Hand) HasBlackjack() bool {
  hasAce := false
  hasTen := false
  if len(h) == 2 {
    for _, c := range h {
      if min(int(c.Rank), 10) == 10 {
        hasTen = true
      }
      if c.Rank == deck.Ace {
        hasAce = true
      }
    }
    return hasAce && hasTen
  }
  return false
}

func (h Hand) MinScore() int {
  score := 0
  for _, c := range h {
    score += min(int(c.Rank), 10)
  }
  return score
}

func (h Hand) MaxScore() int {
  score := 0
  for _, c := range h {
    if c.Rank == deck.Ace {
      score += 11
    } else {
      score += min(int(c.Rank), 10)
    }
  }
  return score
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

func main() {
  cards := deck.New(deck.Deck(3), deck.Shuffle)


  var player, dealer Hand = dealStartingHands(cards)
  player, dealer = playGame(cards, player, dealer)

  var card deck.Card
  for dealer.MaxScore() <= 16 || (dealer.MaxScore() == 17 && dealer.MinScore() != 17) {
    card, cards = draw(cards)
    dealer = append(dealer, card)
  }

  printResult(player, dealer)
}

func playGame(cards []deck.Card, player, dealer Hand) (Hand, Hand) {
  var card deck.Card
  var input string
  for input != "s" && player.MinScore() <= 21 {
    fmt.Println("Player:", player)
    fmt.Println("Dealer:", dealer.DealerString())
    fmt.Println("What will you do? (h)it, (s)tand")
    fmt.Scanf("%s\n", &input)
    switch input {
    case "h":
      card, cards = draw(cards)
      player = append(player, card)
    case "s":
    default:
      fmt.Println("Invalid Option")
    }
  }
  return player, dealer
}

func dealStartingHands(cards []deck.Card) (Hand, Hand) {
  var card deck.Card
  var player, dealer Hand
  for i := 0; i < 2; i++ {
    for _, hand := range []*Hand{&player, &dealer} {
      card, cards = draw(cards)
      *hand = append(*hand, card)
    }
  }
  return player, dealer
}

func printResult(player, dealer Hand) {
  pScore, dScore := player.MaxScore(), dealer.MaxScore()
  fmt.Println("==FINAL HANDS==")
  fmt.Println("Player:", player, "\nScore:", player.MinScore(), pScore)
  fmt.Println("Dealer:", dealer, "\nScore:", dealer.MinScore(), dScore)
  switch {
  case dealer.HasBlackjack():
    fmt.Println("Dealer Has Blackjack, You lose!")
  case pScore > 21:
    fmt.Println("You busted, Dealer wins!")
  case dScore > 21:
    fmt.Println("Dealer busted, You wins!")
  case pScore > dScore:
    fmt.Println("You win!")
  case dScore > pScore:
    fmt.Println("Dealer wins!")
  case dScore == pScore:
    fmt.Println("Draw. Push..")
  }
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
  return cards[0], cards[1:]
}

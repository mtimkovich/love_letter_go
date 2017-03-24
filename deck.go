package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Deck struct {
    cards [16]*Card
    index int
}

func NewDeck() *Deck {
    d := &Deck{}
    d.index = 0
    deck := &d.cards

    start := 0
    i := 0

    for i = 0; i < 5; i++ {
        deck[start + i] = Batman()
    }

    start += i

    for i = 0; i < 2; i++ {
        deck[start + i] = Catwoman()
        deck[start + i + 1] = Bane()
        deck[start + i + 2] = Robin()
        deck[start + i + 3] = PoisonIvy()
        start += i + 3
    }

    start++

    deck[start] = TwoFace()
    deck[start + 1] = HarleyQuinn()
    deck[start + 2] = Joker()

    return d
}

func (d *Deck) Shuffle() {
    for i := range d.cards {
        j := rand.Intn(i + 1)

        d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
    }
}

func (d *Deck) Draw() *Card {
    drawn := d.cards[d.index]
    d.index++

    return drawn
}

func (d *Deck) Print() {
    for i, c := range d.cards {
        fmt.Printf("%d: %s\n", i, c.name)
    }
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    deck := NewDeck()
    deck.Shuffle()
    fmt.Println(deck.Draw())
}

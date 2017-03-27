package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Deck struct {
    Cards []*Card
    index int
}

func NewDeck() *Deck {
    d := new(Deck)
    d.index = 0
    deck := make([]*Card, 0, 16)

    for i := 0; i < 5; i++ {
        deck = append(deck, Batman())
    }

    for i := 0; i < 2; i++ {
        deck = append(deck, Catwoman())
        deck = append(deck, Bane())
        deck = append(deck, Robin())
        deck = append(deck, PoisonIvy())
    }

    deck = append(deck, TwoFace())
    deck = append(deck, HarleyQuinn())
    deck = append(deck, Joker())

    d.Cards = deck
    d.Shuffle()
    return d
}

func (d *Deck) Shuffle() {
    for i := range d.Cards {
        j := rand.Intn(i + 1)

        d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
    }

    d.index = 0
}

func (d *Deck) Draw() *Card {
    if d.IsEmpty() {
        return nil
    }

    drawn := d.Cards[d.index]
    d.index++

    return drawn
}

func (d *Deck) IsEmpty() bool {
    return d.index == len(d.Cards)
}

func (d *Deck) Print() {
    for i, c := range d.Cards {
        fmt.Printf("%d: %s\n", i, c.name)
    }
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

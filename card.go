package main

import "fmt"

type Card struct {
    name string
    value int
    rules string
    action string
}

func (c *Card) String() string {
    return fmt.Sprintf("%s(%d)", c.name, c.value)
}

func Batman() *Card {
    c := new(Card)
    c.name = "Batman"
    c.value = 1
    c.rules = "Guess a player's hand"
    c.action = "guess"

    return c
}

func Catwoman() *Card {
    c := new(Card)
    c.name = "Catwoman"
    c.value = 2
    c.rules = "Look at a hand"
    c.action = "look"

    return c
}

func Bane() *Card {
    c := new(Card)
    c.name = "Bane"
    c.value = 3
    c.rules = "Compare hands; lower hand is out"
    c.action = "compare"

    return c
}

func Robin() *Card {
    c := new(Card)
    c.name = "Robin"
    c.value = 4
    c.rules = "Protection until next turn"
    c.action = "immune"

    return c
}

func PoisonIvy() *Card {
    c := new(Card)
    c.name = "Poison Ivy"
    c.value = 5
    c.rules = "One player discards their hand"
    c.action = "discard"

    return c
}

func TwoFace() *Card {
    c := new(Card)
    c.name = "Two-Face"
    c.value = 6
    c.rules = "Trade hands"
    c.action = "trade"

    return c
}

func HarleyQuinn() *Card {
    c := new(Card)
    c.name = "Harley Quinn"
    c.value = 7
    c.rules = "Discard if caught with TWO-FACE or POISON IVY"
    c.action = "nop"

    return c
}

func Joker() *Card {
    c := new(Card)
    c.name = "Joker"
    c.value = 8
    c.rules = "Lose if discarded"
    c.action = "lose"

    return c
}

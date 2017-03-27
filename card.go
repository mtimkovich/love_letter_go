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
    return &Card{
        "Batman",
        1,
        "Guess a player's hand",
        "guess",
    }
}

func Catwoman() *Card {
    return &Card{
        "Catwoman",
        2,
        "Look at a hand",
        "look",
    }
}

func Bane() *Card {
    return &Card{
        "Bane",
        3,
        "Compare hands; lower hand is out",
        "compare",
    }
}

func Robin() *Card {
    return &Card{
        "Robin",
        4,
        "Protection until next turn",
        "immune",
    }
}

func PoisonIvy() *Card {
    return &Card{
        "Poison Ivy",
        5,
        "One player discards their hand",
        "discard",
    }
}

func TwoFace() *Card {
    return &Card{
        "Two-Face",
        6,
        "Trade hands",
        "trade",
    }
}

func HarleyQuinn() *Card {
    return &Card{
        "Harley Quinn",
        7,
        "Discard if caught with TWO-FACE or POISON IVY",
        "nop",
    }
}

func Joker() *Card {
    return &Card{
        "Joker",
        8,
        "Lose if discarded",
        "lose",
    }
}

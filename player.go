package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

type Player struct {
    name string
    hand []*Card
    deck *Deck
    players []*Player
    actionMap map[string]func()
    index int
    points int
    isImmune bool
    out bool
}

func (p *Player) prompt(msg string) int {
    reader := bufio.NewReader(os.Stdin)
    idx := 0

    for {
        fmt.Printf("%s: ", msg)
        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)

        idx, err := strconv.Atoi(text)

        if err == nil && idx >= 0 && idx < len(p.players) {
            break
        }
    }

    return idx
}

func (p *Player) status(msg string) {
    fmt.Printf("Player '%s' %s\n", p.name, msg)
}

func (p *Player) Win() {
    p.points++
}

func (p *Player) Draw() {
    p.status("draws a card")
    p.hand = append(p.hand, p.deck.Draw())
}

func (p *Player) Play(cardIndex int) {

    card := p.hand[cardIndex]

    if action, e := p.actionMap[card.action]; e {
        action()
    } else {
        fmt.Printf("TODO: Impliment %s\n", card.action)
    }
}

func (p *Player) Immune() {
    p.status("is immune")
    p.isImmune = true
}

func (p *Player) Lose() {
    p.status("is out of this round")
    p.out = true
}

func (p *Player) Discard() {
    target := p.prompt("Player to make discard")
    p.players[target].makeDiscard()
}

func (p *Player) makeDiscard() {
}

func (p *Player) Nop() {
}

func NewPlayer(name string, deck *Deck) *Player {
    p := &Player{}

    p.actionMap = map[string]func(){
        "immune": p.Immune,
        "lose": p.Lose,
        "nop": n.Nop,
    }

    p.deck = deck
    p.hand = make([]*Card, 0, 2)
    p.name = name

    return p
}

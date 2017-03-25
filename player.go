package main

import (
    "bufio"
    "bytes"
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

func (p *Player) options(me bool) string {
    var options bytes.Buffer

    options.WriteRune('[')

    for i, _ := range p.players {
        if me || i != p.index {
            options.WriteRune('0' + rune(i))

            if i < len(p.players) - 1 {
                options.WriteString(", ")
            }
        }
    }

    options.WriteRune(']')

    return options.String()
}

func (p *Player) prompt(msg string, me bool) int {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Printf("%s %s: ", msg, p.options(me))
        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)

        idx, err := strconv.Atoi(text)

        if err == nil &&
           idx >= 0 &&
           idx < len(p.players) &&
           (idx != p.index || me) {
               return idx
        }
    }

}

func (p *Player) status(msg string) {
    fmt.Printf("Player %s %s\n", p.name, msg)
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
    target := p.prompt("Player to make discard", true)
    p.players[target].makeDiscard()
}

func (p *Player) makeDiscard() {
}

func (p *Player) Nop() {
}

func (p *Player) Look() {
    target := p.prompt("Player's hand to look at", false)
    p.players[target].ShowHand()
}

func (p *Player) ShowHand() {
    fmt.Printf("[%s]\n", p.hand[0].name)
}

func NewPlayer(name string, deck *Deck) *Player {
    p := &Player{}

    p.actionMap = map[string]func(){
        "immune": p.Immune,
        "lose": p.Lose,
        "look": p.Look,
        "nop": p.Nop,
    }

    p.deck = deck
    p.hand = make([]*Card, 0, 2)
    p.name = name

    return p
}

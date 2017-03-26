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

func remove(a []*Card, i int) []*Card {
    copy(a[i:], a[i+1:])
    a[len(a)-1] = nil
    return a[:len(a)-1]
}

func (p *Player) options(me bool) string {
    var options bytes.Buffer
    numeros := make([]string, 0, len(p.players))

    options.WriteRune('[')

    for i, _ := range p.players {
        if (me || i != p.index) {
            numeros = append(numeros, strconv.Itoa(i))
        }
    }

    options.WriteString(strings.Join(numeros, ", "))
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

func strRange(end int) string {
    arr := make([]string, 0)
    for i := 0; i < end; i++ {
        arr = append(arr, strconv.Itoa(i))
    }

    return strings.Join(arr, ", ")
}

func (p *Player) promptNum(msg string, end int) int {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Printf("%s [%s]: ", msg, strRange(end))
        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)

        num, err := strconv.Atoi(text)

        if err == nil &&
        num >= 0 &&
        num < end {
            return num
        }
    }
}

func (p *Player) status(msg string) {
    fmt.Printf("Player %s %s\n", p.name, msg)
}

func (p *Player) Draw() {
    p.status("draws a card")
    p.hand = append(p.hand, p.deck.Draw())
}

func (p *Player) Play() {
    cardIndex := p.promptNum("Card to play", 2)
    card := p.hand[cardIndex]

    if action, e := p.actionMap[card.action]; e {
        p.hand = remove(p.hand, cardIndex)
        action()
    } else {
        fmt.Printf("TODO: Implement %s\n", card.action)
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
    other := p.players[target]

    if other.isImmune {
        other.status("is immune")
        return
    }

    other.makeDiscard()
}

func (p *Player) makeDiscard() {
    card := p.hand[0]
    p.status(fmt.Sprintf("discarded %s", card.String()))

    if card.name == "Joker" {
        p.Lose()
    }

    p.hand = remove(p.hand, 0)

    if !p.out {
        p.Draw()
    }
}

func (p *Player) Nop() {}

func (p *Player) Look() {
    target := p.prompt("Player's hand to look at", false)
    other := p.players[target]

    if other.isImmune {
        other.status("is immune")
        return
    }

    other.ShowHand()
}

func (p *Player) Compare() {
    target := p.prompt("Player whose hand to compare to", false)
    other := p.players[target]

    if other.isImmune {
        other.status("is immune")
        return
    }

    fmt.Printf("Player %s has %s, Player %s has %s\n",
               p.name, p.hand[0].String(), other.name, other.hand[0].String())

    c := p.hand[0].value - other.hand[0].value

    if c > 0 {
        other.Lose()
    } else if c < 0 {
        p.Lose()
    } else {
        fmt.Println("Tie! Both players are still in")
    }
}

func (p *Player) Trade() {
    target := p.prompt("Player to trade hands with", false)
    other := p.players[target]

    if other.isImmune {
        other.status("is immune")
        return
    }

    p.hand[0], other.hand[0] = other.hand[0], p.hand[0]

    fmt.Print("New hand: ")
    p.ShowHand()
}

func (p *Player) Guess() {
    target := p.prompt("Player to guess", false)
    other := p.players[target]
    guessNum := p.promptNum("Guess card in player's hand", 9)

    if guessNum == other.hand[0].value {
        p.status(fmt.Sprintf("guessed correctly; Player %s is out", other.name))
        other.out = true
    }
}

func (p *Player) ShowHand() {
    fmt.Print("[")
    for i, c := range p.hand {
        fmt.Printf("%d: %s", i, c.String())

        if i < len(p.hand) - 1 {
            fmt.Print(", ")
        }
    }
    fmt.Println("]")
}

func NewPlayer(name string, deck *Deck) *Player {
    p := &Player{}

    p.actionMap = map[string]func(){
        "immune": p.Immune,
        "lose": p.Lose,
        "look": p.Look,
        "nop": p.Nop,
        "compare": p.Compare,
        "trade": p.Trade,
        "discard": p.Discard,
        "guess": p.Guess,
    }

    p.deck = deck
    p.hand = make([]*Card, 0, 2)
    p.name = name

    return p
}

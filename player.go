package main

import (
    "bufio"
    "bytes"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
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
    cpu bool
}

func remove(a []*Card, i int) []*Card {
    copy(a[i:], a[i+1:])
    a[len(a)-1] = nil
    return a[:len(a)-1]
}

func intInSlice(i int, arr []int) bool {
    for _, v := range arr {
        if v == i {
            return true
        }
    }

    return false
}

func choice(arr []int) int {
    return arr[rand.Intn(len(arr))]
}

func joinInts(arr []int, sep string) string {
    var output bytes.Buffer

    for i, val := range arr {
        output.WriteString(strconv.Itoa(val))

        if i < len(arr) - 1 {
            output.WriteString(sep)
        }
    }

    return output.String()
}

func (p *Player) Name() string {
    return fmt.Sprintf("Player %s", p.name)
}

func (p *Player) options(me bool) []int {
    output := make([]int, 0, len(p.players))

    for i, other := range p.players {
        if (me || i != p.index) && !other.out {
            output = append(output, i)
        }
    }

    return output
}

func intRange(start int, end int) []int {
    arr := make([]int, 0)
    for i := start; i < end; i++ {
        arr = append(arr, i)
    }

    return arr
}

func (p *Player) prompt(msg string, options []int) int {
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Printf("%s [%s]: ", msg, joinInts(options, ", "))

        if p.cpu {
            o := choice(options)
            fmt.Println(o)
            return o
        }

        text, _ := reader.ReadString('\n')
        text = strings.Replace(text, "\n", "", -1)

        idx, err := strconv.Atoi(text)

        if err == nil && intInSlice(idx, options) {
            return idx
        }
    }
}

func (p *Player) promptPlayer(msg string, me bool) int {
    options := p.options(me)

    for {
        idx := p.prompt(msg, options)

        if me || idx != p.index {
            return idx
        }
    }
}

func (p *Player) promptCard(msg string, start int, end int) int {
    options := intRange(start, end)
    return p.prompt(msg, options)
}

func (p *Player) status(msg string) {
    fmt.Printf("%s %s\n", p.Name(), msg)
}

func (p *Player) Win() {
    p.status("wins")
    p.points++
}

func (p *Player) Draw() bool {
    p.status("draws a card")
    card := p.deck.Draw()

    if card == nil {
        return false
    }

    p.hand = append(p.hand, card)
    return true
}

func (p *Player) Play() {
    cardIndex := p.promptCard("Card to play", 0, 2)
    card := p.hand[cardIndex]
    p.status(fmt.Sprintf("played %s", card.String()))

    action := p.actionMap[card.action]
    p.hand = remove(p.hand, cardIndex)
    action()
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
    target := p.promptPlayer("Player to make discard", true)
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
    target := p.promptPlayer("Player's hand to look at", false)
    other := p.players[target]

    if other.isImmune {
        other.status("is immune")
        return
    }

    other.status(other.ShowHand())
}

func (p *Player) Compare() {
    target := p.promptPlayer("Player whose hand to compare to", false)
    other := p.players[target]

    if other.isImmune {
        other.status("is immune")
        return
    }

    fmt.Printf("%s has %s, %s has %s\n",
               p.Name(), p.hand[0].String(), other.Name(), other.hand[0].String())

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
    target := p.promptPlayer("Player to trade hands with", false)
    other := p.players[target]

    if other.isImmune {
        other.status("is immune")
        return
    }

    p.hand[0], other.hand[0] = other.hand[0], p.hand[0]

    fmt.Println("New hand: ", p.ShowHand())
}

func (p *Player) Guess() {
    target := p.promptPlayer("Player to guess", false)
    other := p.players[target]
    guessNum := p.promptCard("Guess card in player's hand", 2, 9)
    fmt.Printf("%s had %s\n", other.Name(), other.hand[0].String()) 

    if guessNum == other.hand[0].value {
        p.status(fmt.Sprintf("guessed correctly; %s is out", other.Name()))
        other.out = true
    }
}

func (p *Player) ShowHand() string {
    var output bytes.Buffer
    output.WriteRune('{')
    for i, c := range p.hand {
        output.WriteString(fmt.Sprintf("%d: %s", i, c.String()))

        if i < len(p.hand) - 1 {
            output.WriteString(", ")
        }
    }

    output.WriteRune('}')

    return output.String()
}

func NewPlayer(name string, deck *Deck) *Player {
    p := new(Player)

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

func init() {
    rand.Seed(time.Now().UnixNano())
}

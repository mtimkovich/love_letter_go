package main

type Player struct {
    name string
    hand []*Card
    deck []*Card
    points int
    is_immune bool
    out bool
}

func (p *Player) win() {
    p.points++
}

func (p *Player) draw() *Card {
    return p.deck.Draw()
}

func NewPlayer(name string, deck []*Card) *Player {
    p = &Player{}

    p.deck = deck
    p.hand = make([]*Card, 0, 2)
    p.name = name

    return p
}

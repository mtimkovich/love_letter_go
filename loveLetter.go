package main

// import "fmt"

func main() {
    deck := NewDeck()
    // deck.Draw()
    deck.Cards[0] = PoisonIvy()
    deck.Cards[1] = Joker()

    players := make([]*Player, 0, 4)

    players = append(players, NewPlayer("one", deck))
    players = append(players, NewPlayer("two", deck))

    for i, p := range players {
        p.players = players
        p.index = i
        p.Draw()
    }

    one := players[0]
    two := players[1]

    one.Draw()
    two.Immune()

    one.ShowHand()
    one.Play()
}

package main

import "fmt"

func playersOut(players []*Player) int {
    out := 0
    for _, p := range players {
        if p.out {
            out++
        }
    }

    return out
}

func main() {
    deck := NewDeck()

    // Simulate burning the top card
    deck.Draw()

    players := make([]*Player, 0, 4)

    for _, name := range []string{"One", "Two", "Three", "Four"} {
        players = append(players, NewPlayer(name, deck))
    }

    for i, p := range players {
        p.players = players
        p.index = i
        p.cpu = true
        p.Draw()
    }

    players[0].cpu = false

    done := false

    for !done && !deck.IsEmpty() {
        for _, p := range players {
            if p.out {
                continue
            }

            if playersOut(players) == len(players) - 1 {
                p.Win()
                done = true
                break
            }

            p.isImmune = false
            fmt.Println()
            success := p.Draw()

            if !success {
                fmt.Println("Deck is empty")
                break
            }

            fmt.Printf("%s: %s\n", p.Name(), p.ShowHand())
            p.Play()
        }
    }

    // TODO: Implement deck empty end of game conditions
}

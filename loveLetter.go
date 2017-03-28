package main

import (
    "fmt"
    "strings"
)

func playersOut(players []*Player) int {
    out := 0
    for _, p := range players {
        if p.out {
            out++
        }
    }

    return out
}

func maxCardValue(players []*Player) int {
    max := 0
    for _, p := range players {
        if len(p.hand) > 0 {
            value := p.hand[0].value

            if !p.out && value > max {
                max = value
            }
        }
    }

    return max
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

    // players[0].cpu = false

    winners := make([]string, 0, 4)

    for len(winners) == 0 && !deck.IsEmpty() {
        for _, p := range players {
            if p.out {
                continue
            }

            if playersOut(players) == len(players) - 1 {
                p.points++
                winners = append(winners, p.Name())
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

    if len(winners) == 0 {
        fmt.Println()

        winningValue := maxCardValue(players)

        for _, p := range players {
            if !p.out {
                // TODO Sometimes players are not out, but also
                // have empty hands, which shouldn't be possible
                fmt.Println(p.Name(), "has", p.hand[0].String())

                if p.hand[0].value == winningValue {
                    winners = append(winners, p.Name())
                }
            }
        }
    }

    fmt.Println()
    if len(winners) == 1 {
        fmt.Println("Winner is", winners[0])
    } else {
        fmt.Println("Winners are", strings.Join(winners, " and "))
    }
}

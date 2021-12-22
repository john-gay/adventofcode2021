package main

import (
	"fmt"
	"log"
	"time"
)

type game struct {
	p1 player
	p2 player
}

type player struct {
	position int
	score    int
	wins     int
}

var winMap = make(map[game][]int64)

func main() {
	start := time.Now()

	g := game{
		p1: player{
			position: 3,
			score:    0,
		},
		p2: player{
			position: 5,
			score:    0,
		},
	}

	wins := g.play()

	fmt.Println(fmt.Sprintf("%+v", wins))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (g *game) play() []int64 {
	if g.p1.score >= 21 {
		return []int64{1, 0}
	}
	if g.p2.score >= 21 {
		return []int64{0, 1}
	}
	if w, ok := winMap[*g]; ok {
		return w
	}

	win := []int64{0, 0}
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				p1Pos := nextPosition(g.p1.position, i+j+k)
				p1Score := g.p1.score + p1Pos

				g2 := game{
					p1: player{
						position: g.p2.position,
						score:    g.p2.score,
					},
					p2: player{
						position: p1Pos,
						score:    p1Score,
					},
				}
				win2 := g2.play()

				win[0] += win2[1]
				win[1] += win2[0]
			}
		}
	}
	winMap[*g] = win
	return win
}

func (g *game) playPracticeRound() {

}

func nextPosition(pos, move int) int {
	next := pos + move

	for next > 10 {
		next -= 10
	}
	return next
}

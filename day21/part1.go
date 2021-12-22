package main

import (
	"fmt"
	"log"
	"time"
)

type game struct {
	p1       player
	p2       player
	die      int
	rolls    int
	complete bool
}

type player struct {
	position int
	score    int
}

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
		die:      1,
		rolls:    0,
		complete: false,
	}

	for {
		g.playPracticeRound()

		if g.complete {
			break
		}
	}

	fmt.Println(fmt.Sprintf("Part 1: %d", g.loosingScore()))

	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}

func (g *game) playPracticeRound() {
	g.p1.practiceMove(g)
	if !g.complete {
		g.p2.practiceMove(g)
	}
}

func (g *game) nextPracticeMove() int {
	move := 0
	for i := 0; i < 3; i++ {
		g.rolls += 1
		move += g.die

		if g.die < 100 {
			g.die++
		} else {
			g.die = 1
		}
	}

	return move
}

func (g *game) loosingScore() int {
	score := g.p1.score
	if score >= 1000 {
		score = g.p2.score
	}

	return score * g.rolls
}

func (p *player) practiceMove(g *game) {
	move := g.nextPracticeMove()

	nextPosition := p.position + move

	for nextPosition > 10 {
		nextPosition -= 10
	}

	p.position = nextPosition
	p.score += nextPosition

	if p.score >= 1000 {
		g.complete = true
	}
}

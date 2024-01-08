package main

import "math/rand"

type player struct {
	id    int
	hold  int
	score int
}

func (p *player) turn() {
	var turnScore int
	toScore := findMin(p.hold, 100-p.score)
	for turnScore < toScore {
		roll := rand.Intn(5) + 1
		if roll == 1 {
			return
		} else {
			turnScore += roll
		}
	}
	p.score += turnScore
}

func (p *player) reset() {
	p.score = 0
}

package eoc

import (
	"fmt"
	"math/rand"
)

type Player interface {
	PlayWith(p Player) bool
	JoinArena(a *Arena)
	ID() int64
	Name() string
	ReceiveMatchResult(playerID int64, coop bool)
}

type PlayerA struct {
	id     int64
	memory map[int64][]bool
}

type Random struct {
	id int64
}

func (p *Random) ID() int64 {
	return p.id
}

func (p *Random) JoinArena(a *Arena) {
	p.id = a.NewPlayerID()
}

func (p *Random) Name() string {
	return fmt.Sprintf("random_%d", p.ID())
}

func (p *Random) PlayWith(p2 Player) bool {
	if rand.Intn(10) > 5 {
		return true
	}
	return false
}

func (p *Random) ReceiveMatchResult(playerID int64, coop bool) {
	return
}

// Fish has no memory, always coop/betray
type Fish struct {
	id   int64
	coop bool
}

func NewFish(coop bool) *Fish {
	return &Fish{
		coop: coop,
	}
}

func (p *Fish) PlayWith(p2 Player) bool {
	return p.coop
}

func (p *Fish) JoinArena(a *Arena) {
	p.id = a.NewPlayerID()
}

func (p *Fish) ID() int64 {
	return p.id
}

func (p *Fish) Name() string {
	position := "bad"
	if p.coop {
		position = "good"
	}
	return fmt.Sprintf("fish_%s_%d", position, p.ID())
}

func (p *Fish) ReceiveMatchResult(playerID int64, coop bool) {
	return
}

type Tic4Tac struct {
	id     int64
	memory map[int64][]bool
}

func NewTic4TacPlayer() *Tic4Tac {
	return &Tic4Tac{
		memory: make(map[int64][]bool),
	}
}

func (p *Tic4Tac) ID() int64 {
	return p.id
}

func (p *Tic4Tac) Name() string {
	return fmt.Sprintf("tic4tac_%d", p.ID())
}

func (p *Tic4Tac) PlayWith(p2 Player) bool {
	if hist, ok := p.memory[p2.ID()]; ok == true {
		if len(hist) > 0 {
			if !hist[len(hist)-1] {
				return false
			}
		}
	}
	return true
}

func (p *Tic4Tac) ReceiveMatchResult(playerID int64, coop bool) {
	if _, ok := p.memory[playerID]; ok == false {
		p.memory[playerID] = make([]bool, 0)
	}
	p.memory[playerID] = append(p.memory[playerID], coop)
}

func (p *Tic4Tac) JoinArena(a *Arena) {
	p.id = a.NewPlayerID()
}

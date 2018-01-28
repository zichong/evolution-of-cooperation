package eoc

import "math/rand"

type Player interface {
	PlayWith(p Player) bool
	JoinArena(a *Arena)
	ID() int64
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

func (p *Random) PlayWith(p2 Player) bool {
	if rand.Intn(10) > 5 {
		return true
	}
	return false
}

func (p *Random) ReceiveMatchResult(playerID int64, coop bool) {
	return
}

type Fish struct {
	// fish has no memory
}

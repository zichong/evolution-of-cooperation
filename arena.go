package eoc

import (
	"fmt"
	"math/rand"
)

type Arena struct {
	currentPlayerID int64
	players         []Player
	interval        int
	coopGain        int
	betrayerGain    int
	betrayedGain    int
	bothBetrayGain  int
	scores          map[int64]int
}

func NewArena() *Arena {
	return &Arena{
		currentPlayerID: 1,
		betrayedGain:    0,
		betrayerGain:    5,
		bothBetrayGain:  1,
		coopGain:        3,
		scores:          make(map[int64]int),
	}
}

func (a *Arena) AddScore(pID int64, s int) {
	a.scores[pID] += s
}

func (a *Arena) AddPlayer(player Player) {
	a.players = append(a.players, player)
	player.JoinArena(a)
}

func (a *Arena) PlayRound() {
	n := len(a.players)
	if n < 2 {
		return
	}
	a.shuffleplayers()
	for i := 0; i < n; i += 2 {
		p1, p2 := a.players[i], a.players[i+1]
		a.Match(p1, p2)
	}
}

func (a *Arena) NewPlayerID() int64 {
	a.currentPlayerID += (rand.Int63n(5) + 1)
	return a.currentPlayerID
}

func (a *Arena) Match(p1 Player, p2 Player) {
	p1Coop := p1.PlayWith(p2)
	p2Coop := p2.PlayWith(p1)
	if p1Coop == p2Coop {
		var gain int
		if p1Coop {
			gain = a.coopGain
		} else {
			gain = a.bothBetrayGain
		}
		a.AddScore(p1.ID(), gain)
		a.AddScore(p2.ID(), gain)
	}
	if p1Coop && !p2Coop {
		a.AddScore(p1.ID(), a.betrayedGain)
		a.AddScore(p2.ID(), a.betrayerGain)
	}
	if !p1Coop && p2Coop {
		a.AddScore(p1.ID(), a.betrayerGain)
		a.AddScore(p2.ID(), a.betrayedGain)
	}
	a.acknowledge(p1, p2.ID(), p2Coop)
	a.acknowledge(p2, p1.ID(), p1Coop)

	var p1Action, p2Action string
	if p1Coop {
		p1Action = "coop"
	} else {
		p1Action = "betray"
	}
	if p2Coop {
		p2Action = "coop"
	} else {
		p2Action = "betry"
	}
	fmt.Printf("player %d: %s, player %d: %s\n", p1.ID(), p1Action, p2.ID(), p2Action)
}

func (a *Arena) acknowledge(p Player, matchedPlayerID int64, coop bool) {
	p.ReceiveMatchResult(matchedPlayerID, coop)
}

func (a *Arena) shuffleplayers() {
	n := len(a.players)
	for i := n - 1; i > 0; i-- {
		j := rand.Intn(i)
		a.players[i], a.players[j] = a.players[j], a.players[i]
	}
}

func (a *Arena) LeaderBoard() map[int64]int {
	return a.scores
}

package main

import (
	"fmt"
	"math/rand"

	eoc "github.com/zichong/evolution-of-cooperation"
)

func main() {
	arena := eoc.NewArena()
	for i := 0; i < 10; i++ {
		p := &eoc.Random{}
		arena.AddPlayer(p)
	}
	for i := 0; i < 20; i++ {
		x := rand.Float64()
		p := eoc.NewFish(x > 0.5)
		arena.AddPlayer(p)
	}
	for i := 0; i < 10; i++ {
		p := eoc.NewTic4TacPlayer()
		arena.AddPlayer(p)
	}

	for i := 0; i < 10000; i++ {
		arena.PlayRound()
	}
	fmt.Println(arena.LeaderBoard())
}

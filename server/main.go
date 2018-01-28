package main

import (
	"eoc"
	"fmt"
)

func main() {
	arena := eoc.NewArena()
	p1 := &eoc.Random{}
	p2 := &eoc.Random{}

	arena.AddPlayer(p1)
	arena.AddPlayer(p2)

	arena.PlayRound()
	arena.PlayRound()
	fmt.Println(arena.LeaderBoard())

	arena.PlayRound()
	arena.PlayRound()
	fmt.Println(arena.LeaderBoard())
}

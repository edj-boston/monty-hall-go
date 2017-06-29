package game

import (
	"math/rand"
	"time"
)

type Game struct {
	Car    int
	Guest  int
	Host   int
	Change bool
	Win    bool
}

func (g *Game) Setup() *Game {
	rand.Seed(time.Now().UTC().UnixNano())
	g.Car = rand.Intn(3)
	return g
}

func (g *Game) Choose(s string) *Game {
	switch s {
	case "guest":
		g.Guest = rand.Intn(3)
	case "host":
		g.Host = rand.Intn(3)
		for g.Host == g.Guest || g.Host == g.Car {
			g.Host = rand.Intn(3)
		}
	}

	return g
}

func (g *Game) Decide() *Game {
	if rand.Intn(2) == 1 {
		g.Change = true
		choice := rand.Intn(3)
		for choice == g.Guest || choice == g.Host {
			choice = rand.Intn(3)
		}
		g.Guest = choice
	} else {
		g.Change = false
	}

	if g.Guest == g.Car {
		g.Win = true
	} else {
		g.Win = false
	}

	return g
}

func (g *Game) Play() {
	g.Setup().Choose("guest").Choose("host").Decide()
}

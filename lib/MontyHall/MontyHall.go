package montyhall

import (
	"github.com/edj-boston/monty-hall-go/lib/Game"
)

type MontyHall struct {
	Results struct {
		Total            int
		Changes          int
		ChangeWins       int
		ChangeWinPercent float64
		Stays            int
		StayWins         int
		StayWinPercent   float64
	}
	Games []game.Game
}

func (mh *MontyHall) Run() {
	for i := 0; i < 100000; i++ {
		g := game.Game{}
		g.Play()
		mh.Games = append(mh.Games, g)

		mh.Results.Total++

		if g.Change {
			mh.Results.Changes++
			if g.Win {
				mh.Results.ChangeWins++
			}
		} else {
			mh.Results.Stays++
			if g.Win {
				mh.Results.StayWins++
			}
		}
	}

	mh.Results.ChangeWinPercent = float64(mh.Results.ChangeWins) / float64(mh.Results.Changes) * 100
	mh.Results.StayWinPercent = float64(mh.Results.StayWins) / float64(mh.Results.Stays) * 100
}

package montyhall

import "testing"

func TestMontyHall(t *testing.T) {
	mh := MontyHall{}

	mh.Run()
	if mh.Results.Changes+mh.Results.Stays != mh.Results.Total {
		t.Error("The number of change wins plus stay wins should equal the total", mh.Results.Changes, mh.Results.Stays, mh.Results.Total)
	}
}

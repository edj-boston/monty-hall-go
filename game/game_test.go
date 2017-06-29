package game

import "testing"

func TestGameSetup(t *testing.T) {
	g := Game{}

	for i := 0; i < 100000; i++ {
		g.Setup()
		if g.Car < 0 || g.Car > 2 {
			t.Error("The car field must be 0, 1, or 2", g.Car)
		}
	}
}

func TestGameChoose(t *testing.T) {
	g := Game{}

	for i := 0; i < 100000; i++ {
		g.Setup().Choose("guest").Choose("host")
		if g.Guest == g.Host {
			t.Error("The guest and host must choose different doors", g.Guest, g.Host)
		}
	}
}

func TestGameDecide(t *testing.T) {
	g := Game{}

	for i := 0; i < 100000; i++ {
		g.Setup().Choose("guest").Choose("host")
		choice := g.Guest
		g.Decide()

		if g.Change == true {
			if choice == g.Guest {
				t.Error("The guest choice should be different if a decision to change is made", choice, g.Guest)
			}
		} else {
			if choice != g.Guest {
				t.Error("The guest choice should stay the same if a decision to not change is made", choice, g.Guest)
			}
		}

		if g.Guest == g.Car && g.Win == false {
			t.Error("The game should be scored a win if the guest choose the car", g.Guest, g.Car, g.Win)
		}
	}
}

func TestGamePlay(t *testing.T) {
	g := Game{}

	for i := 0; i < 100000; i++ {
		g.Play()
		if g.Guest != g.Car && g.Win {
			t.Error("The Play() method should produce a valid game", g.Guest, g.Car, g.Win)
		}
	}
}

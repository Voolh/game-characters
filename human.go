package main

type human struct {
}

func (h *human) makeMage() iMage {
	return &humanMage{
		mage: mage{
			name:   "",
			health: 80,
			mana:   80,
		},
	}
}

func (h *human) makeWarrior() iWarrior {
	return &humanWarrior{
		warrior: warrior{
			name:   "",
			health: 100,
			rage:   100,
		},
	}
}

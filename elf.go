package main

type elf struct {
}

func (e *elf) makeMage() iMage {
	return &elfMage{
		mage: mage{
			name:   "",
			health: 70,
			mana:   100,
		},
	}
}

func (e *elf) makeWarrior() iWarrior {
	return &elfWarrior{
		warrior: warrior{
			name:   "",
			health: 80,
			rage:   80,
		},
	}
}

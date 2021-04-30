package main

import "fmt"

type iRaceFactory interface {
	makeMage() iMage
	makeWarrior() iWarrior
}

func getRaceFactory(race string) (iRaceFactory, error) {
	if race == "human" {
		return &human{}, nil
	}

	if race == "elf" {
		return &elf{}, nil
	}

	return nil, fmt.Errorf("Wrong race selected")
}

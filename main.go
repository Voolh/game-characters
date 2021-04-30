package main

import (
	"fmt"
)

func main() {
	humanFactory, _ := getRaceFactory("human")
	elfFactory, _ := getRaceFactory("elf")

	humanMage := humanFactory.makeMage()
	humanWarrior := humanFactory.makeWarrior()

	elfMage := elfFactory.makeMage()
	elfWarrior := elfFactory.makeWarrior()

	printMageStats(humanMage)
	printWarriorStats(humanWarrior)

	printMageStats(elfMage)
	printWarriorStats(elfWarrior)
}

func printMageStats(s iMage) {
	fmt.Printf("Name: %s", s.Name())
	fmt.Println()
	fmt.Printf("Health Points: %d", s.Health())
	fmt.Println()
	fmt.Printf("Mana Points: %d", s.Mana())
	fmt.Println()
}

func printWarriorStats(s iWarrior) {
	fmt.Printf("Name: %s", s.Name())
	fmt.Println()
	fmt.Printf("Health Points: %d", s.Health())
	fmt.Println()
	fmt.Printf("Rage Points: %d", s.Rage())
	fmt.Println()
}

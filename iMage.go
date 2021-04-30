package main

type iMage interface {
	SetName(name string)
	Name() string
	SetHealth(health int)
	Health() int
	SetMana(mana int)
	Mana() int
}

type mage struct {
	name   string
	health int
	mana   int
}

func (m *mage) Name() string {
	return m.name
}

func (m *mage) SetName(name string) {
	m.name = name
}

func (m *mage) Health() int {
	return m.health
}

func (m *mage) SetHealth(health int) {
	m.health = health
}

func (m *mage) Mana() int {
	return m.mana
}

func (m *mage) SetMana(mana int) {
	m.mana = mana
}

package main

type iWarrior interface {
	SetName(name string)
	Name() string
	SetHealth(health int)
	Health() int
	SetRage(mana int)
	Rage() int
}

type warrior struct {
	name   string
	health int
	rage   int
}

func (w *warrior) Name() string {
	return w.name
}

func (w *warrior) SetName(name string) {
	w.name = name
}

func (w *warrior) Health() int {
	return w.health
}

func (w *warrior) SetHealth(health int) {
	w.health = health
}

func (w *warrior) Rage() int {
	return w.rage
}

func (w *warrior) SetRage(rage int) {
	w.rage = rage
}

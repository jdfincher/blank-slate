// Package gameobj implements the weapons, armor and other items
package gameobj

type Weapon struct {
	Name     string
	Price    int
	Damage   int
	Range    int
	Reload   int
	Critical int
	AimBonus int
	Capacity int
	Weight   int
	Count    int
}

type Armor struct {
	Name         string
	Price        int
	AC           int
	PassiveCheck int
	Weight       int
	Helmet       bool
	Body         bool
	Count        int
}

type Item struct {
	Name   string
	Info   string
	Weight int
	Count  int
}

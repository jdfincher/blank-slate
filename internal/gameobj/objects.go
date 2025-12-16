// Package gameobj implements the weapons, armor and other items
package gameobj

type Weapon struct {
	Name     string
	Price    int
	DmgMin   int
	DmgMax   int
	Range    int
	Reload   int
	Capacity int
	Critical int
	AimBonus int
	Caliber  int
	Count    int
	Type     string
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

// Package gameobj implements the weapons, armor and other items
package gameobj

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

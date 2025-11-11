// Package player implements the structs for the player
package player

import (
	"github.com/jdfincher/blank-slate/internal/gameobj"
)

type Player struct {
	Name      string
	Race      string
	Type      string
	Stats     PlayerStats
	Inventory PlayerInventory
	Skills    []Skill
	Methods   []Method
}

type PlayerStats struct {
	Level       int
	Life        int
	Clicks      int
	Money       int
	Strength    int
	Dexterity   int
	Fitness     int
	Agility     int
	Wisdom      int
	Intellect   int
	Charisma    int
	Appearance  int
	Defense     int
	DamageRange gameobj.DamageRange
}

type PlayerStatMods struct {
	Strength   int
	Dexterity  int
	Fitness    int
	Agility    int
	Wisdom     int
	Intellect  int
	Charisma   int
	Appearance int
}

type PlayerInventory struct {
	Armor  []gameobj.Armor
	Weapon []gameobj.Weapon
	Item   []gameobj.Item
}

type Equipped struct {
	BodyArmor gameobj.Armor
	Helmet    gameobj.Armor
	RightHand gameobj.Weapon
	LeftHand  gameobj.Weapon
}

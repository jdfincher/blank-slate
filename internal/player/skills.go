package player

import (
	"reflect"
)

var (
	Gut = Skill{
		Name:      "Gut",
		Info:      "Trust your gut to never let you down.",
		Level:     1,
		StatKey:   "Wisdom",
		StatBonus: 0,
		ViewState: false,
	}
	Sneak = Skill{
		Name:      "Sneak",
		Info:      "Move like a panther stalking its prey.",
		Level:     1,
		StatKey:   "Agility",
		StatBonus: 0,
		ViewState: false,
	}
	Climb = Skill{
		Name:      "Climb",
		Info:      "Scale objects with ease where others falter.",
		Level:     1,
		StatKey:   "Fitness",
		StatBonus: 0,
		ViewState: false,
	}
	Spot = Skill{
		Name:      "Spot",
		Info:      "A keen eye for opportunity and danger.",
		Level:     1,
		StatKey:   "Intellect",
		StatBonus: 0,
		ViewState: false,
	}
	Lore = Skill{
		Name:      "Lore",
		Info:      "The stories of the dark have never left you.",
		Level:     1,
		StatKey:   "Wisdom",
		StatBonus: 0,
		ViewState: false,
	}
	Hide = Skill{
		Name:      "Hide",
		Info:      "Not all foes should be faced head on.",
		Level:     1,
		StatKey:   "Dexterity",
		StatBonus: 0,
		ViewState: false,
	}
	Listen = Skill{
		Name:      "Listen",
		Info:      "Hear them before they hear you.",
		Level:     1,
		StatKey:   "Wisdom",
		StatBonus: 0,
		ViewState: false,
	}
	Glib = Skill{
		Name:      "Glib",
		Info:      "Say what they want to hear.",
		Level:     1,
		StatKey:   "Charisma",
		StatBonus: 0,
		ViewState: false,
	}
	StockSkillSet = &[]Skill{Gut, Sneak, Climb, Spot, Lore, Hide, Listen, Glib}
)

type Skill struct {
	Name      string
	Info      string
	Level     int
	StatKey   string
	StatBonus int
	ViewState bool
}

func (s *Skill) SetSkillStatBonus(key string, pM *PlayerStatMods) {
	rvPlayer := reflect.ValueOf(pM).Elem()
	fieldValuePlayer := rvPlayer.FieldByName(key)

	s.StatBonus = int(fieldValuePlayer.Int())
}

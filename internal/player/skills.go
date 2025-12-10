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
		StatKey:   "Strength",
		StatBonus: 0,
		ViewState: false,
	}
	Skill3 = Skill{
		Name:      "Skill3",
		Info:      "Skill placeholder index number three",
		Level:     1,
		StatKey:   "Dexterity",
		StatBonus: 0,
		ViewState: false,
	}
	Skill4 = Skill{
		Name:      "Skill4",
		Info:      "Skill placeholder index 4",
		Level:     1,
		StatKey:   "Intellect",
		StatBonus: 0,
		ViewState: false,
	}
	Skill5 = Skill{
		Name:      "Skill5",
		Info:      "Skill placeholder index 5",
		Level:     1,
		StatKey:   "Appearance",
		StatBonus: 0,
		ViewState: false,
	}
	Skill6 = Skill{
		Name:      "Skill6",
		Info:      "Skill placeholder index 6",
		Level:     1,
		StatKey:   "Charisma",
		StatBonus: 0,
		ViewState: false,
	}
	Skill7 = Skill{
		Name:      "Skill7",
		Info:      "Skill placeholder index 7",
		Level:     1,
		StatKey:   "Fitness",
		StatBonus: 0,
		ViewState: false,
	}
	StockSkillSet = &[]Skill{Gut, Sneak, Climb, Skill3, Skill4, Skill5, Skill6, Skill7}
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

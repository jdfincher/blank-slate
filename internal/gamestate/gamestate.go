// Package gamestate
package gamestate

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdfincher/blank-slate/internal/player"
	"github.com/jdfincher/blank-slate/internal/ui"
)

type State struct {
	Menu        bool
	Create      bool
	Idle        bool
	Combat      bool
	Player      *player.Player
	CreateState *CreateState
}

type CreateState struct {
	Name   bool
	Race   bool
	Type   bool
	Stat   bool
	Skill  bool
	Method bool
	Inv    bool
}

func (c *CreateState) Reset() {
	c.Name = true
	c.Race = false
	c.Type = false
	c.Stat = false
	c.Skill = false
	c.Method = false
	c.Inv = false
}

func (s *State) StartMenu() {
	s.Menu = true
	s.Create = false
	s.Idle = false
	s.Combat = false
	s.Player.ResetPlayer()

	if clicked := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50), 200, 50), "New Game"); clicked {
		fmt.Println("New Game Clicked")
		s.Menu = false
		s.Create = true

	} else if clicked := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+60, 200, 50), "Continue"); clicked {
		fmt.Println("Continue Clicked")
		s.Menu = false
		s.Idle = true

	}
}

func (s *State) HandleStyle() {
	if s.Menu {
		rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
	} else if s.Create {
		rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
	} else if s.Idle {
		rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
	}
}

func (s *State) NewGameCreate() {
	s.Menu = false
	s.Create = true

	if s.CreateState.Name {
		s.CreateState.Race = s.Player.NewPlayerName()
		s.CreateState.Name = !s.CreateState.Race
	} else if s.CreateState.Race {
		s.CreateState.Type = s.Player.NewPlayerRace()
		s.CreateState.Race = !s.CreateState.Type
	} else if s.CreateState.Type {
		s.CreateState.Stat = s.Player.NewPlayerType()
		s.CreateState.Type = !s.CreateState.Stat
	} else if s.CreateState.Stat {
		s.CreateState.Skill = s.Player.NewPlayerStats()
		s.CreateState.Stat = !s.CreateState.Skill
	} else if s.CreateState.Skill {
		s.CreateState.Method = s.Player.NewPlayerSkills()
		s.CreateState.Skill = !s.CreateState.Method
	} else if s.CreateState.Method {
		s.CreateState.Inv = s.Player.NewPlayerMethods()
		s.CreateState.Method = !s.CreateState.Inv
	}
	// Go To Main Menu
	if clicked := rg.Button(rl.NewRectangle(float32(rl.GetScreenWidth())-265, ui.BottomAlign(50), 200, 50), "Main Menu"); clicked {
		s.Menu = true
		s.Create = false
		s.CreateState.Reset()

	}
}

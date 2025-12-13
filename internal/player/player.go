// Package player implements the structs for the player
package player

import (
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdfincher/blank-slate/internal/gameobj"
	"github.com/jdfincher/blank-slate/internal/ui"
)

var (
	StatView          = true
	statPointView     = true
	skillView         = false
	methodView        = false
	addSkillInfoView  = false
	addMethodInfoView = false
	StatsCopy         = PlayerStats{}
	SkillSet          = []Skill{}
	MethodSet         = []Method{}
	StockMethodSet    = &[]Method{}

	Human = PlayerStats{
		Level:      1,
		Life:       20,
		Clicks:     6,
		Aim:        2,
		Interrupts: 0,
		Money:      250,
		Strength:   6,
		Dexterity:  13,
		Fitness:    10,
		Agility:    10,
		Wisdom:     8,
		Intellect:  7,
		Charisma:   7,
		Appearance: 8,
	}

	Mutant = PlayerStats{
		Level:      1,
		Life:       30,
		Clicks:     6,
		Interrupts: 0,
		Money:      100,
		Strength:   12,
		Dexterity:  7,
		Fitness:    8,
		Agility:    6,
		Wisdom:     6,
		Intellect:  5,
		Charisma:   2,
		Appearance: 1,
	}

	Cyborg = PlayerStats{
		Level:      1,
		Life:       25,
		Clicks:     6,
		Interrupts: 0,
		Money:      500,
		Strength:   8,
		Dexterity:  8,
		Fitness:    9,
		Agility:    9,
		Wisdom:     9,
		Intellect:  13,
		Charisma:   6,
		Appearance: 5,
	}
)

type Player struct {
	Name       string
	Race       string
	Type       string
	Stats      *PlayerStats
	StatMods   *PlayerStatMods
	StatPoints int
	Inventory  *PlayerInventory
	Skills     *map[string]Skill
	Methods    *map[string]Method
}

type PlayerStats struct {
	Level       int
	Life        int
	Clicks      int
	Initiative  int
	Aim         int
	Interrupts  int
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
	DamageRange *gameobj.DamageRange
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
	Armor  *[]gameobj.Armor
	Weapon *[]gameobj.Weapon
	Item   *[]gameobj.Item
}

type Equipped struct {
	BodyArmor *gameobj.Armor
	Helmet    *gameobj.Armor
	RightHand *gameobj.Weapon
	LeftHand  *gameobj.Weapon
}

func (p *Player) ResetPlayer() {
	p.Name = ""
	p.Race = ""
	p.Type = ""
	p.Stats = new(PlayerStats)
	p.StatMods = new(PlayerStatMods)
	p.StatPoints = 0
	p.NewSkillMap()
	p.NewMethodMap()
	p.CreateInventory()
	StatsCopy = p.UpdateStatsCopy()
	p.UpdateSkillSetCopy()
	p.UpdateMethodSetCopy()
	statPointView = true
}

func (p *Player) NewPlayerName() bool {
	rl.DrawText(p.Name, ui.CenterTextX(p.Name, 32), int32(ui.CenterY(50))-60, 32, rl.Red)
	p.DrawPlayerStats()
	// Next
	if clicked := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+60, 200, 50), "Next"); clicked {
		fmt.Println("Next was clicked")
		return true
	}

	rec := rl.NewRectangle(ui.CenterX(300), ui.CenterY(50), 300, 50)
	rg.TextBox(rec, &p.Name, 14, true)
	return false
}

func (p *Player) NewPlayerRace() bool {
	p.DrawPlayerStats()
	var (
		w float32 = 925
		h float32 = 200
		x         = ui.CenterX(w)
		y         = ui.CenterY(h) - h
	)

	if human := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)-60, 200, 50), "Human"); human {
		p.Race = "Human"
		p.SetRaceStats(Human)
		fmt.Printf("Player Race is %s\n", p.Race)
	} else if mutant := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50), 200, 50), "Mutant"); mutant {
		p.Race = "Mutant"
		p.SetRaceStats(Mutant)
		fmt.Printf("Player Race is %s\n", p.Race)
	} else if cyborg := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+60, 200, 50), "Cyborg"); cyborg {
		p.Race = "Cyborg"
		p.SetRaceStats(Cyborg)
		fmt.Printf("Player Race is %s\n", p.Race)
	}
	if clicked := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+130, 200, 50), "Next"); clicked {
		StatsCopy = p.UpdateStatsCopy()
		return true
	}
	// Popup Panel Race Description
	mousePos := rl.GetMousePosition()
	// Human
	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_CENTER))
	if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)-60, 200, 50)) {
		rg.Panel(rl.NewRectangle(x, y, w, h), "Human")
		rg.Label(rl.NewRectangle(x, y+75, w, h), HumanDesc)

		// Mutant
	} else if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(ui.CenterX(200), ui.CenterY(50), 200, 50)) {
		rg.Panel(rl.NewRectangle(x, y, w, h), "Mutant")
		rg.Label(rl.NewRectangle(x, y+75, w, h), MutantDesc)
		// Cyborg
	} else if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+60, 200, 50)) {
		rg.Panel(rl.NewRectangle(x, y, w, h), "Cyborg")
		rg.Label(rl.NewRectangle(x, y+75, w, h), CyborgDesc)

	}

	return false
}

func (p *Player) NewPlayerType() bool {
	p.DrawPlayerStats()
	var (
		w float32 = 945
		h float32 = 250
		x         = ui.CenterX(w)
		y         = ui.CenterY(h) - h
	)
	if gun := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)-60, 200, 50), "Gunslinger"); gun {
		p.Type = "Gunslinger"
		fmt.Printf("Player Type is %s\n", p.Type)
	} else if nat := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50), 200, 50), "Naturalist"); nat {
		p.Type = "Naturalist"
		fmt.Printf("Player Type is %s\n", p.Type)
	} else if che := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+60, 200, 50), "Chemist"); che {
		p.Type = "Chemist"
		fmt.Printf("Player Type is %s\n", p.Type)
	}
	if clicked := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+130, 200, 50), "Next"); clicked {
		return true
	}
	// Popup Panel Type Description
	mousePos := rl.GetMousePosition()
	// Gunslinger
	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_CENTER))
	if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)-60, 200, 50)) {
		rg.Panel(rl.NewRectangle(x, y, w, h), "Gunslinger")
		rg.Label(rl.NewRectangle(x, y+100, w, h), GunslingerDesc)

		// Mutant
	} else if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(ui.CenterX(200), ui.CenterY(50), 200, 50)) {
		rg.Panel(rl.NewRectangle(x, y, w, h), "Naturalist")
		rg.Label(rl.NewRectangle(x, y+100, w, h), NaturalistDesc)

		// Cyborg
	} else if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+60, 200, 50)) {
		rg.Panel(rl.NewRectangle(x, y, w, h), "Chemist")
		rg.Label(rl.NewRectangle(x, y+100, w, h), ChemistDesc)

	}

	return false
}

func (p *Player) NewPlayerStats() bool {
	p.DrawPlayerStats()

	if statPointView {
		p.DistributeStatPoints()
		p.SetStatMods()
	} else {
		p.SetStatMods()
		return true
	}

	return false
}

func (p *Player) DrawPlayerStats() {
	var (
		panelx float32 = 25
		panely float32 = 25
		panelw float32 = 400
		panelh         = float32(rl.GetScreenHeight() - 50)
	)
	// rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_CENTER))
	bounds := rl.NewRectangle(panelx, panely, panelw, panelh)
	rg.Panel(bounds, fmt.Sprintf("%s -- Level %d", p.Name, p.Stats.Level))

	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 28)

	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))
	rg.Label(rl.NewRectangle(panelx, panely+25, 200, 25), p.Race)
	rg.Label(rl.NewRectangle(panelx, panely+50, 200, 25), p.Type)

	rg.Panel(rl.NewRectangle(panelx, panely+80, 400, 75), "")
	rg.Panel(rl.NewRectangle(panelx, panely+80, 200, 75), "")

	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))

	rg.Label(rl.NewRectangle(panelx, panely+85, 150, 25), "Life:")
	rg.Label(rl.NewRectangle(panelx, panely+105, 150, 25), "Clicks:")
	rg.Label(rl.NewRectangle(panelx, panely+125, 150, 25), "Aim:")

	rg.Label(rl.NewRectangle(panelx+200, panely+85, 150, 25), "Initiative:")
	rg.Label(rl.NewRectangle(panelx+200, panely+105, 150, 25), "Interrupts:")

	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_LEFT))

	rg.Label(rl.NewRectangle(panelx+225, panely+125, 150, 25), "Money:")

	rg.Label(rl.NewRectangle(panelx+150, panely+85, 25, 25), fmt.Sprint(p.Stats.Life))
	rg.Label(rl.NewRectangle(panelx+150, panely+105, 25, 25), fmt.Sprint(p.Stats.Clicks))
	rg.Label(rl.NewRectangle(panelx+150, panely+125, 25, 25), fmt.Sprint(p.Stats.Aim))

	rg.Label(rl.NewRectangle(panelx+350, panely+85, 25, 25), fmt.Sprint(p.Stats.Initiative))
	rg.Label(rl.NewRectangle(panelx+350, panely+105, 25, 25), fmt.Sprint(p.Stats.Interrupts))
	rg.Label(rl.NewRectangle(panelx+305, panely+125, 95, 25), fmt.Sprint(p.Stats.Money))

	rg.SetStyle(rg.DEFAULT, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_CENTER))
	rg.Panel(rl.NewRectangle(panelx, panely+155, 300, 195), "Attributes")
	rg.Panel(rl.NewRectangle(panelx+300, panely+155, 100, 195), "Mod")

	var (
		w float32 = 250
		h float32 = 25
	)

	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))
	rg.Label(rl.NewRectangle(panelx, panely+180, w, h), "Strength:")
	rg.Label(rl.NewRectangle(panelx, panely+200, w, h), "Dexterity:")
	rg.Label(rl.NewRectangle(panelx, panely+220, w, h), "Fitness:")
	rg.Label(rl.NewRectangle(panelx, panely+240, w, h), "Agility:")
	rg.Label(rl.NewRectangle(panelx, panely+260, w, h), "Wisdom:")
	rg.Label(rl.NewRectangle(panelx, panely+280, w, h), "Intellect:")
	rg.Label(rl.NewRectangle(panelx, panely+300, w, h), "Charisma:")
	rg.Label(rl.NewRectangle(panelx, panely+320, w, h), "Appearance:")

	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_LEFT))
	rg.Label(rl.NewRectangle(panelx+250, panely+180, 25, h), fmt.Sprint(p.Stats.Strength))
	rg.Label(rl.NewRectangle(panelx+250, panely+200, 25, h), fmt.Sprint(p.Stats.Dexterity))
	rg.Label(rl.NewRectangle(panelx+250, panely+220, 25, h), fmt.Sprint(p.Stats.Fitness))
	rg.Label(rl.NewRectangle(panelx+250, panely+240, 25, h), fmt.Sprint(p.Stats.Agility))
	rg.Label(rl.NewRectangle(panelx+250, panely+260, 25, h), fmt.Sprint(p.Stats.Wisdom))
	rg.Label(rl.NewRectangle(panelx+250, panely+280, 25, h), fmt.Sprint(p.Stats.Intellect))
	rg.Label(rl.NewRectangle(panelx+250, panely+300, 25, h), fmt.Sprint(p.Stats.Charisma))
	rg.Label(rl.NewRectangle(panelx+250, panely+320, 25, h), fmt.Sprint(p.Stats.Appearance))

	rg.Line(rl.NewRectangle(panelx+300, panely+180, 100, h), fmt.Sprintf("> %d", p.StatMods.Strength))
	rg.Line(rl.NewRectangle(panelx+300, panely+200, 100, h), fmt.Sprintf("> %d", p.StatMods.Dexterity))
	rg.Line(rl.NewRectangle(panelx+300, panely+220, 100, h), fmt.Sprintf("> %d", p.StatMods.Fitness))
	rg.Line(rl.NewRectangle(panelx+300, panely+240, 100, h), fmt.Sprintf("> %d", p.StatMods.Agility))
	rg.Line(rl.NewRectangle(panelx+300, panely+260, 100, h), fmt.Sprintf("> %d", p.StatMods.Wisdom))
	rg.Line(rl.NewRectangle(panelx+300, panely+280, 100, h), fmt.Sprintf("> %d", p.StatMods.Intellect))
	rg.Line(rl.NewRectangle(panelx+300, panely+300, 100, h), fmt.Sprintf("> %d", p.StatMods.Charisma))
	rg.Line(rl.NewRectangle(panelx+300, panely+320, 100, h), fmt.Sprintf("> %d", p.StatMods.Appearance))

	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
}

func (p *Player) UpdateStatsCopy() PlayerStats {
	return *p.Stats
}

func (p *Player) DistributeStatPoints() {
	var (
		offset float32 = 30
		w      float32 = 300
		h      float32 = 350
		x              = ui.CenterX(w)
		y              = ui.CenterY(h)
	)
	start := StatsCopy
	rg.Panel(rl.NewRectangle(x, y, w, h), fmt.Sprintf("Attribute Points - %d", p.StatPoints))
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 28)
	// Strength
	ui.LabelAlignRight()
	rg.Label(rl.NewRectangle(x, y+offset, 175, 25), "Strength:")
	ui.LabelAlignLeft()
	rg.Label(rl.NewRectangle(x+175, y+offset, 25, 25), fmt.Sprintf("%d", p.Stats.Strength))
	if rg.Button(rl.NewRectangle(x+215, y+offset, 30, 25), "#120#") {
		p.Stats.Strength = p.DecrementStat(p.Stats.Strength, start.Strength)
	}
	if rg.Button(rl.NewRectangle(x+250, y+offset, 30, 25), "#121#") {
		p.Stats.Strength = p.IncrementStat(p.Stats.Strength)
	}

	// Dexterity
	ui.LabelAlignRight()
	rg.Label(rl.NewRectangle(x, y+(offset*2), 175, 25), "Dexterity:")
	ui.LabelAlignLeft()
	rg.Label(rl.NewRectangle(x+175, y+(offset*2), 25, 25), fmt.Sprintf("%d", p.Stats.Dexterity))
	if rg.Button(rl.NewRectangle(x+215, y+(offset*2), 30, 25), "#120#") {
		p.Stats.Dexterity = p.DecrementStat(p.Stats.Dexterity, start.Dexterity)
	}
	if rg.Button(rl.NewRectangle(x+250, y+(offset*2), 30, 25), "#121#") {
		p.Stats.Dexterity = p.IncrementStat(p.Stats.Dexterity)
	}

	// Fitness
	ui.LabelAlignRight()
	rg.Label(rl.NewRectangle(x, y+(offset*3), 175, 25), "Fitness:")
	ui.LabelAlignLeft()
	rg.Label(rl.NewRectangle(x+175, y+(offset*3), 25, 25), fmt.Sprintf("%d", p.Stats.Fitness))
	if rg.Button(rl.NewRectangle(x+215, y+(offset*3), 30, 25), "#120#") {
		p.Stats.Fitness = p.DecrementStat(p.Stats.Fitness, start.Fitness)
	}
	if rg.Button(rl.NewRectangle(x+250, y+(offset*3), 30, 25), "#121#") {
		p.Stats.Fitness = p.IncrementStat(p.Stats.Fitness)
	}

	// Agility
	ui.LabelAlignRight()
	rg.Label(rl.NewRectangle(x, y+(offset*4), 175, 25), "Agility:")
	ui.LabelAlignLeft()
	rg.Label(rl.NewRectangle(x+175, y+(offset*4), 25, 25), fmt.Sprintf("%d", p.Stats.Agility))
	if rg.Button(rl.NewRectangle(x+215, y+(offset*4), 30, 25), "#120#") {
		p.Stats.Agility = p.DecrementStat(p.Stats.Agility, start.Agility)
	}
	if rg.Button(rl.NewRectangle(x+250, y+(offset*4), 30, 25), "#121#") {
		p.Stats.Agility = p.IncrementStat(p.Stats.Agility)
	}

	// Wisdom
	ui.LabelAlignRight()
	rg.Label(rl.NewRectangle(x, y+(offset*5), 175, 25), "Wisdom:")
	ui.LabelAlignLeft()
	rg.Label(rl.NewRectangle(x+175, y+(offset*5), 25, 25), fmt.Sprintf("%d", p.Stats.Wisdom))
	if rg.Button(rl.NewRectangle(x+215, y+(offset*5), 30, 25), "#120#") {
		p.Stats.Wisdom = p.DecrementStat(p.Stats.Wisdom, start.Wisdom)
	}
	if rg.Button(rl.NewRectangle(x+250, y+(offset*5), 30, 25), "#121#") {
		p.Stats.Wisdom = p.IncrementStat(p.Stats.Wisdom)
	}
	// Intellect
	ui.LabelAlignRight()
	rg.Label(rl.NewRectangle(x, y+(offset*6), 175, 25), "Intellect:")
	ui.LabelAlignLeft()
	rg.Label(rl.NewRectangle(x+175, y+(offset*6), 25, 25), fmt.Sprintf("%d", p.Stats.Intellect))
	if rg.Button(rl.NewRectangle(x+215, y+(offset*6), 30, 25), "#120#") {
		p.Stats.Intellect = p.DecrementStat(p.Stats.Intellect, start.Intellect)
	}
	if rg.Button(rl.NewRectangle(x+250, y+(offset*6), 30, 25), "#121#") {
		p.Stats.Intellect = p.IncrementStat(p.Stats.Intellect)
	}
	// Charisma
	ui.LabelAlignRight()
	rg.Label(rl.NewRectangle(x, y+(offset*7), 175, 25), "Charisma:")
	ui.LabelAlignLeft()
	rg.Label(rl.NewRectangle(x+175, y+(offset*7), 25, 25), fmt.Sprintf("%d", p.Stats.Charisma))
	if rg.Button(rl.NewRectangle(x+215, y+(offset*7), 30, 25), "#120#") {
		p.Stats.Charisma = p.DecrementStat(p.Stats.Charisma, start.Charisma)
	}
	if rg.Button(rl.NewRectangle(x+250, y+(offset*7), 30, 25), "#121#") {
		p.Stats.Charisma = p.IncrementStat(p.Stats.Charisma)
	}
	// Appearance
	ui.LabelAlignRight()
	rg.Label(rl.NewRectangle(x, y+(offset*8), 175, 25), "Appearance:")
	ui.LabelAlignLeft()
	rg.Label(rl.NewRectangle(x+175, y+(offset*8), 25, 25), fmt.Sprintf("%d", p.Stats.Appearance))
	if rg.Button(rl.NewRectangle(x+215, y+(offset*8), 30, 25), "#120#") {
		p.Stats.Appearance = p.DecrementStat(p.Stats.Appearance, start.Appearance)
	}
	if rg.Button(rl.NewRectangle(x+250, y+(offset*8), 30, 25), "#121#") {
		p.Stats.Appearance = p.IncrementStat(p.Stats.Appearance)
	}
	if rg.Button(rl.NewRectangle(ui.CenterWithinPanelX(x, w, 200), y+h-65, 200, 50), "Done") {
		statPointView = !statPointView
		StatsCopy = p.UpdateStatsCopy()
	}
}

func (p *Player) DecrementStat(currentStat, startStat int) int {
	if currentStat-1 == startStat {
		p.StatPoints += 1
		return startStat
	} else if currentStat-1 < startStat {
		return currentStat
	} else {
		p.StatPoints += 1
		return currentStat - 1
	}
}

func (p *Player) IncrementStat(currentStat int) int {
	if p.StatPoints == 0 {
		return currentStat
	} else {
		p.StatPoints -= 1
		return currentStat + 1
	}
}

func (p *Player) SetRaceStats(race PlayerStats) {
	p.Stats.Level = race.Level
	p.Stats.Life = race.Life
	p.Stats.Clicks = race.Clicks
	p.Stats.Aim = race.Aim
	p.Stats.Interrupts = race.Interrupts
	p.Stats.Money = race.Money
	p.Stats.Strength = race.Strength
	p.Stats.Dexterity = race.Dexterity
	p.Stats.Fitness = race.Fitness
	p.Stats.Agility = race.Agility
	p.Stats.Wisdom = race.Wisdom
	p.Stats.Intellect = race.Intellect
	p.Stats.Charisma = race.Charisma
	p.Stats.Appearance = race.Appearance
	p.SetStatMods()
	p.StatPoints = 5
}

func (p *Player) SetStatMods() {
	p.StatMods.Strength = CalcStatMod(p.Stats.Strength)
	p.StatMods.Dexterity = CalcStatMod(p.Stats.Dexterity)
	p.StatMods.Fitness = CalcStatMod(p.Stats.Fitness)
	p.StatMods.Agility = CalcStatMod(p.Stats.Agility)
	p.StatMods.Wisdom = CalcStatMod(p.Stats.Wisdom)
	p.StatMods.Intellect = CalcStatMod(p.Stats.Intellect)
	p.StatMods.Charisma = CalcStatMod(p.Stats.Charisma)
	p.StatMods.Appearance = CalcStatMod(p.Stats.Appearance)
	p.Stats.Initiative = max(p.StatMods.Dexterity, p.StatMods.Agility)
}

func CalcStatMod(v int) int {
	switch v {
	case 12:
		return 1
	case 13:
		return 1
	case 14:
		return 2
	case 15:
		return 2
	case 16:
		return 3
	case 17:
		return 3
	case 18:
		return 4
	case 19:
		return 4
	case 20:
		return 5
	case 21:
		return 5
	case 22:
		return 6
	case 23:
		return 6
	case 24:
		return 7
	default:
		return 0
	}
}

func (p *Player) NewPlayerSkills() bool {
	p.DrawPlayerStats()
	p.DrawPlayerSkills()
	return p.DrawAddSkills()
}

func (p *Player) DrawAddSkills() bool {
	skills := *StockSkillSet
	var (
		offset float32 = 30
		w      float32 = 360
		h              = 120 + (float32(len(skills)) * offset)
		y              = ui.CenterY(h)
		x              = ui.CenterX(w)
	)
	rg.Panel(rl.NewRectangle(x, y, w, h), "Generic Skills")

	for i, s := range *StockSkillSet {
		skill := skills[i]

		rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))

		rg.Label(rl.NewRectangle(x, 10+y+offset, 110, 25), fmt.Sprintf("%s ", s.Name))

		rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 28)

		if rg.Button(rl.NewRectangle(x+110, 10+y+offset, 125, 25), "Remove") {
			p.RemoveSkill(s)
			fmt.Println(p.Skills)
		} else if rg.Button(rl.NewRectangle(x+245, 10+y+offset, 75, 25), "Add") {
			p.AddSkill(s)
			fmt.Println(p.Skills)
		}
		rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
		rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_LEFT))

		rg.LabelButton(rl.NewRectangle(x+330, 10+y+offset, 25, 25), "#193#")
		mousePos := rl.GetMousePosition()
		if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(x+330, 10+y+offset, 25, 25)) {
			skill.ViewState = true
			addSkillInfoView = true
			skills[i] = skill
			StockSkillSet = &skills
		} else {
			skill.ViewState = false
			addSkillInfoView = false
			skills[i] = skill
			StockSkillSet = &skills
		}

		if addSkillInfoView {
			if s.ViewState {
				textWidth := float32(rl.MeasureText(s.Info, 28))
				rg.Panel(rl.NewRectangle(x+360, y, textWidth, 200), s.Name)
				rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 28)
				rg.Label(rl.NewRectangle(x+380, 35+y, textWidth, 200), fmt.Sprintf("Info --\n%s\n\nLevel - %d\nBonus From - %s Mod", s.Info, s.Level+s.StatBonus, s.StatKey))
				rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
			}
		}

		offset += 30
	}
	if rg.Button(rl.NewRectangle(ui.CenterWithinPanelX(x, w, 150), (y+h)-65, 150, 50), "Done") {
		fmt.Println(len(*p.Skills))
		if len(*p.Skills) >= 3 {
			return true
		}
	}

	return false
}

func (p *Player) DrawPlayerSkills() {
	var (
		offset float32 = 25
		x      float32 = 25
		y      float32 = 375
		w      float32 = 400
		h              = float32((len(*p.Skills) * 25) + 30)
	)
	rg.Panel(rl.NewRectangle(x, y, w, h), "Skills")
	skills := *p.Skills
	for _, s := range SkillSet {
		ui.SizeText(28)
		skill := skills[s.Name]

		rg.LabelButton(rl.NewRectangle(370+x, y+offset, 25, 25), ">>")

		mousePos := rl.GetMousePosition()
		if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(370+x, y+offset, 25, 25)) {
			skill.ViewState = true
			skillView = true
			skills[s.Name] = skill
			p.Skills = &skills
		} else {
			skill.ViewState = false
			skillView = false
			skills[s.Name] = skill
			p.Skills = &skills
		}

		rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))
		rg.Label(rl.NewRectangle(x, y+offset, 150, 25), s.Name)
		rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_LEFT))
		rg.Label(rl.NewRectangle(150+x, y+offset, 250, 25), fmt.Sprintf(": %d - %s", s.Level+s.StatBonus, s.StatKey))
		if skillView {
			if skill.ViewState {
				w := float32(rl.MeasureText(skill.Info, 32))
				rg.Panel(rl.NewRectangle(400+x, y+offset, w-65, 200), skill.Name)
				rg.Label(rl.NewRectangle(420+x, 40+y+offset, w, 200), fmt.Sprintf("Info --\n%s\n\nLevel - %d\nBonus From - %s Mod", skill.Info, skill.Level+skill.StatBonus, skill.StatKey))
			}
		}

		offset += 25
	}
	ui.ResetText32()
}

func (p *Player) AddSkill(skill Skill) {
	skills := *p.Skills
	_, ok := skills[skill.Name]
	if ok {
		fmt.Println("Skill already in skillset")
		fmt.Println("Skills-------")
		fmt.Println(skills)
	} else {
		fmt.Println("Adding Skill to skillset")
		skill.SetSkillStatBonus(skill.StatKey, p.StatMods)
		skills[skill.Name] = skill
		fmt.Println("Skills-------")
		fmt.Println(skills)
	}

	p.Skills = &skills
	p.UpdateSkillSetCopy()
}

func (p *Player) RemoveSkill(skill Skill) {
	skills := *p.Skills
	_, ok := skills[skill.Name]
	if ok {
		fmt.Println("Removing Skill from skillset")
		delete(skills, skill.Name)
		fmt.Println("Skills-------")
		fmt.Println(skills)
	} else {
		fmt.Println("Skill does not exist in skillset")
		fmt.Println("Skills-------")
		fmt.Println(skills)
	}

	p.Skills = &skills
	p.UpdateSkillSetCopy()
}

func (p *Player) NewSkillMap() {
	s := make(map[string]Skill)
	p.Skills = &s
}

func (p *Player) UpdateSkillSetCopy() {
	var skills []Skill
	for _, s := range *p.Skills {
		skills = append(skills, s)
	}
	SkillSet = skills
}

func (p *Player) NewPlayerMethods() bool {
	p.DrawPlayerStats()
	p.DrawPlayerSkills()
	p.DrawPlayerMethods()

	return p.DrawTypeMethods()
}

func (p *Player) DrawPlayerMethods() {
	var (
		offset float32 = 25
		x      float32 = 25
		y              = float32(375 + ((len(*p.Skills) * 25) + 30))
		w      float32 = 400
		h              = float32((len(*p.Methods) * 25) + 30)
	)
	rg.Panel(rl.NewRectangle(x, y, w, h), "Methods")
	methods := *p.Methods

	for _, m := range MethodSet {
		ui.SizeText(28)
		method := methods[m.Name]
		ui.LabelAlignRight()
		rg.Label(rl.NewRectangle(x, y+offset, 150, 25), m.Name)

		ui.LabelAlignLeft()
		rg.Label(rl.NewRectangle(x+150, y+offset, 200, 25), fmt.Sprintf(": %d - %s", m.StatBonus, m.StatKey))
		rg.LabelButton(rl.NewRectangle(370+x, y+offset, 25, 25), ">>")

		mousePos := rl.GetMousePosition()
		if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(370+x, y+offset, 25, 25)) {
			method.ViewState = true
			methodView = true
			methods[m.Name] = method
			p.Methods = &methods
		} else {
			method.ViewState = false
			methodView = false
			methods[m.Name] = method
			p.Methods = &methods
		}

		if methodView {
			if method.ViewState {
				w := float32(rl.MeasureText(method.Info, 32))
				rg.Panel(rl.NewRectangle(400+x, y+offset, w-65, 200), method.Name)
				rg.Label(rl.NewRectangle(420+x, 40+y+offset, w, 200), fmt.Sprintf("Info --\n%s\n\n%s", method.Info, method.StatKey))
			}
		}

		offset += 25

	}
	ui.ResetText32()
}

func (p *Player) DrawTypeMethods() bool {
	methods := []Method{}
	title := ""
	switch p.Type {
	case "Gunslinger":
		title = "Gunslinger Methods"
		methods = *GunslingerMeth
	case "Naturalist":
		title = "Naturalist Methods"
		methods = *NaturalistMeth
	case "Chemist":
		title = "Chemist Methods"
		methods = *ChemistMeth
	default:
		fmt.Println("DEFAULT")
	}
	var (
		offset float32 = 30
		w      float32 = 415
		h              = 120 + (float32(len(methods)) * offset)
		y              = ui.CenterY(h)
		x              = ui.CenterX(w)
	)

	rg.Panel(rl.NewRectangle(x, y, w, h), title)

	for i, m := range methods {
		method := methods[i]
		rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))
		rg.Label(rl.NewRectangle(x, 10+y+offset, 150, 25), m.Name)

		rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 28)

		if rg.Button(rl.NewRectangle(x+160, 10+y+offset, 125, 25), "Remove") {
			p.RemoveMethod(m)
			fmt.Println(p.Methods)
		} else if rg.Button(rl.NewRectangle(x+295, 10+y+offset, 75, 25), "Add") {
			p.AddMethod(m)
			fmt.Println(p.Methods)
		}

		rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
		rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_LEFT))

		rg.LabelButton(rl.NewRectangle(x+380, 10+y+offset, 25, 25), "#193#")
		mousePos := rl.GetMousePosition()

		if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(x+380, 10+y+offset, 25, 25)) {
			method.ViewState = true
			addMethodInfoView = true
			methods[i] = method
			StockMethodSet = &methods
		} else {
			method.ViewState = false
			addMethodInfoView = false
			methods[i] = method
			StockMethodSet = &methods
		}

		if addMethodInfoView {
			if m.ViewState {
				textWidth := float32(rl.MeasureText(m.Info, 28))
				rg.Panel(rl.NewRectangle(x+415, y, textWidth, 200), m.Name)
				rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 28)
				rg.Label(rl.NewRectangle(x+440, 35+y, textWidth, 200), fmt.Sprintf("Info --\n%s\n\nBonus From - %s Mod", m.Info, m.StatKey))
				rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
			}
		}

		offset += 30
	}

	if rg.Button(rl.NewRectangle(ui.CenterWithinPanelX(x, w, 150), (y+h)-65, 150, 50), "Done") {
		fmt.Println(len(*p.Methods))
		if len(*p.Methods) >= 1 {
			return true
		}
	}
	return false
}

func (p *Player) AddMethod(m Method) {
	fmt.Println("Add Method Called")
	methods := *p.Methods
	_, ok := methods[m.Name]
	if ok {
		fmt.Println("Method already in methodset")
		fmt.Println(*p.Methods)
	} else {
		fmt.Println("Adding Method to methodset")
		methods[m.Name] = m
		p.Methods = &methods
		fmt.Println(*p.Methods)
	}
	p.UpdateMethodSetCopy()
}

func (p *Player) RemoveMethod(m Method) {
	fmt.Println("Remove Method Called")
	methods := *p.Methods
	_, ok := methods[m.Name]
	if ok {
		fmt.Println("Removing Method from methodset")
		delete(methods, m.Name)
		p.Methods = &methods
		fmt.Println(*p.Methods)
	} else {
		fmt.Println("Method does not exist in methodset")
		fmt.Println(*p.Methods)
	}
	p.UpdateMethodSetCopy()
}

func (p *Player) NewMethodMap() {
	m := make(map[string]Method)
	p.Methods = &m
}

func (p *Player) UpdateMethodSetCopy() {
	var methods []Method
	for _, m := range *p.Methods {
		methods = append(methods, m)
	}
	MethodSet = methods
}

func (p *Player) NewPlayerInventory() bool {
	p.DrawPlayerStats()
	p.DrawPlayerSkills()
	p.DrawPlayerMethods()
	p.DrawPlayerInventory()
	return false
}

func (p *Player) DrawPlayerInventory() {
	var (
		offset float32 = 25
		w      float32 = 400
		h              = float32(rl.GetScreenHeight() - 50)
		x              = float32(rl.GetScreenWidth() - 425)
		y      float32 = 25
	)

	rg.Panel(rl.NewRectangle(x, y, w, h), "Inventory")
	rg.Label(rl.NewRectangle(x+offset, y+offset, 200, 25), "Test Item")
}

func (p *Player) CreateInventory() {
	var armor []gameobj.Armor
	var weapon []gameobj.Weapon
	var item []gameobj.Item
	p.Inventory = &PlayerInventory{
		Armor:  &armor,
		Weapon: &weapon,
		Item:   &item,
	}
}

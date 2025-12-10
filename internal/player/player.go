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
	statView         = true
	statPointView    = true
	skillView        = false
	addSkillInfoView = false
	StatsCopy        = PlayerStats{}
	SkillSet         = []Skill{}

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
	Methods    *[]Method
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
	p.NewMethodSlice()
	p.CreateInventory()
	p.UpdateStatsCopy(*p.Stats)
	p.UpdateSkillSetCopy()
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
	rg.TextBox(rec, &p.Name, 25, true)
	return false
}

func (p *Player) NewPlayerRace() bool {
	p.DrawPlayerStats()
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
		return true
	}
	return false
}

func (p *Player) NewPlayerType() bool {
	p.DrawPlayerStats()
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
		StatsCopy = p.UpdateStatsCopy(StatsCopy)
		return true
	}
	return false
}

func (p *Player) NewPlayerStats() bool {
	if rl.IsKeyPressed(rl.KeyC) {
		statView = !statView
	}
	if statView {
		p.DrawPlayerStats()
	}

	if statPointView {
		p.DistributeStatPoints()
		p.SetStatMods()
	} else {
		p.SetStatMods()
		return true
	}

	roll := false

	if rg.Button(rl.NewRectangle(ui.CenterX(200), ui.BottomAlign(50)-60, 200, 50), "Roll Stats") {
		roll = !roll
	}

	if roll {
		p.Stats.RollStats()
		StatsCopy = p.UpdateStatsCopy(StatsCopy)
	}

	if rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50), 200, 50), "Confirm Stats") {
		if p.StatPoints == 0 {
			statPointView = !statPointView
		} else {
			rg.Label(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+60, 200, 50), "Attribute")
		}
	}

	return false
}

func (p *Player) DrawPlayerStats() {
	// rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_CENTER))
	bounds := rl.NewRectangle(0, 0, 400, float32(rl.GetScreenHeight()))
	rg.Panel(bounds, fmt.Sprintf("Character -- Level %d", p.Stats.Level))

	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 28)
	rg.Label(rl.NewRectangle(25, 25, 200, 50), p.Name)
	rg.Label(rl.NewRectangle(25, 50, 200, 50), p.Race)
	rg.Label(rl.NewRectangle(25, 75, 200, 50), p.Type)

	var (
		w float32 = 155
		h float32 = 25
	)
	rg.Panel(rl.NewRectangle(0, 125, 400, 75), "")
	rg.Panel(rl.NewRectangle(0, 125, 200, 75), "")

	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))
	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 26)

	rg.Label(rl.NewRectangle(0, 130, w, h), "Life:")
	rg.Label(rl.NewRectangle(0, 150, w, h), "Clicks:")
	rg.Label(rl.NewRectangle(0, 170, w, h), "Aim:")

	rg.Label(rl.NewRectangle(200, 130, w, h), "Initiative:")
	rg.Label(rl.NewRectangle(200, 150, w, h), "Interrupts:")

	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_LEFT))

	rg.Label(rl.NewRectangle(225, 170, w, h), "Money:")

	rg.Label(rl.NewRectangle(155, 130, 25, h), fmt.Sprint(p.Stats.Life))
	rg.Label(rl.NewRectangle(155, 150, 25, h), fmt.Sprint(p.Stats.Clicks))
	rg.Label(rl.NewRectangle(155, 170, 25, h), fmt.Sprint(p.Stats.Aim))

	rg.Label(rl.NewRectangle(355, 130, 25, h), fmt.Sprint(p.Stats.Initiative))
	rg.Label(rl.NewRectangle(355, 150, 25, h), fmt.Sprint(p.Stats.Interrupts))
	rg.Label(rl.NewRectangle(300, 170, 95, h), fmt.Sprint(p.Stats.Money))

	rg.SetStyle(rg.DEFAULT, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_CENTER))
	rg.Panel(rl.NewRectangle(0, 200, 200, 205), "Attribute")
	rg.Panel(rl.NewRectangle(200, 200, 65, 205), "Mod")

	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))
	rg.Label(rl.NewRectangle(0, 225, w, h), "Strength:")
	rg.Label(rl.NewRectangle(0, 245, w, h), "Dexterity:")
	rg.Label(rl.NewRectangle(0, 265, w, h), "Fitness:")
	rg.Label(rl.NewRectangle(0, 285, w, h), "Agility:")
	rg.Label(rl.NewRectangle(0, 305, w, h), "Wisdom:")
	rg.Label(rl.NewRectangle(0, 325, w, h), "Intellect:")
	rg.Label(rl.NewRectangle(0, 345, w, h), "Charisma:")
	rg.Label(rl.NewRectangle(0, 365, w, h), "Appearance:")

	rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_LEFT))
	rg.Label(rl.NewRectangle(155, 225, 25, h), fmt.Sprint(p.Stats.Strength))
	rg.Label(rl.NewRectangle(155, 245, 25, h), fmt.Sprint(p.Stats.Dexterity))
	rg.Label(rl.NewRectangle(155, 265, 25, h), fmt.Sprint(p.Stats.Fitness))
	rg.Label(rl.NewRectangle(155, 285, 25, h), fmt.Sprint(p.Stats.Agility))
	rg.Label(rl.NewRectangle(155, 305, 25, h), fmt.Sprint(p.Stats.Wisdom))
	rg.Label(rl.NewRectangle(155, 325, 25, h), fmt.Sprint(p.Stats.Intellect))
	rg.Label(rl.NewRectangle(155, 345, 25, h), fmt.Sprint(p.Stats.Charisma))
	rg.Label(rl.NewRectangle(155, 365, 25, h), fmt.Sprint(p.Stats.Appearance))

	rg.Line(rl.NewRectangle(200, 238, 65, 1), fmt.Sprintf("> %d", p.StatMods.Strength))
	rg.Line(rl.NewRectangle(200, 258, 65, 1), fmt.Sprintf("> %d", p.StatMods.Dexterity))
	rg.Line(rl.NewRectangle(200, 278, 65, 1), fmt.Sprintf("> %d", p.StatMods.Fitness))
	rg.Line(rl.NewRectangle(200, 298, 65, 1), fmt.Sprintf("> %d", p.StatMods.Agility))
	rg.Line(rl.NewRectangle(200, 318, 65, 1), fmt.Sprintf("> %d", p.StatMods.Wisdom))
	rg.Line(rl.NewRectangle(200, 338, 65, 1), fmt.Sprintf("> %d", p.StatMods.Intellect))
	rg.Line(rl.NewRectangle(200, 358, 65, 1), fmt.Sprintf("> %d", p.StatMods.Charisma))
	rg.Line(rl.NewRectangle(200, 378, 65, 1), fmt.Sprintf("> %d", p.StatMods.Appearance))

	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
}

func (p *Player) UpdateStatsCopy(stats PlayerStats) PlayerStats {
	return *p.Stats
}

func (p *Player) DistributeStatPoints() {
	start := StatsCopy
	rg.Panel(rl.NewRectangle(265, 200, 60, 205), fmt.Sprint(p.StatPoints))
	// Strength
	if rg.LabelButton(rl.NewRectangle(275, 225, 10, 25), "-") {
		p.Stats.Strength = p.DecrementStat(p.Stats.Strength, start.Strength)
	}
	if rg.LabelButton(rl.NewRectangle(300, 225, 10, 25), "+") {
		p.Stats.Strength = p.IncrementStat(p.Stats.Strength)
	}

	// Dexterity
	if rg.LabelButton(rl.NewRectangle(275, 245, 10, 25), "-") {
		p.Stats.Dexterity = p.DecrementStat(p.Stats.Dexterity, start.Dexterity)
	}
	if rg.LabelButton(rl.NewRectangle(300, 245, 10, 25), "+") {
		p.Stats.Dexterity = p.IncrementStat(p.Stats.Dexterity)
	}

	// Fitness
	if rg.LabelButton(rl.NewRectangle(275, 265, 10, 25), "-") {
		p.Stats.Fitness = p.DecrementStat(p.Stats.Fitness, start.Fitness)
	}
	if rg.LabelButton(rl.NewRectangle(300, 265, 10, 25), "+") {
		p.Stats.Fitness = p.IncrementStat(p.Stats.Fitness)
	}

	// Agility
	if rg.LabelButton(rl.NewRectangle(275, 285, 10, 25), "-") {
		p.Stats.Agility = p.DecrementStat(p.Stats.Agility, start.Agility)
	}
	if rg.LabelButton(rl.NewRectangle(300, 285, 10, 25), "+") {
		p.Stats.Agility = p.IncrementStat(p.Stats.Agility)
	}

	// Wisdom
	if rg.LabelButton(rl.NewRectangle(275, 305, 10, 25), "-") {
		p.Stats.Wisdom = p.DecrementStat(p.Stats.Wisdom, start.Wisdom)
	}
	if rg.LabelButton(rl.NewRectangle(300, 305, 10, 25), "+") {
		p.Stats.Wisdom = p.IncrementStat(p.Stats.Wisdom)
	}
	// Intellect
	if rg.LabelButton(rl.NewRectangle(275, 325, 10, 25), "-") {
		p.Stats.Intellect = p.DecrementStat(p.Stats.Intellect, start.Intellect)
	}
	if rg.LabelButton(rl.NewRectangle(300, 325, 10, 25), "+") {
		p.Stats.Intellect = p.IncrementStat(p.Stats.Intellect)
	}
	// Charisma
	if rg.LabelButton(rl.NewRectangle(275, 345, 10, 25), "-") {
		p.Stats.Charisma = p.DecrementStat(p.Stats.Charisma, start.Charisma)
	}
	if rg.LabelButton(rl.NewRectangle(300, 345, 10, 25), "+") {
		p.Stats.Charisma = p.IncrementStat(p.Stats.Charisma)
	}
	// Appearance
	if rg.LabelButton(rl.NewRectangle(275, 365, 10, 25), "-") {
		p.Stats.Appearance = p.DecrementStat(p.Stats.Appearance, start.Appearance)
	}
	if rg.LabelButton(rl.NewRectangle(300, 365, 10, 25), "+") {
		p.Stats.Appearance = p.IncrementStat(p.Stats.Appearance)
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

func (p *Player) NewPlayerInventory() bool {
	return false
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

func (p *Player) NewPlayerSkills() bool {
	if p.DrawAddSkills() {
		return true
	}
	p.DrawPlayerStats()
	p.DrawSkillSet()
	return false
}

func (p *Player) DrawAddSkills() bool {
	rg.Panel(rl.NewRectangle(400, 405, 550, float32(rl.GetScreenHeight())), "Skills")
	skills := *StockSkillSet
	var offset float32 = 25
	for i, s := range *StockSkillSet {
		skill := skills[i]
		rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))
		rg.Label(rl.NewRectangle(400, 405+offset, 100, 25), s.Name)

		rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_LEFT))
		rg.Label(rl.NewRectangle(500, 405+offset, 200, 25), fmt.Sprintf(" - %s", s.StatKey))

		if rg.Button(rl.NewRectangle(700, 405+offset, 125, 25), "Remove") {
			p.RemoveSkill(s)
			fmt.Println(p.Skills)
		} else if rg.Button(rl.NewRectangle(835, 405+offset, 75, 25), "Add") {
			p.AddSkill(s)
			fmt.Println(p.Skills)
		}

		rg.LabelButton(rl.NewRectangle(920, 405+offset, 25, 25), ">>")
		mousePos := rl.GetMousePosition()
		if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(920, 405+offset, 25, 25)) {
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
				w := float32(rl.MeasureText(s.Info, 32))
				rg.Panel(rl.NewRectangle(950, 405+offset, w-65, 200), s.Name)
				rg.Label(rl.NewRectangle(970, 445+offset, w, 200), fmt.Sprintf("Info --\n%s\n\nLevel - %d\nBonus From - %s Mod", s.Info, s.Level+s.StatBonus, s.StatKey))
			}
		}

		offset += 25
	}
	if rg.Button(rl.NewRectangle(ui.CenterWithinPanelX(400, 550, 150), ui.BottomAlign(50), 150, 50), "Done") {
		fmt.Println(len(*p.Skills))
		if len(*p.Skills) >= 3 {
			return true
		}
	}

	return false
}

func (p *Player) DrawSkillSet() {
	var offset float32 = 25
	rg.Panel(rl.NewRectangle(0, 405, 400, 230), "Skills")
	skills := *p.Skills

	for _, s := range SkillSet {

		skill := skills[s.Name]

		rg.LabelButton(rl.NewRectangle(370, 405+offset, 25, 25), ">>")

		mousePos := rl.GetMousePosition()
		if rl.CheckCollisionPointRec(mousePos, rl.NewRectangle(370, 405+offset, 25, 25)) {
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
		rg.Label(rl.NewRectangle(0, 405+offset, 150, 25), s.Name)
		rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_LEFT))
		rg.Label(rl.NewRectangle(150, 405+offset, 250, 25), fmt.Sprintf(": %d - %s", s.Level+s.StatBonus, s.StatKey))
		if skillView {
			if skill.ViewState {
				w := float32(rl.MeasureText(skill.Info, 32))
				rg.Panel(rl.NewRectangle(400, 405+offset, w-65, 200), skill.Name)
				rg.Label(rl.NewRectangle(420, 445+offset, w, 200), fmt.Sprintf("Info --\n%s\n\nLevel - %d\nBonus From - %s Mod", skill.Info, skill.Level+skill.StatBonus, skill.StatKey))
			}
		}

		offset += 25
	}
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
	p.DrawSkillSet()
	p.DrawPlayerMethods()
	return false
}

func (p *Player) DrawPlayerMethods() {
	var offset float32 = 25
	rg.Panel(rl.NewRectangle(0, 635, 400, 250), "Methods")

	for _, m := range TestMethods {
		rg.SetStyle(rg.LABEL, rg.TEXT_ALIGNMENT, int64(rg.TEXT_ALIGN_RIGHT))
		rg.Label(rl.NewRectangle(0, 635+offset, 400, 25), m.Name)
		offset += 25

	}
}

func (p *Player) NewMethodSlice() {
	var m []Method
	p.Methods = &m
}

func (pS *PlayerStats) RollStats() {
	pS.Level = 1
	pS.Life = RollDiceSim(4, 6)
	pS.Clicks = 6
	pS.Initiative = 0
	pS.Aim = 1
	pS.Interrupts = 0
	pS.Money = RollDiceSim(4, 6)
	pS.Strength = RollDiceSim(4, 6)
	pS.Dexterity = RollDiceSim(4, 6)
	pS.Fitness = RollDiceSim(4, 6)
	pS.Agility = RollDiceSim(4, 6)
	pS.Wisdom = RollDiceSim(4, 6)
	pS.Intellect = RollDiceSim(4, 6)
	pS.Charisma = RollDiceSim(4, 6)
	pS.Appearance = RollDiceSim(4, 6)
}

func RollDiceSim(dice, sides int) int {
	lowest := sides + 1
	var total int
	for range dice {
		num := int(rl.GetRandomValue(1, int32(sides)))
		total += num
		if num < lowest {
			lowest = num
		}
	}

	return total - lowest
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

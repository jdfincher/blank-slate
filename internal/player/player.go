// Package player implements the structs for the player
package player

import (
	"fmt"
	"strconv"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/jdfincher/blank-slate/internal/gameobj"
	"github.com/jdfincher/blank-slate/internal/ui"
)

type Player struct {
	Name      string
	Race      string
	Type      string
	Stats     *PlayerStats
	StatMods  *PlayerStatMods
	Inventory *PlayerInventory
	Skills    *[]Skill
	Methods   *[]Method
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

func (p *Player) ResetPlayer() {
	p.Name = ""
	p.Race = ""
	p.Type = ""
	p.Stats = new(PlayerStats)
	p.StatMods = new(PlayerStatMods)
}

func (p *Player) NewPlayerName() bool {
	rl.DrawText(p.Name, ui.CenterTextX(p.Name, 32), int32(ui.CenterY(50))-60, 32, rl.Red)
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
	if human := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)-60, 200, 50), "Human"); human {
		p.Race = "Human"
		fmt.Printf("Player Race is %s\n", p.Race)
		return true
	} else if mutant := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50), 200, 50), "Mutant"); mutant {
		p.Race = "Mutant"
		fmt.Printf("Player Race is %s\n", p.Race)
		return true
	} else if cyborg := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+60, 200, 50), "Cyborg"); cyborg {
		p.Race = "Cyborg"
		fmt.Printf("Player Race is %s\n", p.Race)
		return true
	}
	return false
}

func (p *Player) NewPlayerType() bool {
	if gun := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)-60, 200, 50), "Gunslinger"); gun {
		p.Type = "Gunslinger"
		fmt.Printf("Player Type is %s\n", p.Type)
		return true
	} else if nat := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50), 200, 50), "Naturalist"); nat {
		p.Type = "Naturalist"
		fmt.Printf("Player Type is %s\n", p.Type)
		return true
	} else if che := rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50)+60, 200, 50), "Chemist"); che {
		p.Type = "Chemist"
		fmt.Printf("Player Type is %s\n", p.Type)
		return true
	}
	return false
}

var (
	textLife = "0"
	editLife = false

	textLevel = "0"
	editLevel = false

	textClicks = "0"
	editClicks = false

	statView = true
)

func (p *Player) NewPlayerStats() bool {
	if rl.IsKeyPressed(rl.KeyC) {
		statView = !statView
	}
	if statView {
		p.DrawPlayerStats()
	}

	roll := false

	if rg.Button(rl.NewRectangle(ui.CenterX(200), ui.CenterY(50), 200, 50), "Roll Stats") {
		roll = !roll
	}

	if roll {
		p.Stats.RollStats()
	}
	p.SetStatMods()

	if rg.TextBox(rl.NewRectangle(350, 500, 50, 50), &textLife, 32, editLife) {
		editLife = !editLife
		life, err := strconv.Atoi(textLife)
		if err != nil {
			fmt.Println("Error converting life to int")
		}
		p.Stats.Strength = life
		fmt.Printf("Player Life is -> %d\n", p.Stats.Life)

	} else if rg.TextBox(rl.NewRectangle(350, 560, 50, 50), &textLevel, 32, editLevel) {
		editLevel = !editLevel
		level, err := strconv.Atoi(textLevel)
		if err != nil {
			fmt.Println("Error converting level to int")
		}
		p.Stats.Money = level
		fmt.Printf("Player Level is -> %d\n", p.Stats.Money)
	}

	return false
}

func NewPlayerInventory() PlayerInventory {
	return PlayerInventory{}
}

func NewPlayerSkills() []Skill {
	return make([]Skill, 1)
}

func NewPlayerMethods() []Method {
	return make([]Method, 1)
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

	rg.Line(rl.NewRectangle(200, 238, 65, 1), fmt.Sprintf("+ %d", p.StatMods.Strength))
	rg.Line(rl.NewRectangle(200, 258, 65, 1), fmt.Sprintf("+ %d", p.StatMods.Dexterity))
	rg.Line(rl.NewRectangle(200, 278, 65, 1), fmt.Sprintf("+ %d", p.StatMods.Fitness))
	rg.Line(rl.NewRectangle(200, 298, 65, 1), fmt.Sprintf("+ %d", p.StatMods.Agility))
	rg.Line(rl.NewRectangle(200, 318, 65, 1), fmt.Sprintf("+ %d", p.StatMods.Wisdom))
	rg.Line(rl.NewRectangle(200, 338, 65, 1), fmt.Sprintf("+ %d", p.StatMods.Intellect))
	rg.Line(rl.NewRectangle(200, 358, 65, 1), fmt.Sprintf("+ %d", p.StatMods.Charisma))
	rg.Line(rl.NewRectangle(200, 378, 65, 1), fmt.Sprintf("+ %d", p.StatMods.Appearance))

	rg.SetStyle(rg.DEFAULT, rg.TEXT_SIZE, 32)
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

package gameobj

const (
	GUN = iota
	MELEE
	THROWN
)

var WeaponTypes = map[int]string{
	GUN:    "Projectile",
	MELEE:  "Melee",
	THROWN: "Thrown",
}

type Weapon struct {
	Name     string // Name for Ui display
	Price    int
	DmgMin   int
	DmgMax   int
	Range    int     // Distance in grid squares *  for ft
	Reload   int     // Clicks needed for full reload
	Capacity int     // Amount of attacks before reload needed
	Critical int     // Min roll needed for critical
	CritMin  float32 // Min Crit Multiplier
	CritMax  float32 // Max Crit Multiplier
	AimBonus int
	Caliber  int // Used to keep track of ammo in inventory 0 for Melee
	Count    int // Count of actual weapon
	Type     int // AREA, GUN, MELEE
	AOE      int // x and y dimension of effected grid squares
}

var (
	GunslingerStartWeapon = []Weapon{Revolver22}

	Revolver22 = Weapon{
		Name:     "Revolver .22",
		Price:    75,
		DmgMin:   1,
		DmgMax:   6,
		Range:    80,
		Reload:   6,
		Capacity: 6,
		Critical: 20,
		CritMin:  1.5,
		CritMax:  2.0,
		AimBonus: 1,
		Caliber:  22,
		Count:    1,
		Type:     GUN,
		AOE:      1,
	}
	Revolver38 = Weapon{
		Name:     "Revolver .38",
		Price:    125,
		DmgMin:   2,
		DmgMax:   8,
		Range:    100,
		Reload:   6,
		Capacity: 6,
		Critical: 20,
		CritMin:  1.5,
		CritMax:  2.0,
		AimBonus: 1,
		Caliber:  38,
		Count:    1,
		Type:     GUN,
		AOE:      1,
	}

	NaturalistStartWeapon = []Weapon{SpikeClub, Crowbar}

	SpikeClub = Weapon{
		Name:     "Spike Club",
		Price:    75,
		DmgMin:   2,
		DmgMax:   8,
		Range:    5,
		Reload:   1,
		Capacity: 1,
		Critical: 19,
		CritMin:  2.0,
		CritMax:  2.5,
		AimBonus: 2,
		Caliber:  0,
		Count:    1,
		Type:     MELEE,
		AOE:      1,
	}
	Crowbar = Weapon{
		Name:     "Crowbar",
		Price:    50,
		DmgMin:   1,
		DmgMax:   4,
		Range:    5,
		Reload:   0,
		Capacity: 0,
		Critical: 19,
		CritMin:  1.5,
		CritMax:  1.75,
		AimBonus: 2,
		Caliber:  0,
		Count:    1,
		Type:     MELEE,
		AOE:      1,
	}

	ChemistStartWeapon = []Weapon{AcidFlask, HydroFlask}

	AcidFlask = Weapon{
		Name:     "Acid Flask",
		Price:    100,
		DmgMin:   4,
		DmgMax:   10,
		Range:    20,
		Reload:   3,
		Capacity: 1,
		Critical: 24,
		CritMin:  2.0,
		CritMax:  3.0,
		AimBonus: 4,
		Caliber:  0,
		Count:    3, // count should decrease when flask is used
		Type:     THROWN,
		AOE:      3,
	}

	HydroFlask = Weapon{
		Name:     "Hydro Flask",
		Price:    100,
		DmgMin:   2,
		DmgMax:   12,
		Range:    20,
		Reload:   3,
		Capacity: 1,
		Critical: 24,
		CritMin:  1.75,
		CritMax:  2.25,
		AimBonus: 4,
		Caliber:  0,
		Count:    3,
		Type:     THROWN,
		AOE:      3,
	}
)

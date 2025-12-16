package gameobj

var (
	Revolver22 = Weapon{
		Name:     "Revolver .22",
		Price:    75,
		DmgMin:   1,
		DmgMax:   6,
		Range:    80,
		Reload:   6,
		Capacity: 6,
		Critical: 20,
		AimBonus: 1,
		Caliber:  22,
		Count:    1,
		Type:     "Gun",
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
		AimBonus: 1,
		Caliber:  38,
		Count:    1,
		Type:     "Gun",
	}
)

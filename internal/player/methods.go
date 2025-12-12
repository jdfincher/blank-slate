package player

var (
	DoubleUp = Method{
		Name:      "Double Up",
		Info:      "Wield two single handed weapons at once.",
		StatBonus: -1,
		StatKey:   "Aim",
		Active:    false,
		ViewState: false,
	}
	OakWill = Method{
		Name:      "Oak Will",
		Info:      "Through sheer will-power alone stop bleeding.",
		StatBonus: 1,
		StatKey:   "Strength",
		Active:    false,
		ViewState: false,
	}
	Concoct = Method{
		Name:      "Concoct",
		Info:      "Cook up 3 flasks, once per day, sell the scraps.",
		StatBonus: 10,
		StatKey:   "Money",
		Active:    false,
		ViewState: false,
	}

	GunslingerMeth = &[]Method{DoubleUp, Concoct, OakWill}
	NaturalistMeth = &[]Method{OakWill, DoubleUp, Concoct}
	ChemistMeth    = &[]Method{Concoct, OakWill, DoubleUp}
)

type Method struct {
	Name      string
	Info      string
	StatKey   string
	StatBonus int
	Active    bool
	ViewState bool
}

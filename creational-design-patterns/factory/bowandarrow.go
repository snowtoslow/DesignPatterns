package factory

type BowAndArrow struct {
	Weapon
}

func NewBowAndArrow() IWeapon {
	return &BowAndArrow{
		Weapon{
			Name:  "Bow And Arrow",
			Power: 1,
		},
	}
}

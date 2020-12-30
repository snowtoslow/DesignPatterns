package factory

type BowAndArrow struct {
	Weapon
}

func (s BowAndArrow) CreateWeapon() (Weapon IWeapon) {
	return NewBowAndArrow()
}

func NewBowAndArrow() IWeapon {
	return &BowAndArrow{
		Weapon{
			Name:  "Bow And Arrow",
			Power: 1,
		},
	}
}

package factory

type Spear struct {
	Weapon
}

func NewSpear() IWeapon {
	return &Spear{
		Weapon{
			Name:  "Spear",
			Power: 7,
		},
	}
}

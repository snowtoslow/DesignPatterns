package factory

type Spear struct {
	Weapon
}

func (s Spear) CreateWeapon() (Weapon IWeapon) {
	return NewSpear()
}

func NewSpear() IWeapon {
	return &Spear{
		Weapon{
			Name:  "Spear",
			Power: 7,
		},
	}
}

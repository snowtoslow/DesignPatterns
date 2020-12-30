package factory

type Sword struct {
	Weapon
}

func (s Sword) CreateWeapon() (Weapon IWeapon) {
	return NewSword()
}

func NewSword() IWeapon {
	return &Sword{
		Weapon{
			Name:  "Sword",
			Power: 9,
		},
	}
}

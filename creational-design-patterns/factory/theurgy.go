package factory

type Theurgy struct {
	Weapon
}

func (s Theurgy) CreateWeapon() (Weapon IWeapon) {
	return NewTheurgy()
}

func NewTheurgy() IWeapon {
	return &Theurgy{
		Weapon{
			Name:  "Theurgy",
			Power: 12,
		},
	}
}

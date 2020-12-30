package factory

type Dagger struct {
	Weapon
}

func (d Dagger) CreateWeapon() (Weapon IWeapon) {
	return NewDagger()
}

func NewDagger() IWeapon {
	return &Dagger{
		Weapon{
			Name:  "Dagger",
			Power: 6,
		},
	}
}

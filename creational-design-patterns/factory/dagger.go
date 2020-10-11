package factory

type Dagger struct {
	Weapon
}

func NewDagger() IWeapon {
	return &Dagger{
		Weapon{
			Name:  "Dagger",
			Power: 6,
		},
	}
}

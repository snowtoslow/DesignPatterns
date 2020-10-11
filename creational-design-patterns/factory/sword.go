package factory

type Sword struct {
	Weapon
}

func NewSword() IWeapon {
	return &Sword{
		Weapon{
			Name:  "Sword",
			Power: 9,
		},
	}
}

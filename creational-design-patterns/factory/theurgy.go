package factory

type Theurgy struct {
	Weapon
}

func NewTheurgy() IWeapon {
	return &Theurgy{
		Weapon{
			Name:  "Theurgy",
			Power: 12,
		},
	}
}

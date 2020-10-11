package factory

type IWeapon interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}

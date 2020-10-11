package factory

type Weapon struct {
	Name  string
	Power int
}

func (w *Weapon) SetName(name string) {
	w.Name = name
}

func (w *Weapon) GetName() string {
	return w.Name
}

func (w *Weapon) SetPower(power int) {
	w.Power = power
}

func (w *Weapon) GetPower() int {
	return w.Power
}

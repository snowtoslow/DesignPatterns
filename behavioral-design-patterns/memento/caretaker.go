package memento

type CareTaker struct {
	MementoArray []*Memento
}

func (c *CareTaker) AddMemento(m *Memento) {
	c.MementoArray = append(c.MementoArray, m)
}

func (c *CareTaker) GetMenento(index int) *Memento {
	return c.MementoArray[index]
}

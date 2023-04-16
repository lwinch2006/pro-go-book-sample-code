package models

type Person struct {
	Name string
}

func NewPerson(name string) *Person {
	return &Person{name}
}

func (p *Person) GetName() string {
	return p.Name
}

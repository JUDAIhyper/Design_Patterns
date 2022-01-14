package model

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			Name:  "musket gun",
			Power: 1,
		},
	}
}

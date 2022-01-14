package model

type Adidas struct{}

func (a *Adidas) MakeShoe() IShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 42,
		},
	}
}

func (a *Adidas) MakeShirt() IShirt {
	return &adidasShirt{
		shirt: shirt{
			logo: "adidas",
			size: 14,
		},
	}
}

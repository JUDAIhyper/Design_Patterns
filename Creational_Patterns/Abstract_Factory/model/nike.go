package model

type Nike struct{}

func (n *Nike) MakeShoe() IShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 41,
		},
	}
}

func (n *Nike) MakeShirt() IShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 13,
		},
	}
}

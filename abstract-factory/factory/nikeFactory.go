package factory

type Nike struct {
}

type NikeShoe struct {
	Shoe
}

type NikeShirt struct {
	Shirt
}

func (nike *Nike) CreateShoe() IShoe {
	return &AdidasShoe{
		Shoe{
			logo: "nike",
			size: 41,
		},
	}
}

func (nike *Nike) CreateShirt() IShirt {
	return &AdidasShirt{
		Shirt{
			logo: "nike",
			size: "2xl",
		},
	}

}

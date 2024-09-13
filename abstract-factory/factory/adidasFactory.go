package factory

type Adidas struct {
}

type AdidasShoe struct {
	Shoe
}

type AdidasShirt struct {
	Shirt
}

func (adidas *Adidas) CreateShoe() IShoe {
	return &AdidasShoe{
		Shoe{
			logo: "adidas",
			size: 42,
		},
	}
}

func (adidas *Adidas) CreateShirt() IShirt {
	return &AdidasShirt{
		Shirt{
			logo: "adidas",
			size: "xl",
		},
	}

}

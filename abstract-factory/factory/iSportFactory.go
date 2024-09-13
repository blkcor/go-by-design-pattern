package factory

import "errors"

type ISportFactory interface {
	CreateShoe() IShoe
	CreateShirt() IShirt
}

func GetSportFactory(brand string) (ISportFactory, error) {
	if brand == "nike" {
		return &Nike{}, nil
	}
	if brand == "adidas" {
		return &Adidas{}, nil
	}
	return nil, errors.New("Factory Not Found")
}

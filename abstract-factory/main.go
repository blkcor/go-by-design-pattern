package main

import (
	"fmt"
	"github.com/blkcor/go-design-pattern/abstract-factory/factory"
)

func main() {
	f, err := factory.GetSportFactory("nike")
	if err != nil {
		panic(err)
	}
	shoe := f.CreateShoe()
	shirt := f.CreateShirt()
	fmt.Printf("shoe :::::: logo is %s, and the size is %d\n", shoe.GetLogo(), shoe.GetSize())
	fmt.Printf("shirt :::::: logo is %s, and the size is %s\n", shirt.GetLogo(), shirt.GetSize())

	f, err = factory.GetSportFactory("adidas")
	if err != nil {
		panic(err)
	}
	shoe = f.CreateShoe()
	shirt = f.CreateShirt()
	fmt.Printf("shoe :::::: logo is %s, and the size is %d\n", shoe.GetLogo(), shoe.GetSize())
	fmt.Printf("shirt :::::: logo is %s, and the size is %s\n", shirt.GetLogo(), shirt.GetSize())
}

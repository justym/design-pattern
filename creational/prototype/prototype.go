//This example shows Stock Keeping Unit(SKU) in case of cutomized shirts shop

package prototype

import (
	"errors"
	"fmt"
)

type ShirtsCloner interface {
	GetClone(s int) (ItemInfoGetter, error)
}

type ItemInfoGetter interface {
	GetInfo() string
}

const (
	White = iota
	Black
	Blue
)

type ShirtColor byte

type Shirt struct {
	Price float32
	SKU   string
	Color ShirtColor
}

func (s *Shirt) GetInfo() string {
	return fmt.Sprintf("SKU %s, Color %d, Price %f", s.SKU, s.Color, s.Price)
}

func GetShirtsCloner() ShirtsCloner {
	return &ShirtsCache{}
}

var whitePrototype *Shirt = &Shirt{
	Price: 15.00,
	SKU:   "empty",
	Color: White,
}

var blackPrototype *Shirt = &Shirt{
	Price: 12.00,
	SKU:   "empty",
	Color: Black,
}

var bluePrototype *Shirt = &Shirt{
	Price: 17.00,
	SKU:   "empty",
	Color: Blue,
}

func (i *Shirt) GetPrice() float32 {
	return i.Price
}

type ShirtsCache struct{}

func (c *ShirtsCache) GetClone(s int) (ItemInfoGetter, error) {
	switch s {
	case White:
		newItem := *whitePrototype
		return &newItem, nil
	case Black:
		newItem := *blackPrototype
		return &newItem, nil
	case Blue:
		newItem := *blackPrototype
		return &newItem, nil
	}
	return nil, errors.New("This shirt model has not found")
}

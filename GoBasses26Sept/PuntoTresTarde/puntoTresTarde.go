package main

import (
	"fmt"
)

const (
	small  = "small"
	medium = "medium"
	big    = "big"
)

type TypeItem struct {
	Size            string
	ExtraPrice      float64
	ExtraPercentage float64
}
type Item struct {
	Name  string
	Type  TypeItem
	Price float64
}

type Shop struct {
	Items []Item
}

func (shop *Shop) addItem(item Item) {
	var items []Item
	items = append(shop.Items, item)
	shop.Items = items
}

func NewTypeItem(tipo string) TypeItem {
	var typeItem TypeItem
	if tipo == "small" {
		typeItem = TypeItem{
			Size:            tipo,
			ExtraPrice:      0.0,
			ExtraPercentage: 0.0,
		}
	} else if tipo == "medium" {
		typeItem = TypeItem{
			Size:            tipo,
			ExtraPrice:      0.0,
			ExtraPercentage: 0.03,
		}
	} else if tipo == "big" {
		typeItem = TypeItem{
			Size:            tipo,
			ExtraPrice:      2500,
			ExtraPercentage: 0.06,
		}
	}
	return typeItem
}

func (item *Item) NewItem(typeItem TypeItem, name string, price float64) {
	item.Name = name
	item.Type = typeItem
	item.Price = price
}

type Producto interface {
	CalculateRate() float64
}
type Ecommerce interface {
	CalculateTotalRate() float64
	AgregarItem(item Item)
}

func (shop *Shop) CalculateTotalRate() float64 {
	var totalCost float64
	for _, value := range shop.Items {
		totalCost += value.CalculateRate()
	}
	return totalCost
}
func (item Item) CalculateRate() float64 {
	var cost float64
	var totalCost float64
	if item.Type.Size != small {
		if item.Type.Size == medium {
			cost = item.Price * item.Type.ExtraPercentage
			totalCost = cost + item.Price
		} else {
			cost = (item.Price * item.Type.ExtraPercentage) + item.Type.ExtraPrice
			totalCost = cost + item.Price
		}
	} else {
		fmt.Println("Small")
		totalCost = item.Price
	}
	return totalCost
}

func main() {

	typeItem := NewTypeItem(big)

	item := Item{}
	shop := Shop{}
	item.NewItem(typeItem, "Sugar", 30000)
	shop.addItem(item)
	shop.addItem(item)
	totalRate := shop.CalculateTotalRate()
	fmt.Println("El precio total de los productos : ", totalRate)
}

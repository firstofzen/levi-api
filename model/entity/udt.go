package entity

import "time"

type ProductCategory struct {
	l1 string
	l2 string
	l3 string
}
type ProductSale struct {
	CreateAt time.Time
	Expire   time.Time
	Discount float64
}
type ProductSku struct {
	Size  int
	Name  string
	Price string
}
type ProductRating struct {
	AccountId  int
	Comment    string
	PictureUrl string
	Score      float64
	CreateAt   time.Time
}
type ProductInventory struct {
	Total int
	Sold  int
}
type OrderShip struct {
	Price         float64
	ShippingUnit  string
	ShippingState string
	Description   string
	LastUpdate    time.Time
}

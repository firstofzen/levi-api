package entity

import (
	"time"
)

type Account struct {
	AccountId int       `json:"account_id,omitempty"`
	Email     string    `json:"email,omitempty"`
	Name      string    `json:"name,omitempty"`
	AvatarUrl string    `json:"avatar_url,omitempty"`
	IsViolate bool      `json:"is_violate,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Address   string    `json:"address,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	Bank      string    `json:"bank,omitempty"`
	Number    string    `json:"number,omitempty"`
	CreateAt  time.Time `json:"create_at"`
}
type AccountSocial struct {
	AccountId        int   `json:"account_id,omitempty"`
	BoothFollowingId []int `json:"booth_following_id,omitempty"`
	BoothLikeId      []int `json:"booth_like_id,omitempty"`
	RequestFriendId  []int `json:"request_friend_id,omitempty"`
	FriendId         []int `json:"friend_id,omitempty"`
}
type AccountBooth struct {
	AccountId             int    `json:"account_id,omitempty"`
	OrderIdWaitConfirm    []int  `json:"order_id_wait_confirm,omitempty"`
	PictureAvatarUrl      string `json:"picture_avatar_url,omitempty"`
	PictureDescriptionUrl string `json:"picture_description_url,omitempty"`
	Description           string `json:"description,omitempty"`
	IsAuthentic           bool   `json:"is_authentic,omitempty"`
}
type Product struct {
	ProductId           int              `json:"product_id,omitempty"`
	AccountId           int              `json:"account_id,omitempty"`
	Name                string           `json:"name,omitempty"`
	PictureAvatarUrl    string           `json:"picture_avatar_url,omitempty"`
	PicturesDescription []string         `json:"pictures_description,omitempty"`
	TotalScore          float64          `json:"total_score,omitempty"`
	Price               float64          `json:"price,omitempty"`
	Category            ProductCategory  `json:"category"`
	Sale                ProductSale      `json:"sale"`
	Inventory           ProductInventory `json:"inventory"`
	Skus                []ProductSku     `json:"skus,omitempty"`
	Ratings             []ProductRating  `json:"ratings,omitempty"`
	CreateAt            time.Time        `json:"create_at"`
}

type Order struct {
	OrderId     int       `json:"order_id,omitempty"`
	AccountId   int       `json:"account_id,omitempty"`
	ProductId   int       `json:"product_id,omitempty"`
	TotalAmount float64   `json:"total_amount,omitempty"`
	Status      string    `json:"status,omitempty"`
	LastUpdate  time.Time `json:"last_update"`
	Ship        OrderShip `json:"ship"`
	CreateAt    time.Time `json:"create_at"`
}

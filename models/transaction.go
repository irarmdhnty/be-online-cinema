package models

import "time"

type Transaction struct {
	ID     int          `json:"id"`
	UserID int          `json:"user_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User   User         `json:"user"`
	FilmID int          `json:"film_id" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Film   FilmResponse `json:"film"`
	Status string       `json:"status"`
	// for midtrans
	Price int `json:"price"`
	//BuyerID       int                  `json:"buyer_id"`
	//Buyer         UsersProfileResponse `json:"buyer"`
	//SellerID      int                  `json:"seller_id"`
	//Seller        UsersProfileResponse `json:"seller"`
	AccountNumber int       `json:"account_number"`
	TanggalOrder  time.Time `json:"tanggal_order" gorm:"default:Now()"`
}
type TransactionUserResponse struct {
	ID            int          `json:"id"`
	UserID        int          `json:"user_id"`
	User          User         `json:"user"`
	FilmID        int          `json:"film_id"`
	Film          FilmResponse `json:"film"`
	Status        string       `json:"status"`
	AccountNumber int          `json:"account_number"`
	TanggalOrder  time.Time    `json:"tanggal_order" gorm:"default:Now()"`
}

func (TransactionUserResponse) TableName() string {
	return "transactions"
}

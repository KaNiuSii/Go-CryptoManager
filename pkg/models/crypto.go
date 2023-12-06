package models

import (
	"github.com/KaNiuSii/Go-CryptoManager/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Crypto struct {
	gorm.Model
	Name   string `gorm:""json:"name"`
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Crypto{})
}

func (c *Crypto) CreateCrypto() *Crypto {
	db.NewRecord(c)
	db.Create(&c)
	return c
}

func GetAllCrypto() []Crypto {
	var Cryptos []Crypto
	db.Find(&Cryptos)
	return Cryptos
}

func GetCryptoById(Id int64) (*Crypto, *gorm.DB) {
	var getCrypto Crypto
	db := db.Where("ID=?", Id).Find(&getCrypto)
	return &getCrypto, db
}

func DeleteCrypto(Id int64) Crypto {
	var crypto Crypto
	db.Where("ID=?", Id).Delete(crypto)
	return crypto
}

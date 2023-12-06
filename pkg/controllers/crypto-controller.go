package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/KaNiuSii/Go-CryptoManager/pkg/models"
	"github.com/KaNiuSii/Go-CryptoManager/pkg/utils"
	"github.com/gorilla/mux"
)

var NewCrypto models.Crypto

func GetCrypto(w http.ResponseWriter, r *http.Request) {
	newCryptos := models.GetAllCrypto()
	res, _ := json.Marshal(newCryptos)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCryptoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cryptoId := vars["cryptoId"]
	ID, err := strconv.ParseInt(cryptoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	cryptoDetails, _ := models.GetCryptoById(ID)
	res, _ := json.Marshal(cryptoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateCrypto(w http.ResponseWriter, r *http.Request) {
	CreateCrypto := &models.Crypto{}
	utils.ParseBody(r, CreateCrypto)
	c := CreateCrypto.CreateCrypto()
	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCrypto(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cryptoId := vars["cryptoId"]
	ID, err := strconv.ParseInt(cryptoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing ")
	}
	crypto := models.DeleteCrypto(ID)
	res, _ := json.Marshal(crypto)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCrypto(w http.ResponseWriter, r *http.Request) {
	var updateCrypto = &models.Crypto{}
	utils.ParseBody(r, updateCrypto)
	vars := mux.Vars(r)
	cryptoId := vars["cryptoId"]
	ID, err := strconv.ParseInt(cryptoId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing ")
	}
	cryptoDetails, db := models.GetCryptoById(ID)
	if updateCrypto.Name != "" {
		cryptoDetails.Name = updateCrypto.Name
	}
	if updateCrypto.Symbol != "" {
		cryptoDetails.Symbol = updateCrypto.Symbol
	}
	if updateCrypto.Price != "" {
		cryptoDetails.Price = updateCrypto.Price
	}
	db.Save(&cryptoDetails)
	res, _ := json.Marshal(cryptoDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

package routes

import (
	"github.com/KaNiuSii/Go-CryptoManager/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterCryptoRoutes = func(router *mux.Router) {
	router.HandleFunc("/crypto/", controllers.CreateCrypto).Methods("POST")
	router.HandleFunc("/crypto/", controllers.GetCrypto).Methods("GET")
	router.HandleFunc("/crypto/{cryptoId}", controllers.GetCryptoById).Methods("GET")
	router.HandleFunc("/crypto/{cryptoId}", controllers.UpdateCrypto).Methods("PUT")
	router.HandleFunc("/crypto/{cryptoId}", controllers.DeleteCrypto).Methods("DELETE")
}

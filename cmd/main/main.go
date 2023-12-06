package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KaNiuSii/Go-CryptoManager/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterCryptoRoutes(r)
	http.Handle("/", r)
	port := ":8081"
	fmt.Println("listening on localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

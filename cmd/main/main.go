package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	"testGo/pkg/routes"
	"testGo/pkg/repository"
)
func main()  {
	fmt.Println("Запуск...")
	repository.CreateTable()
	router:=mux.NewRouter()
	routes.Router(router)
	log.Fatal(http.ListenAndServe(":8080",router))
}
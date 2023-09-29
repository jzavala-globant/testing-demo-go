package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jzavala-globant/testing-demo-go/internal/controllers"
	"github.com/jzavala-globant/testing-demo-go/internal/repositories"
	"github.com/jzavala-globant/testing-demo-go/internal/services"
	"github.com/jzavala-globant/testing-demo-go/pkg/apiconsumer"
)

type bookStoreController interface {
	GetBook(http.ResponseWriter, *http.Request)
	ListBooks(http.ResponseWriter, *http.Request)
}

type app struct {
	bookController bookStoreController
}

func main() {
	repo := repositories.NewBookRepository()
	services := services.NewBookStoreService(repo, apiconsumer.NewConsumer())
	app := app{
		bookController: controllers.NewBookStoreController(services),
	}

	app.startServer()
}

func (a *app) startServer() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/books/{id}", a.bookController.GetBook).Methods("GET")
	r.HandleFunc("/books", a.bookController.ListBooks).Methods("GET")
	fmt.Println("Server running")
	log.Fatal(http.ListenAndServe(":8080", r))
}

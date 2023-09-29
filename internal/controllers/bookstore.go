package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jzavala-globant/testing-demo-go/internal/models"
)

type Services interface {
	GetBook(context.Context, int64) (*models.APIGetBook, error)
	ListBooks(context.Context) ([]*models.APIGetBook, error)
}

type bookStoreController struct {
	s Services
}

func NewBookStoreController(s Services) *bookStoreController {
	return &bookStoreController{
		s,
	}
}

func (c *bookStoreController) GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	strID := vars["id"]

	id, err := strconv.ParseInt(strID, 10, 64)
	if err != nil {
		fmt.Printf("error parsing provided id %s %v\n", strID, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	response, err := c.s.GetBook(r.Context(), id)
	if err != nil {
		fmt.Printf("internal error getting book by id %s %v\n", strID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (c *bookStoreController) ListBooks(w http.ResponseWriter, r *http.Request) {
	response, err := c.s.ListBooks(r.Context())
	if err != nil {
		fmt.Printf("internal error listing books: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

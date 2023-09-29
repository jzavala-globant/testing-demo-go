package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/jzavala-globant/testing-demo-go/internal/models"
	"github.com/jzavala-globant/testing-demo-go/pkg/apiconsumer"
)

const (
	defaultCoverURL = "http://default.com"
	httpPrefix      = "http://"
	checkPricesURL  = "http://keepa/books/"
)

type Repository interface {
	GetBookByID(context.Context, int64) (*models.DBBook, error)
	ListBooks(context.Context) ([]*models.DBBook, error)
}

type bookStoreService struct {
	r Repository
	c *apiconsumer.Consumer
}

func NewBookStoreService(r Repository, c *apiconsumer.Consumer) *bookStoreService {
	return &bookStoreService{
		r,
		c,
	}
}

func (bss *bookStoreService) GetBook(ctx context.Context, id int64) (*models.APIGetBook, error) {
	book, err := bss.r.GetBookByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("error getting book from DB: %v", err)
	}

	if book == nil {
		return nil, fmt.Errorf("book not found")
	}

	book.CoverURL = formatCoverURL(book.CoverURL)

	return book.TranslateToAPIGetBook(), nil
}

func (bss *bookStoreService) ListBooks(ctx context.Context) ([]*models.APIGetBook, error) {
	books, err := bss.r.ListBooks(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting book from DB: %v", err)
	}

	if books == nil {
		return []*models.APIGetBook{}, nil
	}

	var res []*models.APIGetBook
	for _, book := range books {
		book.CoverURL = formatCoverURL(book.CoverURL)
		res = append(res, book.TranslateToAPIGetBook())
	}

	return res, nil
}

// dependency on a concrete
func (bss *bookStoreService) GetPriceHistory(ctx context.Context, id int64) (*models.HistoricPrice, error) {
	url := fmt.Sprintf("%s%d", checkPricesURL, id)
	status, body, err := bss.c.GetStats(url)
	if err != nil {
		return nil, fmt.Errorf("error getting historic prices for %d: %v", id, err)
	}

	history := new(models.HistoricPrice)
	err = json.Unmarshal(body, history)
	if err != nil {
		return nil, fmt.Errorf("error parsing response for %d: %v", id, err)
	}

	switch status {
	case http.StatusOK:
		return history, nil
	default:
		return nil, fmt.Errorf("unexpected response: %d - %s", status, string(body))
	}
}

func formatCoverURL(url string) string {
	if url == "" {
		return defaultCoverURL
	}

	if strings.HasPrefix(url, httpPrefix) {
		return url
	}

	return fmt.Sprintf("%s%s", httpPrefix, url)
}

package repositories

import (
	"context"
	"fmt"

	"github.com/jzavala-globant/testing-demo-go/internal/models"
)

var books = map[int64]*models.DBBook{
	1: {
		ID:         1,
		Title:      "To Kill a Mockingbird",
		Price:      12.99,
		CoverURL:   "http://example.com/to-kill-a-mockingbird",
		Author:     "Harper Lee",
		ISBN:       "978-0061120084",
		Year:       1960,
		NumOfPages: 281,
		CoverType:  "Hardcover",
		Summary:    "A classic novel set in the American South during the 1930s, dealing with issues of racial injustice and moral growth.",
	},
	2: {
		ID:         2,
		Title:      "1984",
		Price:      9.99,
		CoverURL:   "http://example.com/1984",
		Author:     "George Orwell",
		ISBN:       "978-0451524935",
		Year:       1949,
		NumOfPages: 328,
		CoverType:  "Paperback",
		Summary:    "A dystopian novel that explores totalitarianism and the consequences of government surveillance and censorship.",
	},
	3: {
		ID:         3,
		Title:      "The Great Gatsby",
		Price:      11.99,
		CoverURL:   "http://example.com/the-great-gatsby",
		Author:     "F. Scott Fitzgerald",
		ISBN:       "978-0743273565",
		Year:       1925,
		NumOfPages: 180,
		CoverType:  "Hardcover",
		Summary:    "A novel set in the Roaring Twenties, known for its portrayal of the American Dream and the excesses of the Jazz Age.",
	},
	4: {
		ID:         4,
		Title:      "Pride and Prejudice",
		Price:      10.99,
		CoverURL:   "http://example.com/pride-and-prejudice",
		Author:     "Jane Austen",
		ISBN:       "978-0141439518",
		Year:       1813,
		NumOfPages: 279,
		CoverType:  "Paperback",
		Summary:    "A classic novel of manners that explores themes of love, class, and societal expectations in 19th-century England.",
	},
	5: {
		ID:         5,
		Title:      "The Catcher in the Rye",
		Price:      8.99,
		CoverURL:   "example.com/the-catcher-in-the-rye",
		Author:     "J.D. Salinger",
		ISBN:       "978-0316769488",
		Year:       1951,
		NumOfPages: 224,
		CoverType:  "Paperback",
		Summary:    "A novel that follows the experiences of Holden Caulfield, a teenager dealing with alienation and rebellion in post-World War II America.",
	},
}

type bookRepository struct{}

func NewBookRepository() *bookRepository {
	return &bookRepository{}
}

func (br *bookRepository) GetBookByID(ctx context.Context, id int64) (*models.DBBook, error) {
	if b, ok := books[id]; ok {
		return b, nil
	}
	return nil, fmt.Errorf("DB error, book not found")
}

func (br *bookRepository) ListBooks(ctx context.Context) ([]*models.DBBook, error) {
	var res []*models.DBBook
	for _, book := range books {
		res = append(res, book)
	}
	return res, nil
}

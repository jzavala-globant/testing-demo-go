package models

type coverType string

const (
	HardCover coverType = "Hardcover"
	SoftCover coverType = "Paperback"
)

type DBBook struct {
	ID         int64
	Title      string
	Price      float64
	CoverURL   string
	Author     string
	ISBN       string
	Year       int
	NumOfPages int
	CoverType  coverType
	Summary    string
}

type APIGetBook struct {
	ID        int64
	Title     string
	Price     float64
	CoverURL  string
	CoverType string
}

func (db *DBBook) TranslateToAPIGetBook() *APIGetBook {
	return &APIGetBook{
		ID:        db.ID,
		Title:     db.Title,
		Price:     db.Price,
		CoverURL:  db.CoverURL,
		CoverType: string(db.CoverType),
	}
}

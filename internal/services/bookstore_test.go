package services

import (
	"context"
	"fmt"
	"testing"

	"github.com/jzavala-globant/testing-demo-go/internal/models"
	"github.com/jzavala-globant/testing-demo-go/internal/services/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BookstoreServiceSuite struct {
	suite.Suite
}

func (s *BookstoreServiceSuite) TestNewBookStoreService() {
	res := NewBookStoreService(nil, nil)
	s.NotNil(res)
}

func (s *BookstoreServiceSuite) TestGetBook() {
	r := mocks.NewRepository(&testing.T{})

	// db error
	r.On("GetBookByID", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("some error")).Once()
	service := NewBookStoreService(r, nil)
	book, err := service.GetBook(context.Background(), 1)
	s.Nil(book)
	s.Error(err)

	// book not found
	r.On("GetBookByID", mock.Anything, mock.Anything).Return(nil, nil).Once()
	// service = NewBookStoreService(r, nil)
	book, err = service.GetBook(context.Background(), 1)
	s.Nil(book)
	s.Empty(book)
	s.Error(err)

	// successfull response
	r.On("GetBookByID", mock.Anything, mock.Anything).Return(&models.DBBook{}, nil).Once()
	// service = NewBookStoreService(r, nil)
	book, err = service.GetBook(context.Background(), 1)
	s.NotNil(book)
	s.NoError(err)
}

func (s *BookstoreServiceSuite) TestListBook() {
	r := mocks.NewRepository(&testing.T{})

	// db error
	r.On("ListBooks", mock.Anything).Return(nil, fmt.Errorf("some error")).Once()
	service := NewBookStoreService(r, nil)
	books, err := service.ListBooks(context.Background())
	s.Nil(books)
	s.Error(err)

	// books not found
	r.On("ListBooks", mock.Anything).Return(nil, nil).Once()
	books, err = service.ListBooks(context.Background())
	s.NotNil(books)
	s.Empty(books)
	s.NoError(err)

	// successfull response
	r.On("ListBooks", mock.Anything).Return([]*models.DBBook{{ID: 1}}, nil).Once()
	books, err = service.ListBooks(context.Background())
	s.NotNil(books)
	s.NoError(err)
}

func TestBookstoreServiceSuite(t *testing.T) {
	suite.Run(t, new(BookstoreServiceSuite))
}

func Test_formatCoverURL(t *testing.T) {
	tests := []struct {
		name           string
		url            string
		expectedOutput string
	}{
		{
			name:           "Case 1 When empty input, default url output",
			url:            "",
			expectedOutput: defaultCoverURL,
		},
		{
			name:           "Case 2 When input includes protocol",
			url:            "http://example.com",
			expectedOutput: "http://example.com",
		},
		{
			name:           "Case 3 When input missed the protocol",
			url:            "example.com",
			expectedOutput: "http://example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedOutput, formatCoverURL(tt.url))
		})
	}
}

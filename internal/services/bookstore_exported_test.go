package services_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/jzavala-globant/testing-demo-go/internal/models"
	"github.com/jzavala-globant/testing-demo-go/internal/services"
	"github.com/jzavala-globant/testing-demo-go/internal/services/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type BookStoreServiceSuite struct {
	suite.Suite
}

func (s *BookStoreServiceSuite) TestNewBookStoreService() {
	res := services.NewBookStoreService(nil, nil)
	s.NotNil(res)
}

func (s *BookStoreServiceSuite) TestGetBookDBError() {
	r := newRepoMock()
	r.On("GetBookByID", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("some error"))
	service := services.NewBookStoreService(r, nil)
	book, err := service.GetBook(context.Background(), 1)
	s.Nil(book)
	s.Error(err)
}

func (s *BookStoreServiceSuite) TestGetBookNotFound() {
	r := newRepoMock()
	r.On("GetBookByID", mock.Anything, mock.Anything).Return(nil, nil)
	service := services.NewBookStoreService(r, nil)
	book, err := service.GetBook(context.Background(), 1)
	s.Nil(book)
	s.Error(err)
}

func (s *BookStoreServiceSuite) TestGetBookSuccessfully() {
	r := mocks.NewRepository(&testing.T{})
	r.On("GetBookByID", mock.Anything, mock.Anything).Return(&models.DBBook{}, nil)
	service := services.NewBookStoreService(r, nil)
	book, err := service.GetBook(context.Background(), 1)
	s.NotNil(book)
	s.NoError(err)
}

func TestBookStoreServiceSuite(t *testing.T) {
	suite.Run(t, new(BookStoreServiceSuite))
}

// manually created mock
type repoMock struct {
	mock.Mock
}

func newRepoMock() *repoMock {
	return &repoMock{}
}

func (m *repoMock) GetBookByID(ctx context.Context, id int64) (*models.DBBook, error) {
	args := m.Called(ctx, id)
	if _, ok := args.Get(0).(*models.DBBook); ok {
		return args.Get(0).(*models.DBBook), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *repoMock) ListBooks(ctx context.Context) ([]*models.DBBook, error) {
	args := m.Called(ctx)
	if _, ok := args.Get(0).(*models.DBBook); ok {
		return args.Get(0).([]*models.DBBook), args.Error(1)
	}
	return nil, args.Error(1)
}

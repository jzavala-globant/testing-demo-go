package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

// suite
type UrlFormaterSuite struct {
	suite.Suite
}

// executes before any suite execution
func (s *UrlFormaterSuite) SetupSuite() {
	fmt.Println("Configuring suite...")
}

// executes at the end of all tests
func (s *UrlFormaterSuite) TearDownSuite() {
	fmt.Println("Cleaning suite...")
}

// executes before each test
func (s *UrlFormaterSuite) SetupTest() {
	fmt.Println("Configuring test...")
}

// executes after each test
func (s *UrlFormaterSuite) TearDownTest() {
	fmt.Println("Cleaning test...")
}

func (s *UrlFormaterSuite) TestEmptyInput() {
	res := formatCoverURL("")
	s.Equal(defaultCoverURL, res)
}

func (s *UrlFormaterSuite) TestInputWithNoProtocol() {
	res := formatCoverURL("example.com")
	s.Equal("http://example.com", res)
}

func (s *UrlFormaterSuite) TestInputWithProtocol() {
	res := formatCoverURL("http://example.com")
	s.Equal("http://example.com", res)
}

// entrypoint for all test cases
func TestUrlFormaterSuite(t *testing.T) {
	suite.Run(t, new(UrlFormaterSuite))
}

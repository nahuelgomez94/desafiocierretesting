package products

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
)

type repositoryMock struct {
	productos []Product
	err       error
}

func NewRepositoryMock(p []Product, e error) *repositoryMock {
	return &repositoryMock{
		productos: p,
		err:       e,
	}
}

func (rm repositoryMock) GetAllBySeller(sellerID string) ([]Product, error) {
	return rm.productos, rm.err
}

func TestGetAllBySeller(t *testing.T) {
	testCases := []struct {
		Name          string
		SellerID      string
		Expected      interface{}
		ExpectedError error
		RepoMock      Repository
	}{
		{
			Name:          "Error en Repo",
			SellerID:      "id",
			Expected:      nil,
			ExpectedError: errors.New("Test"),
			RepoMock:      NewRepositoryMock(nil, errors.New("Test")),
		},
		{
			Name:          "Todo ok",
			SellerID:      "id",
			Expected:      []Product{},
			ExpectedError: nil,
			RepoMock:      NewRepositoryMock([]Product{}, nil),
		},
	}

	for i := range testCases {
		tc := testCases[i]

		serv := NewService(tc.RepoMock)
		t.Run(tc.Name, func(t *testing.T) {
			result, err := serv.GetAllBySeller(tc.SellerID)

			assert.Equal(t, tc.ExpectedError, err)
			assert.Equal(t, tc.Expected, result)
		})
	}
}

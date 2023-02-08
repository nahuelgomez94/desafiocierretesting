package products

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func CreateServerProducts(pr Repository) *gin.Engine {
	service := NewService(pr)
	handler := NewHandler(service)

	server := gin.Default()

	rg := server.Group("/api/v1")
	{
		rg.GET("", handler.GetProducts)
	}

	return server
}

func NewRequest(method, path, body string) (req *http.Request, res *httptest.ResponseRecorder) {
	req = httptest.NewRequest(method, path, strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	res = httptest.NewRecorder()

	return
}

func TestGetAllBySellerHandler(t *testing.T) {
	// Act
	testCases := []struct {
		Name           string
		Method         string
		Endpoint       string
		Body           string
		ExpectedCode   int
		ProductsInRepo []Product
		ErrMocked      error
	}{
		{
			Name:         "Ok",
			Method:       http.MethodGet,
			Endpoint:     "/api/v1?seller_id=1",
			Body:         "",
			ExpectedCode: 200,
		},
		{
			Name:         "No query",
			Method:       http.MethodGet,
			Endpoint:     "/api/v1",
			Body:         "",
			ExpectedCode: 400,
		},
		{
			Name:         "Internal Error",
			Method:       http.MethodGet,
			Endpoint:     "/api/v1?seller_id=error",
			Body:         "",
			ExpectedCode: 500,
			ErrMocked:    errors.New("Test"),
		},
	}

	// Assert
	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			server := CreateServerProducts(NewRepositoryMock(tc.ProductsInRepo, tc.ErrMocked))

			req, res := NewRequest(tc.Method, tc.Endpoint, tc.Body)

			server.ServeHTTP(res, req)
			// Assert
			assert.Equal(t, tc.ExpectedCode, res.Code)
		})
	}
}

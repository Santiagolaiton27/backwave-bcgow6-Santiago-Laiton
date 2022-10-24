package products

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/desafio-cierre-testing-SantiagoLaiton/internal/products"
	"github.com/Santiagolaiton27/backwave-bcgow6-Santiago-Laiton/desafio-cierre-testing-SantiagoLaiton/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockStore mocks.MockStore) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := products.NewHandler(service)

	r := gin.Default()
	pr := r.Group("/api/v1/products")
	pr.GET("", p.GetProducts)
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	requerid := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	requerid.Header.Add("Content-type", "application/json")
	return requerid, httptest.NewRecorder()
}
func TestGetAllProducts(t *testing.T) {
	//arrange
	mockStorage := mocks.MockStore{
		DataMock: []products.Product{{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
		},
	}
	var resp []products.Product
	r := createServer(mockStorage)
	request, rr := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=mock", "")

	//act
	r.ServeHTTP(rr, request)
	//assert
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Equal(t, len(mockStorage.DataMock), len(resp))
}

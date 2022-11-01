package seller

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/domain"
	"github.com/extmatperez/meli_bootcamp_go_w6-1/internal/seller/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	//arrange
	cx := gin.Context{}
	db := []domain.Seller{
		{
			ID:          9,
			CID:         345433,
			CompanyName: "Meli",
			Address:     "Calle 7 70-78",
			Telephone:   "3168747880",
		}, {
			ID:          96,
			CID:         34115433,
			CompanyName: "Meli",
			Address:     "Calle 7a 70-78",
			Telephone:   "3138912730",
		},
	}
	//act
	mockRepository := mocks.MockStorage{DataMock: db}
	service := NewService(&mockRepository)
	//assert
	seller, err := service.GetAll(cx)
	assert.Nil(t, err)
	assert.Equal(t, len(db), len(seller))
}
func TestGetById(t *testing.T) {
	//arrange
	cx := gin.Context{}
	db := []domain.Seller{
		{
			ID:          9,
			CID:         345433,
			CompanyName: "Meli",
			Address:     "Calle 7 70-78",
			Telephone:   "3168747880",
		}, {
			ID:          96,
			CID:         34115433,
			CompanyName: "Meli",
			Address:     "Calle 7a 70-78",
			Telephone:   "3138912730",
		},
	}
	//act
	mockRepository := mocks.MockStorage{DataMock: db}
	service := NewService(&mockRepository)
	seller, err := service.GetSellerById(cx, 9)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, db[0], seller)
}
func TestGetByIdFake(t *testing.T) {
	//arrange
	cx := gin.Context{}
	db := []domain.Seller{
		{
			ID:          9,
			CID:         345433,
			CompanyName: "Meli",
			Address:     "Calle 7 70-78",
			Telephone:   "3168747880",
		}, {
			ID:          96,
			CID:         34115433,
			CompanyName: "Meli",
			Address:     "Calle 7a 70-78",
			Telephone:   "3138912730",
		},
	}
	//act
	mockRepository := mocks.MockStorage{DataMock: db}
	service := NewService(&mockRepository)
	_, err := service.GetSellerById(cx, 10)
	//assert
	assert.NotNil(t, err)
	assert.Equal(t, "the seller not exist", err.Error())
}
func TestDelete(t *testing.T) {
	//arrange
	cx := gin.Context{}
	db := []domain.Seller{
		{
			ID:          9,
			CID:         345433,
			CompanyName: "Meli",
			Address:     "Calle 7 70-78",
			Telephone:   "3168747880",
		}, {
			ID:          96,
			CID:         34115433,
			CompanyName: "Meli",
			Address:     "Calle 7a 70-78",
			Telephone:   "3138912730",
		},
	}
	//act
	mockRepository := mocks.MockStorage{DataMock: db}
	service := NewService(&mockRepository)
	err := service.Delete(cx, 9)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(mockRepository.DataMock))
}
func TestDeleteFake(t *testing.T) {
	//arrange
	cx := gin.Context{}
	db := []domain.Seller{
		{
			ID:          9,
			CID:         345433,
			CompanyName: "Meli",
			Address:     "Calle 7 70-78",
			Telephone:   "3168747880",
		}, {
			ID:          96,
			CID:         34115433,
			CompanyName: "Meli",
			Address:     "Calle 7a 70-78",
			Telephone:   "3138912730",
		},
	}
	//act
	mockRepository := mocks.MockStorage{DataMock: db}
	service := NewService(&mockRepository)
	err := service.Delete(cx, 10)
	assert.NotNil(t, err)
	assert.Equal(t, "the seller don't exist", err.Error())
}
func TestCreate(t *testing.T) {
	//arrange
	cx := gin.Context{}
	db := []domain.Seller{
		{
			ID:          9,
			CID:         345433,
			CompanyName: "Meli",
			Address:     "Calle 7 70-78",
			Telephone:   "3168747880",
		}, {
			ID:          96,
			CID:         34115433,
			CompanyName: "Meli",
			Address:     "Calle 7a 70-78",
			Telephone:   "3138912730",
		},
	}
	sellerNew := domain.Seller{
		ID:          48,
		CID:         15,
		CompanyName: "COMSILA",
		Address:     "Carrera 8 #190 91",
		Telephone:   "3213403427",
	}
	//act
	mockRepository := mocks.MockStorage{DataMock: db}
	service := NewService(&mockRepository)
	_, err := service.Save(cx, sellerNew)
	assert.Nil(t, err)
	assert.Equal(t, 3, len(mockRepository.DataMock))
}
func TestCreateFake(t *testing.T) {
	//arrange
	cx := gin.Context{}
	db := []domain.Seller{
		{
			ID:          9,
			CID:         345433,
			CompanyName: "Meli",
			Address:     "Calle 7 70-78",
			Telephone:   "3168747880",
		}, {
			ID:          96,
			CID:         34115433,
			CompanyName: "Meli",
			Address:     "Calle 7a 70-78",
			Telephone:   "3138912730",
		},
	}
	sellerNew := domain.Seller{
		ID:          0,
		CID:         345433,
		CompanyName: "COMSILA",
		Address:     "Carrera 8 #190 91",
		Telephone:   "3213403427",
	}
	//act
	mockRepository := mocks.MockStorage{DataMock: db}
	service := NewService(&mockRepository)
	_, err := service.Save(cx, sellerNew)
	assert.NotNil(t, err)
	assert.Equal(t, 2, len(mockRepository.DataMock))
}
func TestUpdate(t *testing.T) {
	//arrange
	cx := gin.Context{}
	db := []domain.Seller{
		{
			ID:          9,
			CID:         345433,
			CompanyName: "Meli",
			Address:     "Calle 7 70-78",
			Telephone:   "3168747880",
		}, {
			ID:          96,
			CID:         34115433,
			CompanyName: "Meli",
			Address:     "Calle 7a 70-78",
			Telephone:   "3138912730",
		},
	}
	//act
	mockRepository := mocks.MockStorage{DataMock: db}
	service := NewService(&mockRepository)
	sellerNew := domain.Seller{
		ID:          0,
		CID:         345433,
		CompanyName: "COMSILA",
		Address:     "Carrera 8 #190 91",
		Telephone:   "3213403427",
	}
	service.Update(cx, sellerNew, 9)

}

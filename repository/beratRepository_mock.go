package repository

import (
	"errors"
	"time"

	"github.com/jagadwp/gin-crud-web/model"
	"github.com/stretchr/testify/mock"
)

type BeratRepositoryMock struct {
	Mock mock.Mock
}

func (b *BeratRepositoryMock) GetBeratById(id int) (*model.Berat, error) {
	arg := b.Mock.Called(id)
	if arg.Get(0) == nil && arg.Get(1) != nil {
		return nil, errors.New("query error")
	} else {
		berat := arg.Get(0).(model.Berat)
		return &berat, nil
	}
}

//GetBerats : get all berat
func (b *BeratRepositoryMock) GetAllBerat() (*[]model.Berat, error) {
	return nil, errors.New("query error")
}

//CreateBerat : create a berat
func (b *BeratRepositoryMock) CreateBerat(date time.Time, max, min, diff int) (*model.Berat, error) {
	return nil, errors.New("query error")
}

//UpdateBerat : edit a berat which find by ID
func (b *BeratRepositoryMock) UpdateBerat(berat *model.Berat) (*model.Berat, error) {
	return nil, errors.New("query error")
}

//DeleteBerat : delete a berat which find by ID
func (b *BeratRepositoryMock) DeleteBerat(berat *model.Berat) (*model.Berat, error) {
	return nil, errors.New("query error")
}

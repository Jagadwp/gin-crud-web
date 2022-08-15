package repository

import (
	"errors"
	"time"

	"github.com/jagadwp/gin-crud-web/helper"
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
	arg := b.Mock.Called()
	if arg.Get(0) == nil && arg.Get(1) != nil {
		return nil, errors.New("query error")
	} else {
		berat := arg.Get(0).([]model.Berat)
		return &berat, nil
	}
}

//CreateBerat : create a berat
func (b *BeratRepositoryMock) CreateBerat(date time.Time, max, min, diff int) (*model.Berat, error) {
	req := helper.BeratRequest{
		Max:  max,
		Min:  min,
		Date: date.Format("2006-01-02"),
	}

	arg := b.Mock.Called(&req)

	if arg.Get(0) == nil && arg.Get(1) != nil {
		return nil, errors.New("query error")
	} else {
		berat := arg.Get(0).(model.Berat)
		return &berat, nil
	}
}

//UpdateBerat : edit a berat which find by ID
func (b *BeratRepositoryMock) UpdateBerat(berat *model.Berat) (*model.Berat, error) {
	id := 3
	req := helper.BeratRequest{
		Max:  70,
		Min:  67,
		Date: "2022-12-04",
	}

	arg := b.Mock.Called(id, &req)

	if arg.Get(0) == nil && arg.Get(1) != nil {
		return nil, errors.New("query error")
	} else {
		berat := arg.Get(0).(model.Berat)
		return &berat, nil
	}
}

//DeleteBerat : delete a berat which find by ID
func (b *BeratRepositoryMock) DeleteBerat(berat *model.Berat) (*model.Berat, error) {
	arg := b.Mock.Called(3)

	if arg.Get(0) == nil && arg.Get(1) != nil {
		return nil, errors.New("query error")
	} else {
		berat := arg.Get(0).(model.Berat)
		return &berat, nil
	}
}

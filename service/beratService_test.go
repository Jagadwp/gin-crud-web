package service

import (
	"errors"
	"testing"
	"time"

	"github.com/jagadwp/gin-crud-web/helper"
	"github.com/jagadwp/gin-crud-web/model"
	"github.com/jagadwp/gin-crud-web/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var beratRepository = &repository.BeratRepositoryMock{Mock: mock.Mock{}}
var beratService = BeratService{BeratRepo: beratRepository}

func TestBeratService_GetAllBeratSuccess(t *testing.T) {
	date1, _ := time.Parse("2006-01-02", "2022-08-14")
	date2, _ := time.Parse("2006-01-02", "2022-08-15")
	berats := []model.Berat{
		{
			ID:   1,
			Max:  60,
			Min:  55,
			Diff: 5,
			Date: date1,
		},
		{
			ID:   2,
			Max:  77,
			Min:  75,
			Diff: 2,
			Date: date2,
		},
	}

	beratRepository.Mock.On("GetAllBerat").Return(berats, nil)

	result, err := beratService.GetAllBerat()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, berats[0].ID, result.Berat[0].ID)
	assert.Equal(t, berats[0].Max, result.Berat[0].Max)
	assert.Equal(t, berats[0].Min, result.Berat[0].Min)
	assert.Equal(t, berats[0].Diff, result.Berat[0].Diff)
	assert.Equal(t, berats[0].Date, result.Berat[0].Date)
}

func TestBeratService_GetBeratByIdSuccess(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2022-08-14")
	berat := model.Berat{
		ID:   3,
		Max:  60,
		Min:  55,
		Diff: 5,
		Date: date,
	}
	beratRepository.Mock.On("GetBeratById", 3).Return(berat, nil)

	result, err := beratService.GetBeratById(3)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, berat.ID, result.ID)
	assert.Equal(t, berat.Max, result.Max)
	assert.Equal(t, berat.Min, result.Min)
	assert.Equal(t, berat.Diff, result.Diff)
	assert.Equal(t, berat.Date, result.Date)
}

func TestBeratService_GetBeratByIdNotFound(t *testing.T) {
	beratRepository.Mock.On("GetBeratById", 999).Return(nil, errors.New("query error"))

	res, err := beratService.GetBeratById(999)
	assert.Nil(t, res)
	assert.NotNil(t, err)

}
func TestBeratService_CreateBeratSuccess(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2022-08-13")
	berat := model.Berat{
		ID:   1,
		Max:  60,
		Min:  55,
		Diff: 5,
		Date: date,
	}

	req := helper.BeratRequest{
		Max:  60,
		Min:  55,
		Date: "2022-08-13",
	}

	beratRepository.Mock.On("CreateBerat", &req).Return(berat, nil)

	res, err := beratService.CreateBerat(&req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, berat.ID, res.ID)
	assert.Equal(t, berat.Max, res.Max)
	assert.Equal(t, berat.Min, res.Min)
	assert.Equal(t, berat.Diff, res.Diff)
	assert.Equal(t, berat.Date, res.Date)
}

func TestBeratService_CreateBeratInputNotValid(t *testing.T) {
	req := helper.BeratRequest{
		Max:  -10,
		Min:  -1,
		Date: "22-44-99",
	}

	msg := "Required fields are empty or not valid"
	beratRepository.Mock.On("CreateBerat", &req).Return(nil, errors.New(msg))

	res, err := beratService.CreateBerat(&req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestBeratService_UpdateBeratSuccess(t *testing.T) {
	TestBeratService_GetBeratByIdSuccess(t)
	id := 3
	date, _ := time.Parse("2006-01-02", "2022-12-04")
	berat := model.Berat{
		ID:   3,
		Max:  70,
		Min:  67,
		Diff: 3,
		Date: date,
	}

	req := helper.BeratRequest{
		Max:  70,
		Min:  67,
		Date: "2022-12-04",
	}

	beratRepository.Mock.On("UpdateBerat", id, &req).Return(berat, nil)

	res, err := beratService.UpdateBerat(id, &req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, berat.ID, res.ID)
	assert.Equal(t, berat.Max, res.Max)
	assert.Equal(t, berat.Min, res.Min)
	assert.Equal(t, berat.Diff, res.Diff)
	assert.Equal(t, berat.Date, res.Date)
}

func TestBeratService_UpdateBeratInputNotValid(t *testing.T) {
	TestBeratService_GetBeratByIdSuccess(t)
	id := 3
	req := helper.BeratRequest{
		Max:  2,
		Min:  3,
		Date: "2022-112-004",
	}

	msg := "Required fields are empty or not valid"
	beratRepository.Mock.On("UpdateBerat", id, &req).Return(nil, msg)

	res, err := beratService.UpdateBerat(id, &req)
	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestBeratService_DeleteBeratSuccess(t *testing.T) {
	TestBeratService_GetBeratByIdSuccess(t)
	date, _ := time.Parse("2006-01-02", "2022-08-14")
	berat := model.Berat{
		ID:   3,
		Max:  60,
		Min:  55,
		Diff: 5,
		Date: date,
	}

	beratRepository.Mock.On("DeleteBerat", 3).Return(berat, nil)
	
	result, err := beratService.DeleteBerat(3)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, berat.ID, result.ID)
	assert.Equal(t, berat.Max, result.Max)
	assert.Equal(t, berat.Min, result.Min)
	assert.Equal(t, berat.Diff, result.Diff)
	assert.Equal(t, berat.Date, result.Date)
}

func TestBeratService_DeleteBeratNotFound(t *testing.T) {
	TestBeratService_GetBeratByIdNotFound(t)

	beratRepository.Mock.On("DeleteBerat", 999).Return(nil,)
	
	result, err := beratService.DeleteBerat(999)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}
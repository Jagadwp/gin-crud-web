package service

import (
	"errors"
	"math"
	"time"

	"github.com/jagadwp/gin-crud-web/helper"
	"github.com/jagadwp/gin-crud-web/repository"
)

type BeratService struct {
	BeratRepo repository.IBeratRepository
}

func NewBeratService(BeratRepo repository.IBeratRepository) *BeratService {
	return &BeratService{BeratRepo: BeratRepo}
}

//GetBeratById : get a berat which find by ID
func (b *BeratService) GetBeratById(id int) (*helper.BeratResponse, error) {
	berat, err := b.BeratRepo.GetBeratById(id)

	if err != nil {
		return nil, err
	}

	return &helper.BeratResponse{
		ID:   berat.ID,
		Min:  berat.Min,
		Max:  berat.Max,
		Diff: berat.Diff,
		Date: berat.Date,
	}, nil
}

//GetBerats : get all berat
func (b *BeratService) GetAllBerat() (*helper.IndexResponse, error) {
	berats, err := b.BeratRepo.GetAllBerat()

	if err != nil {
		return nil, err
	}

	var data helper.IndexResponse

	lenData := float64(len(*berats))
	totMax, totMin, totDiff := 0, 0, 0

	for _, berat := range *berats {
		tempData := helper.BeratResponse{
			ID:   berat.ID,
			Min:  berat.Min,
			Max:  berat.Max,
			Diff: berat.Diff,
			Date: berat.Date,
		}

		totMax += berat.Max
		totMin += berat.Min
		totDiff += berat.Diff

		data.Berat = append(data.Berat, tempData)
	}

	data.Max = roundFloat((float64(totMax) / lenData), 2)
	data.Min = roundFloat((float64(totMin) / lenData), 2)
	data.Diff = roundFloat((float64(totDiff) / lenData), 2)

	return &data, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

// //CreateBerat : create a berat
func (b *BeratService) CreateBerat(req *helper.BeratRequest) (*helper.BeratResponse, error) {
	diff := req.Max - req.Min
	date, err := time.Parse("2006-01-02", req.Date)

	if err != nil {
		return nil, err
	}

	berat, err := b.BeratRepo.CreateBerat(date, req.Max, req.Min, diff)

	if err != nil {
		return nil, err
	}

	return &helper.BeratResponse{
		ID:   berat.ID,
		Max:  berat.Max,
		Min:  berat.Min,
		Diff: berat.Diff,
		Date: berat.Date,
	}, nil
}

//UpdateBerat : edit a berat which find by ID
func (b *BeratService) UpdateBerat(id int, req *helper.BeratRequest) (*helper.BeratResponse, error) {
	beratTemp, err := b.BeratRepo.GetBeratById(id)
	if err != nil {
		return nil, err
	}

	if req.Date != "" {
		date, err := time.Parse("2006-01-02", req.Date)

		if err != nil {
			return nil, err
		}
		beratTemp.Date = date
	}

	if req.Max != 0 {
		beratTemp.Max = req.Max
	}

	if req.Min != 0 {
		beratTemp.Min = req.Min
	}

	beratTemp.Diff = beratTemp.Max - beratTemp.Min
	if beratTemp.Diff < 0 {
		return nil, errors.New("min cannot be greater than max")
	}
	beratTemp.UpdatedAt = time.Now()

	beratUpdated, err := b.BeratRepo.UpdateBerat(beratTemp)

	if err != nil {
		return nil, err
	}

	return &helper.BeratResponse{
		ID:   beratUpdated.ID,
		Max:  beratUpdated.Max,
		Min:  beratUpdated.Min,
		Diff: beratUpdated.Diff,
		Date: beratUpdated.Date,
	}, nil
}

//DeleteBerat : delete a berat which find by ID
func (s *BeratService) DeleteBerat(id int) (*helper.BeratResponse, error) {
	berat, err := s.BeratRepo.GetBeratById(id)

	if err != nil {
		return nil, err
	}

	berat, err = s.BeratRepo.DeleteBerat(berat)

	if err != nil {
		return nil, err
	}

	return &helper.BeratResponse{
		ID:   berat.ID,
		Max:  berat.Max,
		Min:  berat.Min,
		Diff: berat.Diff,
		Date: berat.Date,
	}, nil
}

//isUniquePhone : check if the phone number already taken
// func isUniquePhone(berat model.Berat, isUpdate bool) bool {
// 	var res model.Berat
// 	db := database.GetDB()
// 	if isUpdate {
// 		db.Where("phone = ? AND id <> ?", berat.Phone, berat.ID).First(&res)
// 	} else {
// 		db.Where("phone = ?", berat.Phone).First(&res)
// 	}

// 	if res.ID != 0 {

// 		return false
// 	}

// 	return true
// }

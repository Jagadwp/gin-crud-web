package repository

import (
	"time"

	"github.com/jagadwp/gin-crud-web/model"
	"github.com/jinzhu/gorm"
)

type BeratRepository struct {
	db *gorm.DB
}

func NewBeratRepository(db *gorm.DB) *BeratRepository {
	return &BeratRepository{db: db}
}

//GetBeratById : get a berat which find by ID
func (b *BeratRepository) GetBeratById(id int) (*model.Berat, error) {
	var berat model.Berat

	if err := b.db.First(&berat, id).Error; err != nil {
		return nil, err
	}

	return &berat, nil
}

//GetBerats : get all berat
func (b *BeratRepository) GetAllBerat() (*[]model.Berat, error) {
	var berats []model.Berat

	if err := b.db.Order("date").Find(&berats).Error; err != nil {
		return nil, err
	}

	return &berats, nil
}

//CreateBerat : create a berat
func (b *BeratRepository) CreateBerat(date time.Time, max, min, diff int) (*model.Berat, error) {
	var berat *model.Berat = &model.Berat{
		Max:       max,
		Min:       min,
		Diff:      diff,
		Date:      date,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := b.db.Create(&berat).Error; err != nil {
		return nil, err
	}

	return berat, nil
}

//UpdateBerat : edit a berat which find by ID
func (b *BeratRepository) UpdateBerat(berat *model.Berat) (*model.Berat, error) {
	if err := b.db.Save(berat).Error; err != nil {
		return &model.Berat{}, err
	}

	return berat, nil
}

//DeleteBerat : delete a berat which find by ID
func (b *BeratRepository) DeleteBerat(berat *model.Berat) (*model.Berat, error) {
	if err := b.db.Delete(berat).Error; err != nil {
		return &model.Berat{}, err
	}

	return berat, nil
}

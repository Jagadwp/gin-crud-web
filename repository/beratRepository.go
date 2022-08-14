package repository

import (
	"fmt"

	"github.com/bonggar/gorestapi/model"
	"github.com/jinzhu/gorm"
)

type BeratRepository struct {
	db *gorm.DB
}

func NewBeratRepository(db *gorm.DB) *BeratRepository {
	return &BeratRepository{db: db}
}

// //isUniquePhone : check if the phone number already taken
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

//GetBerat : get a berat which find by ID
// func GetBerat(c *gin.Context) {
// 	db := database.GetDB()
// 	id := c.Params.ByName("id")
// 	var berat model.Berat
// 	db.First(&berat, id)

// 	if berat.ID != 0 {
// 		helper.RespondJSON(c, 200, "", berat, nil)
// 	} else {
// 		helper.RespondJSON(c, 404, "Berat not found", nil, nil)
// 	}
// }

//GetBerats : get all berat
func (b *BeratRepository) GetAllBerat() (*[]model.Berat, error) {
	var berats []model.Berat
	fmt.Println("MASUKKK REPO")

	if err := b.db.Find(&berats).Error; err != nil {
		return nil, err
	}

	return &berats, nil
}

// //CreateBerat : create a berat
// func CreateBerat(c *gin.Context) {
// 	var berat model.Berat

// 	if err := c.ShouldBind(&berat); err != nil {
// 		errorFields := helper.ConstructErrors(err)
// 		helper.RespondJSON(c, 422, "Could not create berat", nil, errorFields)
// 		return
// 	}

// 	db := database.GetDB()
// 	//manually check phone unique constraint
// 	if !isUniquePhone(berat, false) {
// 		errorFields := [1]helper.ErrorField{{ID: "phone", Message: "Phone already taken"}}
// 		helper.RespondJSON(c, 422, "Could not create berat", nil, errorFields)
// 		return
// 	}

// 	//finally create the berat
// 	db.Create(&berat)

// 	if berat.ID != 0 {
// 		helper.RespondJSON(c, 201, "Berat has been created successfully", berat, nil)
// 	} else {
// 		helper.RespondJSON(c, 409, "Can not create berat", nil, nil)
// 	}
// }

// //UpdateBerat : edit a berat which find by ID
// func UpdateBerat(c *gin.Context) {
// 	var berat model.Berat
// 	db := database.GetDB()
// 	id := c.Params.ByName("id")
// 	db.First(&berat, id)

// 	if berat.ID != 0 {
// 		var updatedBerat model.Berat
// 		c.ShouldBind(&updatedBerat)

// 		//manually check email unique constraint
// 		if !isUniqueEmail(updatedBerat, true) {
// 			errorFields := [1]helper.ErrorField{{ID: "email", Message: "Email already taken"}}
// 			helper.RespondJSON(c, 422, "Could not create berat", nil, errorFields)
// 			return
// 		}

// 		//manually check phone unique constraint
// 		if !isUniquePhone(updatedBerat, true) {
// 			errorFields := [1]helper.ErrorField{{ID: "phone", Message: "Phone already taken"}}
// 			helper.RespondJSON(c, 422, "Could not create berat", nil, errorFields)
// 			return
// 		}

// 		//finally update the berat
// 		db.Model(&berat).Updates(updatedBerat)

// 		helper.RespondJSON(c, 200, "Berat has been updated successfully", berat, nil)
// 	} else {
// 		helper.RespondJSON(c, 404, "Berat not found", nil, nil)
// 	}
// }

// //DeleteBerat : delete a berat which find by ID
// func DeleteBerat(c *gin.Context) {
// 	var berat model.Berat
// 	db := database.GetDB()
// 	id := c.Params.ByName("id")
// 	db.First(&berat, id)

// 	if berat.ID != 0 {
// 		db.Delete(&berat)
// 		helper.RespondJSON(c, 200, "Berat has been deleted successfully", nil, nil)
// 	} else {
// 		helper.RespondJSON(c, 404, "Berat not found", nil, nil)
// 	}
// }

// //OptionsBerat : supporting options for CORS
// func OptionsBerat(c *gin.Context) {
// 	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
// 	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
// 	c.Next()
// }

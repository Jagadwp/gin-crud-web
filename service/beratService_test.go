package service

// var beratRepository = &repository.BeratRepositoryMock{Mock: mock.Mock{}}
// var beratService = BeratService{BeratRepo: beratRepository}

// func TestBeratService_GetNotFound(t *testing.T) {

// 	// program mock
// 	beratRepository.Mock.On("FindById", 1).Return(nil, "ok")

// 	berat, err := beratService.Get(1)
// 	assert.Nil(t, berat)
// 	assert.NotNil(t, err)

// }

// func TestBeratService_GetSuccess(t *testing.T) {
// 	berat := entity.Berat{
// 		Id:   "1",
// 		Name: "Laptop",
// 	}
// 	beratRepository.Mock.On("FindById", 2).Return(berat)

// 	result, err := beratService.Get(2)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, result)
// 	assert.Equal(t, berat.Id, result.Id)
// 	assert.Equal(t, berat.Name, result.Name)
// }

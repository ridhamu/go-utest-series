package service

import (
	"belajar-golang-utest/entity"
	"belajar-golang-utest/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindById", "1").Return(nil) // simulate nil return
	category, err := categoryService.Get("1")

	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSucceed(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindById", "2").Return(category)
	result, err := categoryService.Get("2")
	
	// checking here
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.Id, category.Id)
	assert.Equal(t, result.Name, category.Name)
}

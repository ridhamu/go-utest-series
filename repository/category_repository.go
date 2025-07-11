package repository

import "belajar-golang-utest/entity"

// create repository for category struct
type CategoryRepository interface {
	FindById(id string) *entity.Category
}

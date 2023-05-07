package domain

import (
	"fmt"
	"github.com/ningenMe/miiko-api/pkg/infra"
)

type CategoryService struct{}

var categoryRepository = infra.CategoryRepository{}

func (CategoryService) GetCategoryByCategoryId(categoryId string) (*infra.CategoryDto, error) {
	categoryDtoList := categoryRepository.GetList()
	for _, categoryDto := range categoryDtoList {
		if categoryId == categoryDto.CategoryId {
			return categoryDto, nil
		}
	}
	return nil, fmt.Errorf("category not found")
}

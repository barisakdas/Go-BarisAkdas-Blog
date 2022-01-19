package Repository

import (
	model "Go-BarisAkdas-Blog/Domain/Model"
	"time"
)

type ICategoryRepository interface {
	GetAllCategories() ([]CategoryDTO, error)
	GetAllActiveCategories() ([]CategoryDTO, error)
	GetCategoryById(id int) (CategoryDTO, error)
	CreateCategory(newCategoryDto CategoryDTO) (CategoryDTO, error)
	UpdateCategory(updateCategoryDto CategoryDTO) (CategoryDTO, error)
	DeleteCategory(id int) (CategoryDTO, error)
}

type CategoryRepository struct{}

type CategoryDTO struct {
	ID          uint
	Name        string
	CreatedDate time.Time
	IsDeleted   bool
}

func NewCategoryRepository() ICategoryRepository {
	return &CategoryRepository{}
}

func (c CategoryRepository) GetAllCategories() ([]CategoryDTO, error) {
	data, err := model.Category{}.GetAll()
	if err != nil {
		return nil, err
	}

	var dataDtos []CategoryDTO

	for _, value := range data {
		newData := CategoryDTO{
			ID:          value.ID,
			Name:        value.Name,
			CreatedDate: value.CreatedAt,
			IsDeleted:   value.IsDeleted,
		}
		dataDtos = append(dataDtos, newData)
	}

	return dataDtos, nil
}

func (c CategoryRepository) GetAllActiveCategories() ([]CategoryDTO, error) {
	data, err := model.Category{}.GetAll("is_deleted = false")
	if err != nil {
		return nil, err
	}

	var dataDtos []CategoryDTO

	for _, value := range data {
		newData := CategoryDTO{
			ID:          value.ID,
			Name:        value.Name,
			CreatedDate: value.CreatedAt,
			IsDeleted:   value.IsDeleted,
		}
		dataDtos = append(dataDtos, newData)
	}

	return dataDtos, nil
}
func (c CategoryRepository) GetCategoryById(id int) (CategoryDTO, error) {
	data, err := model.Category{}.Get(id)
	if err != nil {
		return CategoryDTO{}, err
	}

	return CategoryDTO{
		ID:          data.ID,
		Name:        data.Name,
		CreatedDate: data.CreatedAt,
		IsDeleted:   data.IsDeleted,
	}, nil
}

func (c CategoryRepository) CreateCategory(newCategoryDto CategoryDTO) (CategoryDTO, error) {
	result, err := model.Category{
		Name:      newCategoryDto.Name,
		IsDeleted: newCategoryDto.IsDeleted}.Add()
	if err != nil {
		return CategoryDTO{}, err
	}

	return CategoryDTO{
		ID:          result.ID,
		Name:        result.Name,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted,
	}, nil
}

func (c CategoryRepository) UpdateCategory(updateCategoryDto CategoryDTO) (CategoryDTO, error) {
	data, err := model.Category{}.Get(updateCategoryDto.ID)
	if err != nil {
		return CategoryDTO{}, err
	}

	result, err := data.Update(model.Category{
		Name:      data.Name,
		IsDeleted: data.IsDeleted,
	})
	if err != nil {
		return CategoryDTO{}, err
	}

	return CategoryDTO{
		ID:          result.ID,
		Name:        result.Name,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted,
	}, nil

}

func (c CategoryRepository) DeleteCategory(id int) (CategoryDTO, error) {
	data, err := model.Category{}.Get(id)
	if err != nil {
		return CategoryDTO{}, err
	}

	result, err := data.Update(model.Category{
		IsDeleted: true,
	})
	if err != nil {
		return CategoryDTO{}, err
	}

	return CategoryDTO{
		ID:          result.ID,
		Name:        result.Name,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted,
	}, nil
}

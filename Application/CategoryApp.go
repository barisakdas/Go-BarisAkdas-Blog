package Application

import repo "Go-BarisAkdas-Blog/Domain/Repository"

type CategoryApplication struct{}

var categoryRepo = repo.NewCategoryRepository()

func (c CategoryApplication) GetAllCategories() ([]repo.CategoryDTO, error) {
	return categoryRepo.GetAllCategories()
}
func (c CategoryApplication) GetAllActiveCategories() ([]repo.CategoryDTO, error) {
	return categoryRepo.GetAllActiveCategories()
}
func (c CategoryApplication) GetCategoryById(id int) (repo.CategoryDTO, error) {
	return categoryRepo.GetCategoryById(id)
}
func (c CategoryApplication) CreateCategory(newCategoryDto repo.CategoryDTO) (repo.CategoryDTO, error) {
	return categoryRepo.CreateCategory(newCategoryDto)
}
func (c CategoryApplication) UpdateCategory(updateCategoryDto repo.CategoryDTO) (repo.CategoryDTO, error) {
	return categoryRepo.UpdateCategory(updateCategoryDto)
}
func (c CategoryApplication) DeleteCategory(id int) (repo.CategoryDTO, error) {
	return categoryRepo.DeleteCategory(id)
}

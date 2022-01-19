package Application

import repo "Go-BarisAkdas-Blog/Domain/Repository"

type AboutApplication struct{}

var aboutRepo = repo.NewAboutRepositoory()

func (a AboutApplication) GetAllAbouts() ([]repo.AboutDTO, error) {
	return aboutRepo.GetAllAbouts()
}
func (a AboutApplication) GetAllActiveAbouts() ([]repo.AboutDTO, error) {
	return aboutRepo.GetAllActiveAbouts()
}
func (a AboutApplication) GetAboutById(id int) (repo.AboutDTO, error) {
	return aboutRepo.GetAboutById(id)
}
func (a AboutApplication) CreateAbout(about *repo.AboutDTO) (repo.AboutDTO, error) {
	return aboutRepo.CreateAbout(about)
}
func (a AboutApplication) UpdateAbout(about *repo.AboutDTO) (repo.AboutDTO, error) {
	return aboutRepo.UpdateAbout(about)
}
func (a AboutApplication) DeleteAbout(id int) (repo.AboutDTO, error) {
	return aboutRepo.DeleteAbout(id)
}

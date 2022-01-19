package Application

import repo "Go-BarisAkdas-Blog/Domain/Repository"

type UserApplication struct{}

var userRepo = repo.NewUserRepository()

func (u UserApplication) GetAllUsers() ([]repo.UserDTO, error) {
	return userRepo.GetAllUsers()
}
func (u UserApplication) GetAllActiveUsers() ([]repo.UserDTO, error) {
	return userRepo.GetAllActiveUsers()
}
func (u UserApplication) GetUserById(id int) (repo.UserDTO, error) {
	return userRepo.GetUserById(id)
}
func (u UserApplication) CreateUser(newUserDto repo.UserDTO) (repo.UserDTO, error) {
	return userRepo.CreateUser(newUserDto)
}
func (u UserApplication) UpdateUser(updateUserDto repo.UserDTO) (repo.UserDTO, error) {
	return userRepo.UpdateUser(updateUserDto)
}
func (u UserApplication) DeleteUser(id int) (repo.UserDTO, error) {
	return userRepo.DeleteUser(id)
}

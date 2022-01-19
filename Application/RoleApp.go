package Application

import repo "Go-BarisAkdas-Blog/Domain/Repository"

type RoleApplication struct{}

var roleRepo = repo.NewRoleRepository()

func (r RoleApplication) GetAllRoles() ([]repo.RoleDTO, error) {
	return roleRepo.GetAllRoles()
}
func (r RoleApplication) GetAllActiveRoles() ([]repo.RoleDTO, error) {
	return roleRepo.GetAllActiveRoles()
}
func (r RoleApplication) GetRoleById(id int) (repo.RoleDTO, error) {
	return roleRepo.GetRoleById(id)
}
func (r RoleApplication) CreateRole(newroleDto repo.RoleDTO) (repo.RoleDTO, error) {
	return roleRepo.CreateRole(newroleDto)
}
func (r RoleApplication) UpdateRole(updateRoleDto repo.RoleDTO) (repo.RoleDTO, error) {
	return roleRepo.UpdateRole(updateRoleDto)
}
func (r RoleApplication) DeleteRole(id int) (repo.RoleDTO, error) {
	return roleRepo.DeleteRole(id)
}

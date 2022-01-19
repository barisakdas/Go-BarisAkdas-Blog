package Repository

import (
	model "Go-BarisAkdas-Blog/Domain/Model"
	"time"
)

type IRoleRepository interface {
	GetAllRoles() ([]RoleDTO, error)
	GetAllActiveRoles() ([]RoleDTO, error)
	GetRoleById(id int) (RoleDTO, error)
	CreateRole(newRoleDto RoleDTO) (RoleDTO, error)
	UpdateRole(updateRoleDto RoleDTO) (RoleDTO, error)
	DeleteRole(id int) (RoleDTO, error)
}

type RoleRepository struct{}

type RoleDTO struct {
	Id          uint
	Name        string
	Description string
	SelectAuth  bool
	CreateAuth  bool
	UpdateAuth  bool
	DeleteAuth  bool
	CreatedDate time.Time
	IsDeleted   bool
}

func NewRoleRepository() IRoleRepository {
	return &RoleRepository{}
}

func (r RoleRepository) GetAllRoles() ([]RoleDTO, error) {
	data, err := model.Role{}.GetAll()
	if err != nil {
		return nil, err
	}

	var dataDtos []RoleDTO

	for _, value := range data {
		newData := RoleDTO{
			Id:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			SelectAuth:  value.SelectAuth,
			CreateAuth:  value.CreateAuth,
			UpdateAuth:  value.UpdateAuth,
			DeleteAuth:  value.DeleteAuth,
			CreatedDate: value.CreatedAt,
			IsDeleted:   value.IsDeleted,
		}

		dataDtos = append(dataDtos, newData)
	}

	return dataDtos, nil
}

func (r RoleRepository) GetAllActiveRoles() ([]RoleDTO, error) {
	data, err := model.Role{}.GetAll("is_deleted = false")
	if err != nil {
		return nil, err
	}

	var dataDtos []RoleDTO

	for _, value := range data {
		newData := RoleDTO{
			Id:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			SelectAuth:  value.SelectAuth,
			CreateAuth:  value.CreateAuth,
			UpdateAuth:  value.UpdateAuth,
			DeleteAuth:  value.DeleteAuth,
			CreatedDate: value.CreatedAt,
			IsDeleted:   value.IsDeleted,
		}

		dataDtos = append(dataDtos, newData)
	}

	return dataDtos, nil
}

func (r RoleRepository) GetRoleById(id int) (RoleDTO, error) {
	data, err := model.Role{}.Get(id)
	if err != nil {
		return RoleDTO{}, err
	}

	return RoleDTO{
		Id:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		SelectAuth:  data.SelectAuth,
		CreateAuth:  data.CreateAuth,
		UpdateAuth:  data.UpdateAuth,
		DeleteAuth:  data.DeleteAuth,
		CreatedDate: data.CreatedAt,
		IsDeleted:   data.IsDeleted,
	}, nil
}

func (r RoleRepository) CreateRole(newRoleDto RoleDTO) (RoleDTO, error) {
	result, err := model.Role{
		Name:        newRoleDto.Name,
		Description: newRoleDto.Description,
		SelectAuth:  newRoleDto.SelectAuth,
		CreateAuth:  newRoleDto.CreateAuth,
		UpdateAuth:  newRoleDto.UpdateAuth,
		DeleteAuth:  newRoleDto.DeleteAuth,
		IsDeleted:   newRoleDto.IsDeleted,
	}.Add()
	if err != nil {
		return RoleDTO{}, err
	}

	return RoleDTO{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		SelectAuth:  result.SelectAuth,
		CreateAuth:  result.CreateAuth,
		UpdateAuth:  result.UpdateAuth,
		DeleteAuth:  result.DeleteAuth,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted,
	}, nil
}

func (r RoleRepository) UpdateRole(updateRoleDto RoleDTO) (RoleDTO, error) {
	data, err := model.Role{}.Get(updateRoleDto.Id)
	if err != nil {
		return RoleDTO{}, err
	}

	result, err := data.Update(model.Role{
		Name:        updateRoleDto.Name,
		Description: updateRoleDto.Description,
		SelectAuth:  updateRoleDto.SelectAuth,
		CreateAuth:  updateRoleDto.CreateAuth,
		UpdateAuth:  updateRoleDto.UpdateAuth,
		DeleteAuth:  updateRoleDto.DeleteAuth,
		IsDeleted:   updateRoleDto.IsDeleted,
	})

	return RoleDTO{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		SelectAuth:  result.SelectAuth,
		CreateAuth:  result.CreateAuth,
		UpdateAuth:  result.UpdateAuth,
		DeleteAuth:  result.DeleteAuth,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted,
	}, nil
}

func (r RoleRepository) DeleteRole(id int) (RoleDTO, error) {
	data, err := model.Role{}.Get(id)
	if err != nil {
		return RoleDTO{}, err
	}

	result, err := data.Update(model.Role{
		IsDeleted: true,
	})

	return RoleDTO{
		Id:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		SelectAuth:  result.SelectAuth,
		CreateAuth:  result.CreateAuth,
		UpdateAuth:  result.UpdateAuth,
		DeleteAuth:  result.DeleteAuth,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted,
	}, nil
}

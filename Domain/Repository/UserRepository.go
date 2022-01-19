package Repository

import (
	model "Go-BarisAkdas-Blog/Domain/Model"
	"time"
)

type IUserRepository interface {
	GetAllUsers() ([]UserDTO, error)
	GetAllActiveUsers() ([]UserDTO, error)
	GetUserById(id int) (UserDTO, error)
	CreateUser(newUserDto UserDTO) (UserDTO, error)
	UpdateUser(updateUserDto UserDTO) (UserDTO, error)
	DeleteUser(id int) (UserDTO, error)
}

type UserRepository struct{}

type UserDTO struct {
	ID          uint
	FirstName   string
	LastName    string
	UserName    string
	Password    string
	BirthDate   string
	RoleId      int
	CreatedDate time.Time
	IsDeleted   bool
}

func NewUserRepository() IUserRepository {
	return &UserRepository{}
}

func (u UserRepository) GetAllUsers() ([]UserDTO, error) {
	data, err := model.User{}.GetAll()
	if err != nil {
		return nil, err
	}

	var dataDtos []UserDTO

	for _, value := range data {
		newData := UserDTO{
			ID:          value.ID,
			FirstName:   value.FirstName,
			LastName:    value.LastName,
			UserName:    value.UserName,
			Password:    value.Password,
			BirthDate:   value.BirthDate,
			RoleId:      value.RoleId,
			CreatedDate: value.CreatedAt,
			IsDeleted:   value.IsDeleted,
		}

		dataDtos = append(dataDtos, newData)
	}

	return dataDtos, nil
}

func (u UserRepository) GetAllActiveUsers() ([]UserDTO, error) {
	data, err := model.User{}.GetAll("is_deleted = false")
	if err != nil {
		return nil, err
	}

	var dataDtos []UserDTO

	for _, value := range data {
		newData := UserDTO{
			ID:          value.ID,
			FirstName:   value.FirstName,
			LastName:    value.LastName,
			UserName:    value.UserName,
			Password:    value.Password,
			BirthDate:   value.BirthDate,
			RoleId:      value.RoleId,
			CreatedDate: value.CreatedAt,
			IsDeleted:   value.IsDeleted,
		}

		dataDtos = append(dataDtos, newData)
	}

	return dataDtos, nil
}

func (u UserRepository) GetUserById(id int) (UserDTO, error) {
	data, err := model.User{}.Get(id)
	if err != nil {
		return UserDTO{}, err
	}

	return UserDTO{
		ID:          data.ID,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		UserName:    data.UserName,
		Password:    data.Password,
		BirthDate:   data.BirthDate,
		RoleId:      data.RoleId,
		CreatedDate: data.CreatedAt,
		IsDeleted:   data.IsDeleted,
	}, nil
}

func (u UserRepository) CreateUser(newUserDto UserDTO) (UserDTO, error) {
	result, err := model.User{
		FirstName: newUserDto.FirstName,
		LastName:  newUserDto.LastName,
		UserName:  newUserDto.UserName,
		Password:  newUserDto.Password,
		BirthDate: newUserDto.BirthDate,
		RoleId:    newUserDto.RoleId,
		IsDeleted: newUserDto.IsDeleted}.Add()
	if err != nil {
		return UserDTO{}, err
	}

	return UserDTO{
		ID:          result.ID,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		UserName:    result.UserName,
		Password:    result.Password,
		BirthDate:   result.BirthDate,
		RoleId:      result.RoleId,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted}, nil
}

func (u UserRepository) UpdateUser(updateUserDto UserDTO) (UserDTO, error) {
	data, err := model.User{}.Get(updateUserDto.ID)
	if err != nil {
		return UserDTO{}, err
	}

	result, err := data.Update(model.User{
		FirstName: updateUserDto.FirstName,
		LastName:  updateUserDto.LastName,
		UserName:  updateUserDto.UserName,
		Password:  updateUserDto.Password,
		BirthDate: updateUserDto.BirthDate,
		RoleId:    updateUserDto.RoleId,
		IsDeleted: updateUserDto.IsDeleted,
	})
	if err != nil {
		return UserDTO{}, err
	}

	return UserDTO{
		ID:          result.ID,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		UserName:    result.UserName,
		Password:    result.Password,
		BirthDate:   result.BirthDate,
		RoleId:      result.RoleId,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted}, nil
}

func (u UserRepository) DeleteUser(id int) (UserDTO, error) {
	data, err := model.User{}.Get(id)
	if err != nil {
		return UserDTO{}, err
	}

	result, err := data.Update(model.User{
		IsDeleted: true,
	})
	if err != nil {
		return UserDTO{}, err
	}

	return UserDTO{
		ID:          result.ID,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		UserName:    result.UserName,
		Password:    result.Password,
		BirthDate:   result.BirthDate,
		RoleId:      result.RoleId,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted}, nil
}

package Repository

import (
	model "Go-BarisAkdas-Blog/Domain/Model"
	"time"
)

type IAboutRepository interface {
	GetAllAbouts() ([]AboutDTO, error)
	GetAllActiveAbouts() ([]AboutDTO, error)
	GetAboutById(id int) (AboutDTO, error)
	CreateAbout(about *AboutDTO) (AboutDTO, error)
	UpdateAbout(about *AboutDTO) (AboutDTO, error)
	DeleteAbout(id int) (AboutDTO, error)
}

type AboutRepository struct{}

type AboutDTO struct {
	ID          uint
	Title       string
	Description string
	CreatedDate time.Time
	IsDeleted   bool
}

func NewAboutRepositoory() IAboutRepository {
	return &AboutRepository{}
}

func (a AboutRepository) GetAllAbouts() ([]AboutDTO, error) {
	data, err := model.About{}.GetAll()
	if err != nil {
		return nil, err
	}

	var dataDtos []AboutDTO

	for _, value := range data {
		newData := AboutDTO{
			ID:          value.ID,
			Title:       value.Title,
			Description: value.Description,
			CreatedDate: value.CreatedAt,
			IsDeleted:   value.IsDeleted,
		}
		dataDtos = append(dataDtos, newData)
	}
	return dataDtos, nil
}

func (a AboutRepository) GetAllActiveAbouts() ([]AboutDTO, error) {
	data, err := model.About{}.GetAll("is_deleted = false")
	if err != nil {
		return nil, err
	}

	var dataDtos []AboutDTO

	for _, value := range data {
		newData := AboutDTO{
			ID:          value.ID,
			Title:       value.Title,
			Description: value.Description,
			CreatedDate: value.CreatedAt,
			IsDeleted:   value.IsDeleted,
		}
		dataDtos = append(dataDtos, newData)
	}
	return dataDtos, nil
}

func (a AboutRepository) GetAboutById(id int) (AboutDTO, error) {
	data, err := model.About{}.Get(id)
	if err != nil {
		return AboutDTO{}, err
	}

	newData := AboutDTO{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		CreatedDate: data.CreatedAt,
		IsDeleted:   data.IsDeleted,
	}
	return newData, nil
}

func (a AboutRepository) CreateAbout(about *AboutDTO) (AboutDTO, error) {
	result, err := model.About{
		Title:       about.Title,
		Description: about.Description,
		IsDeleted:   about.IsDeleted,
	}.Add()

	if err != nil {
		return AboutDTO{}, err
	}

	return AboutDTO{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted,
	}, nil
}

func (a AboutRepository) UpdateAbout(about *AboutDTO) (AboutDTO, error) {
	data, err := model.About{}.Get(about.ID)
	if err != nil {
		return AboutDTO{}, err
	}

	result, err := data.Update(model.About{
		Title:       data.Title,
		Description: data.Description,
		IsDeleted:   data.IsDeleted,
	})
	if err != nil {
		return AboutDTO{}, err
	}

	return AboutDTO{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted,
	}, nil
}

func (a AboutRepository) DeleteAbout(id int) (AboutDTO, error) {
	data, err := model.About{}.Get(id)
	if err != nil {
		return AboutDTO{}, err
	}

	result, err := data.Update(model.About{
		IsDeleted: true,
	})
	if err != nil {
		return AboutDTO{}, err
	}

	return AboutDTO{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		CreatedDate: result.CreatedAt,
		IsDeleted:   result.IsDeleted,
	}, nil
}

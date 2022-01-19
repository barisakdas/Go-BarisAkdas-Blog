package Repository

import (
	model "Go-BarisAkdas-Blog/Domain/Model"
	"time"
)

type IMessageRepository interface {
	GetAllMessages() ([]MessageDTO, error)
	GetAllActiveMessages() ([]MessageDTO, error)
	GetMessageById(id int) (MessageDTO, error)
	CreateMessage(newMessageDto MessageDTO) (MessageDTO, error)
	UpdateMessage(updateMessageDto MessageDTO) (MessageDTO, error)
	ReplyMessage(id int, answer string) (MessageDTO, error)
	DeleteMessage(id int) (MessageDTO, error)
}

type MessageRepository struct{}

type MessageDTO struct {
	ID          uint
	Name        string
	Email       string
	Phone       string
	Message     string
	Answer      string
	CreatedDate time.Time
	IsDeleted   bool
	IsReplied   bool
}

func NewMessageRepository() IMessageRepository {
	return &MessageRepository{}
}

func (m MessageRepository) GetAllMessages() ([]MessageDTO, error) {
	data, err := model.Message{}.GetAll()
	if err != nil {
		return nil, err
	}

	var dataDtos []MessageDTO

	for _, value := range data {
		newData := MessageDTO{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Phone:       value.Phone,
			Message:     value.Message,
			Answer:      value.Answer,
			CreatedDate: value.CreatedAt,
			IsReplied:   value.IsReplied,
			IsDeleted:   value.IsDeleted,
		}

		dataDtos = append(dataDtos, newData)
	}

	return dataDtos, nil
}

func (m MessageRepository) GetAllActiveMessages() ([]MessageDTO, error) {
	data, err := model.Message{}.GetAll("is_deleted = false")
	if err != nil {
		return nil, err
	}

	var dataDtos []MessageDTO

	for _, value := range data {
		newData := MessageDTO{
			ID:          value.ID,
			Name:        value.Name,
			Email:       value.Email,
			Phone:       value.Phone,
			Message:     value.Message,
			Answer:      value.Answer,
			CreatedDate: value.CreatedAt,
			IsReplied:   value.IsReplied,
			IsDeleted:   value.IsDeleted,
		}

		dataDtos = append(dataDtos, newData)
	}

	return dataDtos, nil
}

func (m MessageRepository) GetMessageById(id int) (MessageDTO, error) {
	data, err := model.Message{}.Get(id)
	if err != nil {
		return MessageDTO{}, err
	}

	return MessageDTO{
		ID:          data.ID,
		Name:        data.Name,
		Email:       data.Email,
		Phone:       data.Phone,
		Message:     data.Message,
		Answer:      data.Answer,
		CreatedDate: data.CreatedAt,
		IsReplied:   data.IsReplied,
		IsDeleted:   data.IsDeleted,
	}, nil
}

func (m MessageRepository) CreateMessage(newMessageDto MessageDTO) (MessageDTO, error) {
	result, err := model.Message{
		Name:      newMessageDto.Name,
		Email:     newMessageDto.Email,
		Phone:     newMessageDto.Phone,
		Message:   newMessageDto.Message,
		IsReplied: newMessageDto.IsReplied,
		IsDeleted: newMessageDto.IsDeleted,
	}.Add()
	if err != nil {
		return MessageDTO{}, err
	}

	return MessageDTO{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		Phone:       result.Phone,
		Message:     result.Message,
		Answer:      result.Answer,
		CreatedDate: result.CreatedAt,
		IsReplied:   result.IsReplied,
		IsDeleted:   result.IsDeleted,
	}, nil
}

func (m MessageRepository) UpdateMessage(updateMessageDto MessageDTO) (MessageDTO, error) {
	data, err := model.Message{}.Get(updateMessageDto.ID)
	if err != nil {
		return MessageDTO{}, err
	}

	result, err := data.Update(model.Message{
		Name:      updateMessageDto.Name,
		Email:     updateMessageDto.Email,
		Phone:     updateMessageDto.Phone,
		Message:   updateMessageDto.Message,
		IsReplied: updateMessageDto.IsReplied,
		IsDeleted: updateMessageDto.IsDeleted,
	})
	if err != nil {
		return MessageDTO{}, err
	}

	return MessageDTO{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		Phone:       result.Phone,
		Message:     result.Message,
		Answer:      result.Answer,
		CreatedDate: result.CreatedAt,
		IsReplied:   result.IsReplied,
		IsDeleted:   result.IsDeleted,
	}, nil
}

func (m MessageRepository) ReplyMessage(id int, answer string) (MessageDTO, error) {
	data, err := model.Message{}.Get(id)
	if err != nil {
		return MessageDTO{}, err
	}

	result, err := data.Update(model.Message{
		Answer: answer,
	})
	if err != nil {
		return MessageDTO{}, err
	}

	return MessageDTO{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		Phone:       result.Phone,
		Message:     result.Message,
		Answer:      result.Answer,
		CreatedDate: result.CreatedAt,
		IsReplied:   result.IsReplied,
		IsDeleted:   result.IsDeleted,
	}, nil
}

func (m MessageRepository) DeleteMessage(id int) (MessageDTO, error) {
	data, err := model.Message{}.Get(id)
	if err != nil {
		return MessageDTO{}, err
	}

	result, err := data.Update(model.Message{
		IsDeleted: true,
	})
	if err != nil {
		return MessageDTO{}, err
	}

	return MessageDTO{
		ID:          result.ID,
		Name:        result.Name,
		Email:       result.Email,
		Phone:       result.Phone,
		Message:     result.Message,
		Answer:      result.Answer,
		CreatedDate: result.CreatedAt,
		IsReplied:   result.IsReplied,
		IsDeleted:   result.IsDeleted,
	}, nil
}

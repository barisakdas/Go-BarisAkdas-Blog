package Application

import repo "Go-BarisAkdas-Blog/Domain/Repository"

type MessageApplication struct{}

var messageRepo = repo.NewMessageRepository()

func (m MessageApplication) GetAllMessages() ([]repo.MessageDTO, error) {
	return messageRepo.GetAllMessages()
}
func (m MessageApplication) GetAllActiveMessages() ([]repo.MessageDTO, error) {
	return messageRepo.GetAllActiveMessages()
}
func (m MessageApplication) GetMessageById(id int) (repo.MessageDTO, error) {
	return messageRepo.GetMessageById(id)
}
func (m MessageApplication) CreateMessage(newMessageDto repo.MessageDTO) (repo.MessageDTO, error) {
	return messageRepo.CreateMessage(newMessageDto)
}
func (m MessageApplication) UpdateMessage(updateMessageDto repo.MessageDTO) (repo.MessageDTO, error) {
	return messageRepo.UpdateMessage(updateMessageDto)
}
func (m MessageApplication) ReplyMessage(id int, answer string) (repo.MessageDTO, error) {
	return messageRepo.ReplyMessage(id, answer)
}
func (m MessageApplication) DeleteMessage(id int) (repo.MessageDTO, error) {
	return messageRepo.DeleteMessage(id)
}

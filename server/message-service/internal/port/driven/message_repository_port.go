package driven

type MessageRepository interface {
	SaveMessage() error
	GetMessages()
}

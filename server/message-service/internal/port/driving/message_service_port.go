package driving

type MessageServicePort interface {
	SaveMessage() error
	GetMessages()
}

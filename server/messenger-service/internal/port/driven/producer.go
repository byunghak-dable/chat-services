package driven

type Producer[T any] interface {
	Produce(message *T) error
}

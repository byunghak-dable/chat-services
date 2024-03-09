package driving

type Broadcaster[T any] interface {
	Broadcast(messageDto *T) error
}

package application

type user interface {
	GetIdx() uint
	GetName() string
	GetImageUrl() string
	Send() error
}

type workerPool interface {
	RegisterJob(func())
}

type ChatApp struct {
	chatRoom map[int][]user
	pool     workerPool
	poolChan chan error
}

func NewChatApp(pool workerPool) *ChatApp {
	return &ChatApp{
		chatRoom: make(map[int][]user),
		pool:     pool,
		poolChan: make(chan error),
	}
}

func CreateRoom() error {
	return nil
}

func JoinRoom() error {
	return nil
}

func SendMessage() error {
	return nil
}

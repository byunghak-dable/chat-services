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
	pool     workerPool
	poolChan chan error
	chatRoom map[int][]user
}

func NewChatApp(pool workerPool) *ChatApp {
	return &ChatApp{
		pool:     pool,
		poolChan: make(chan error),
		chatRoom: make(map[int][]user),
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

package application

type participant interface {
	GetIdx() uint
	GetName() string
	GetImageUrl() string
	Send() error
}

type ChatApp struct {
	chatRoom map[int][]participant
}

func NewChatApp() *ChatApp {
	return &ChatApp{
		chatRoom: make(map[int][]participant),
	}
}

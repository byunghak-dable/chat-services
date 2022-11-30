package domain

type Room struct {
	Idx          uint
	Name         string
	participants []uint
}

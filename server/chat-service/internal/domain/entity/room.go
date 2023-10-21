package entity

import "log"

type Room struct {
	Name         string
	participants []uint
	Idx          uint
}

func test() {
	log.Print()
}

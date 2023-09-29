package entity

import "log"

type Room struct {
	Idx          uint
	Name         string
	participants []uint
}

func test() {
	log.Print()

}

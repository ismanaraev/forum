package models

type Category int64

const (
	Coding  = 0b000001
	Music   = 0b000010
	Art     = 0b000100
	Sports  = 0b001000
	Cooking = 0b010000
	Other   = 0b100000
)

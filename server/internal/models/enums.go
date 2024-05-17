package models

type (
	Gender int8
	State  int8
)

// iota + 1 because validators consider 0 an default field
const (
	Unknown Gender = iota + 1
	Female
	Male
)

const (
	Lost State = iota + 1
	Found
	Adoption
)

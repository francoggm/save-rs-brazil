package models

type (
	UrgencyType int8
	RescueType  int8
	GenderType  int8
)

// iota + 1 because validators consider 0 with default fields
const (
	Person RescueType = iota + 1
	Animal
)

const (
	Low UrgencyType = iota + 1
	Medium
	Critical
)

const (
	Unknown GenderType = iota + 1
	Male
	Female
)

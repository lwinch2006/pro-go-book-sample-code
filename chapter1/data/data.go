package data

type Rvsp struct {
	Name, Email, Phone string
	WillAttend         bool
}

var Responses = make([]*Rvsp, 0, 10)

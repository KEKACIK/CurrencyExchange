package model

type Currency struct {
	Id       int
	Code     string
	FullName string
	Sign     string
}

func (Currency) TableName() string {
	return "currencies"
}

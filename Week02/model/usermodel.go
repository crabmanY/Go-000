package model

//用户信息
type Person struct {
	UserId   int
	Username string
	Sex      string
	Email    string
}

//住址信息
type Place struct {
	Country string
	City    string
	TelCode int
}

package model

type IDStudent int
type Student struct {
	ID          IDStudent `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Age         int       `json:"age"`
	Gender      int       `json:"gender"`
	Class       int       `json:"class"`
	HomeAddress string    `json:"home_address"`
}

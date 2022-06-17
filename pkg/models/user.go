package models

type User struct {
	Id        int64  `json:"id"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Age       int    `json:"age"`
}

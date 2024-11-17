package model

type Register struct{
	Username 	string		`json:"username" biding:"required"`
	Email		string		`json:"email" binding:"required"`
	Password	string		`json:"password" binding:"required"`
}
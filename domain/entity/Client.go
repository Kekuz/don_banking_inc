package entity

type Client struct {
	ClientId int
	FirstName string
	LastName string
	Accounts []Account
}
package domain

type Client struct {
	ClientId int
	FirstName string
	LastName string
	Accounts []Account
}
package service

type AccountStorage interface{
	Read() [][]string 
}
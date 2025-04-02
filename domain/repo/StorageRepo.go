package domain

type StorageRepo interface{
	Read() [][]string 
}
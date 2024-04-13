package app

import "github.com/Muhless/GO-TOKO/app/models"

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: models.User{}},
		{Model: models.Address{}},
	}
}

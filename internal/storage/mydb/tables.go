package mydb

import agecategories "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/enums/ageCategories"

type hall struct {
	maxCapacity int
	clients     int // TODO: Прописать структуру для хранения клиентов
}

type ingridient struct {
	id    int
	name  string
	count int
}

type dish struct {
	id          int
	name        string
	ingridients []*ingridient
	ageCategory agecategories.AgeCat
}

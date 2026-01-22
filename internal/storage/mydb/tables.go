package mydb

import (
	agecategories "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/enums/ageCategories"
)

type hall struct {
	maxCapacity int
	clients     []*client
}

type ingridient struct {
	id    int
	name  string
	count int
}

type dishIngridients struct {
	dishID       int
	ingridientId int
	count        int
}

type dish struct {
	id          int
	name        string
	ageCategory agecategories.AgeCat
}

type client struct {
	id   string
	name string
	age  int
}

package models

import agecategories "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/enums/ageCategories"

type Client struct {
	Id   string
	Name string
	Age  int
}

type Order struct {
	Id     string
	Client Client
	Dish   string
	Count  int
}

type Ingridient struct {
	Id    int
	Name  string
	Count int
}

type Dish struct {
	Id          int
	Name        string
	Ingridients []Ingridient
	AgeCategory agecategories.AgeCat
}

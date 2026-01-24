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

type Ingridients map[string]int

type DishData struct {
	AgeCategory agecategories.AgeCat
	Ingridients Ingridients
}

type Dishes map[string]DishData

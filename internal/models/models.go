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

type Ingredients map[string]int

type DishData struct {
	AgeCategory agecategories.AgeCat
	Ingredients Ingredients
}

type Dishes map[string]DishData

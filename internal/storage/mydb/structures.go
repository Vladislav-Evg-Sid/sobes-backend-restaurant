package mydb

import (
	agecategories "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/enums/ageCategories"
)

type ingridient map[string]int

type dishData struct {
	ageCategory agecategories.AgeCat
	ingredients ingridient
}

type dish map[string]dishData

type clientData struct {
	name string
	age  int
}

type client map[string]clientData

package mydb

import (
	"fmt"

	agecategories "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/enums/ageCategories"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (db *DataBase) GetClientInfo(clientId string) (models.Client, error) {
	currentClient, exists := db.Clients[clientId]
	if !exists {
		return models.Client{}, fmt.Errorf("No data")
	}
	return models.Client{
		Id:   clientId,
		Name: currentClient.name,
		Age:  currentClient.age,
	}, nil
}

func (db *DataBase) GetDishAgeLimit(dishName string) (agecategories.AgeCat, error) {
	currentDish, exist := db.Dishes[dishName]
	if !exist {
		return "", fmt.Errorf("No data")
	}
	return currentDish.ageCategory, nil
}

func (db *DataBase) GetMaxCapacity() int {
	return db.MaxHallCapacity
}

func (db *DataBase) GetCurrentClientCount() int {
	return len(db.Clients)
}

func (db *DataBase) GetDishByName(dishName string) (models.Dishes, error) {
	currentDish, exist := db.Dishes[dishName]
	if !exist {
		return models.Dishes{}, fmt.Errorf("No data")
	}
	return models.Dishes{
		dishName: models.DishData{
			AgeCategory: currentDish.ageCategory,
			Ingridients: mapDishIngridients(currentDish.ingridients),
		},
	}, nil
}

func mapDishIngridients(ingridients ingridient) models.Ingridients {
	ingr := make(models.Ingridients)
	for key, value := range ingridients {
		ingr[key] = value
	}
	return ingr
}

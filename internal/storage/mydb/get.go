package mydb

import (
	"fmt"

	agecategories "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/enums/ageCategories"
	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
	"github.com/samber/lo"
)

func (db *DataBase) GetClientInfo(clientId string) (models.Client, error) {
	for _, client := range db.Clients {
		if client.id == clientId {
			return models.Client{
				Id:   client.id,
				Name: client.name,
				Age:  client.age,
			}, nil
		}
	}
	return models.Client{}, fmt.Errorf("No data")
}

func (db *DataBase) GetDishAgeLimit(dishName string) (agecategories.AgeCat, error) {
	for _, dish := range db.Dishes {
		if dish.name == dishName {
			return dish.ageCategory, nil
		}
	}
	return "", fmt.Errorf("No data")
}

func (db *DataBase) GetMaxCapacity() int {
	return db.Hall.maxCapacity
}

func (db *DataBase) GetCurrentClientCount() int {
	return len(db.Hall.clients)
}

func (db *DataBase) GetDishList() ([]*models.Dish, error) {
	if len(db.Dishes) == 0 {
		return nil, fmt.Errorf("No Data")
	}
	return db.mapDishList(), nil
}

func (db *DataBase) mapDishList() []*models.Dish {
	return lo.Map(db.Dishes, func(d *dish, _ int) *models.Dish {
		return &models.Dish{
			Id:          d.id,
			Name:        d.name,
			AgeCategory: d.ageCategory,
			Ingridients: db.mapIngridientList(d.id),
		}
	})
}

func (db *DataBase) mapIngridientList(dishId int) []models.Ingridient {
	ingridientsForCurrentDish := make([]models.Ingridient, 0, 5)
	for _, dish := range db.DishIngridients {
		if dish.dishID == dishId {
			ingridientsForCurrentDish = append(
				ingridientsForCurrentDish,
				models.Ingridient{
					Id:    dish.ingridientId,
					Name:  db.GetIngridientNameByID(dish.ingridientId),
					Count: dish.count,
				},
			)
		}
	}
	return ingridientsForCurrentDish
}

func (db *DataBase) GetIngridientNameByID(ingridientId int) string {
	for _, ingridient := range db.Ingridients {
		if ingridient.id == ingridientId {
			return ingridient.name
		}
	}
	return ""
}

package mydb

import (
	"sort"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (db *DataBase) GetCurrentIngredientsCount() models.Ingredients {
	return models.Ingredients(db.Ingredients)
}

func (db *DataBase) GetTop5IngredientsMinCount() [5]models.ProductCount {
	topIngredientsCount := make([]models.ProductCount, 0, len(db.Ingredients))
	for key, value := range db.Ingredients {
		topIngredientsCount = append(
			topIngredientsCount,
			models.ProductCount{
				Name:  key,
				Count: value,
			},
		)
	}

	sort.Slice(topIngredientsCount, func(i, j int) bool {
		return topIngredientsCount[i].Count > topIngredientsCount[j].Count
	})

	var result [5]models.ProductCount
	for i := 0; i < min(5, len(topIngredientsCount)); i++ {
		result[i] = topIngredientsCount[i]
	}

	return result
}

func (db *DataBase) GetTop5IngredientsMaxCount() [5]models.ProductCount {
	topIngredientsCount := make([]models.ProductCount, 0, len(db.Ingredients))
	for key, value := range db.Ingredients {
		topIngredientsCount = append(
			topIngredientsCount,
			models.ProductCount{
				Name:  key,
				Count: value,
			},
		)
	}

	sort.Slice(topIngredientsCount, func(i, j int) bool {
		return topIngredientsCount[i].Count < topIngredientsCount[j].Count
	})

	var result [5]models.ProductCount
	for i := 0; i < min(5, len(topIngredientsCount)); i++ {
		result[i] = topIngredientsCount[i]
	}

	return result
}

func (db *DataBase) GetCountWriteOffIngredients() models.Ingredients {
	return models.Ingredients(db.statistic.writedOffIngredients)
}

func (db *DataBase) GetCountAllIngredients() int {
	countAllIngredients := 0
	for _, count := range db.Ingredients {
		countAllIngredients += count
	}
	return countAllIngredients
}

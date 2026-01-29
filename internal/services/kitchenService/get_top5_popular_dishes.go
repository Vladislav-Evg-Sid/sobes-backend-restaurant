package kitchenservice

import (
	"sort"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (k *KitchenService) GetTop5PopularDishes() [5]models.DishCount {
	topDishCount := make([]models.DishCount, 0, len(k.statistic.topProcessingDishes))
	for key, value := range k.statistic.topProcessingDishes {
		topDishCount = append(
			topDishCount,
			models.DishCount{
				Name:  key,
				Count: value,
			},
		)
	}

	sort.Slice(topDishCount, func(i, j int) bool {
		return topDishCount[i].Count < topDishCount[j].Count
	})

	var result [5]models.DishCount
	for i := 0; i < 5; i++ {
		result[i] = topDishCount[i]
	}

	return result
}

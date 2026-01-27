package waiterservice

import (
	"fmt"

	"github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"
)

func (w *WaiterService) TransferOrderToKitchen(currentOrder models.Order) error {
	if w.ClientAge == -1 {
		return fmt.Errorf("No client")
	}

	dish, err := w.Kitchen.GetDishByName(currentOrder.Dish)
	if err != nil {
		return err
	}
	err = w.Kitchen.ProcessingDish(dish, currentOrder.Count)
	if err != nil {
		return err
	}
	w.ClientAge = -1
	return nil
}

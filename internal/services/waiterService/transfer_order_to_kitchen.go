package waiterservice

import "github.com/Vladislav-Evg-Sid/sobes-backend-restaurant/internal/models"

func (w *WaiterService) TransferOrderToKitchen(currentOrder models.Order) error {
	dish, err := w.Kitchen.GetDishByName(currentOrder.Dish)
	if err != nil {
		return err
	}
	return w.Kitchen.ProcessingDish(dish)
}

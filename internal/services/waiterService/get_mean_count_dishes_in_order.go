package waiterservice

func (w *WaiterService) GetMeanCountDishesInOrder() float32 {
	return float32(w.statistic.countServedDishes) / float32(w.statistic.countServedOrders)
}

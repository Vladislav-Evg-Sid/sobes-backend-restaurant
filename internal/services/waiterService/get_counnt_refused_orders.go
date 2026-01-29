package waiterservice

func (w *WaiterService) GetCountRefusedOrders() int {
	return w.statistic.countRefusedOrders
}

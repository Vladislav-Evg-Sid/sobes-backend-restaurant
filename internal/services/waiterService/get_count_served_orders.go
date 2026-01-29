package waiterservice

func (w *WaiterService) GetCountServedOrders() int {
	return w.statistic.countServedOrders
}

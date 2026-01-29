package waiterservice

func (w *WaiterService) GetCountMore18Dishes() int {
	return w.statistic.countMore18Dishes
}

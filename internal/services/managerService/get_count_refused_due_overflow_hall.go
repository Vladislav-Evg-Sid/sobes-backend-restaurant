package managerservice

func (m *ManagerService) GetCountRefusedDueOverflowHall() int {
	return m.statistic.countRefusedDueOverflowHall
}

package managerservice

func (m *ManagerService) GetMaxHallLoad() float32 {
	return m.statistic.maxHallLoad
}

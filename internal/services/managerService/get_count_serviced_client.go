package managerservice

func (m *ManagerService) GetCountServedClients() int {
	return m.statistic.countServedClients
}

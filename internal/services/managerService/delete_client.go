package managerservice

func (m *ManagerService) DeleteClient(clientId string) error {
	err := m.Waiter.DeleteClientAge()
	if err != nil {
		return err
	}
	return m.Storage.DeleteClient(clientId)
}

package managerserviceapi

func (m *ManagerServiceAPI) DeleteClients(clientsId []string) error {
	for _, cliID := range clientsId {
		err := m.Service.DeleteClient(cliID)
		if err != nil {
			return err
		}
	}
	return nil
}

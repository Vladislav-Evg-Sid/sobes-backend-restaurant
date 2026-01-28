package waiterservice

func (w *WaiterService) SetClientAge(clientAge int) error {
	w.ClientAge = clientAge
	return nil
}

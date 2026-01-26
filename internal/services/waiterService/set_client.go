package waiterservice

func (w *WaiterService) DeleteClientAge() error {
	w.ClientAge = -1
	return nil
}

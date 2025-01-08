package helper

type ResponseBooking struct{}

func (ResponseBooking) CreateSuccessfully() string {
	return "Create data booking successfully"
}
func (ResponseBooking) CreateFailed() string {
	return "Create data booking failed"
}
func (ResponseBooking) GetSuccessfully() string {
	return "Fetch data booking successfully"
}
func (ResponseBooking) GetFailed() string {
	return "Fetch data booking failed"
}
func (ResponseBooking) UpdateSuccessfully() string {
	return "Update data booking successfully"
}
func (ResponseBooking) UpdateFailed() string {
	return "Update data booking failed"
}
func (ResponseBooking) IdFailed() string {
	return "Invalid booking ID"
}
func (ResponseBooking) DeleteSuccessfully() string {
	return "Delete booking successfully"
}
func (ResponseBooking) DeleteFailed() string {
	return "Delete booking failed"
}
func (ResponseBooking) RequestFailed() string {
	return "Invalid request payload"
}

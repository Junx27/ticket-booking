package helper

type ResponsePayment struct{}

func (ResponsePayment) CreateSuccessfully() string {
	return "Create data payment successfully"
}
func (ResponsePayment) CreateFailed() string {
	return "Create data payment failed"
}
func (ResponsePayment) GetSuccessfully() string {
	return "Fetch data payment successfully"
}
func (ResponsePayment) GetFailed() string {
	return "Fetch data payment failed"
}
func (ResponsePayment) UpdateSuccessfully() string {
	return "Update data payment successfully"
}
func (ResponsePayment) UpdateFailed() string {
	return "Update data payment failed"
}
func (ResponsePayment) IdFailed() string {
	return "Invalid payment ID"
}
func (ResponsePayment) DeleteSuccessfully() string {
	return "Delete payment successfully"
}
func (ResponsePayment) DeleteFailed() string {
	return "Delete payment failed"
}
func (ResponsePayment) RequestFailed() string {
	return "Invalid request payload"
}

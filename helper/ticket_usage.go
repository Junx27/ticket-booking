package helper

type ResponseTicketUsage struct{}

func (ResponseTicketUsage) CreateSuccessfully() string {
	return "Create data ticket usage successfully"
}
func (ResponseTicketUsage) CreateFailed() string {
	return "Create data ticket usage failed"
}
func (ResponseTicketUsage) GetSuccessfully() string {
	return "Fetch data ticket usage successfully"
}
func (ResponseTicketUsage) GetFailed() string {
	return "Fetch data ticket usage failed"
}
func (ResponseTicketUsage) UpdateSuccessfully() string {
	return "Update data ticket usage successfully"
}
func (ResponseTicketUsage) UpdateFailed() string {
	return "Update data ticket usage failed"
}
func (ResponseTicketUsage) IdFailed() string {
	return "Invalid ticket usage ID"
}
func (ResponseTicketUsage) DeleteSuccessfully() string {
	return "Delete ticket usage successfully"
}
func (ResponseTicketUsage) DeleteFailed() string {
	return "Delete ticket usage failed"
}
func (ResponseTicketUsage) RequestFailed() string {
	return "Invalid request payload"
}

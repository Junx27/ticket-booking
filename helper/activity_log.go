package helper

type ResponseActivityLog struct{}

func (ResponseActivityLog) CreateSuccessfully() string {
	return "Create data activity log successfully"
}
func (ResponseActivityLog) CreateFailed() string {
	return "Create data activity log failed"
}
func (ResponseActivityLog) GetSuccessfully() string {
	return "Fetch data activity log successfully"
}
func (ResponseActivityLog) GetFailed() string {
	return "Fetch data activity log failed"
}
func (ResponseActivityLog) UpdateSuccessfully() string {
	return "Update data activity log successfully"
}
func (ResponseActivityLog) UpdateFailed() string {
	return "Update data activity log failed"
}
func (ResponseActivityLog) IdFailed() string {
	return "Invalid activity log ID"
}
func (ResponseActivityLog) DeleteSuccessfully() string {
	return "Delete activity log successfully"
}
func (ResponseActivityLog) DeleteFailed() string {
	return "Delete activity log failed"
}
func (ResponseActivityLog) RequestFailed() string {
	return "Invalid request payload"
}

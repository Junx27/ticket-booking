package helper

type ResponseNotification struct{}

func (ResponseNotification) CreateSuccessfully() string {
	return "Create data notification successfully"
}
func (ResponseNotification) InvalidCreate() string {
	return "Create data notification failed"
}
func (ResponseNotification) GetSuccessfully() string {
	return "Fetch data notification successfully"
}
func (ResponseNotification) InvalidGet() string {
	return "Fetch data notification failed"
}
func (ResponseNotification) UpdateSuccessfully() string {
	return "Update data notification successfully"
}
func (ResponseNotification) Invalidupdate() string {
	return "Update data notification failed"
}
func (ResponseNotification) InvalidId() string {
	return "Invalid notification ID"
}
func (ResponseNotification) DeleteSuccessfully() string {
	return "Delete notification successfully"
}
func (ResponseNotification) InvalidDelete() string {
	return "Delete notification failed"
}
func (ResponseNotification) InvalidRequest() string {
	return "Invalid request payload"
}

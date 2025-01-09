package helper

import "fmt"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(message string, data any) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func FailedResponse(message string) Response {
	return Response{
		Message: message,
	}
}

type ResponseMessage struct{}

func (ResponseMessage) CreateSuccessfully(name string) string {
	return fmt.Sprintf("Create data %s successfully", name)
}

func (ResponseMessage) CreateFailed(name string) string {
	return fmt.Sprintf("Create data %s failed", name)
}

func (ResponseMessage) GetSuccessfully(name string) string {
	return fmt.Sprintf("Fetch data %s successfully", name)
}

func (ResponseMessage) GetFailed(name string) string {
	return fmt.Sprintf("Fetch data %s failed", name)
}

func (ResponseMessage) UpdateSuccessfully(name string) string {
	return fmt.Sprintf("Update data %s successfully", name)
}

func (ResponseMessage) UpdateFailed(name string) string {
	return fmt.Sprintf("Update data %s failed", name)
}

func (ResponseMessage) IdFailed(id string) string {
	return fmt.Sprintf("Invalid %s ID", id)
}

func (ResponseMessage) DeleteSuccessfully(name string) string {
	return fmt.Sprintf("Delete %s successfully", name)
}

func (ResponseMessage) DeleteFailed(name string) string {
	return fmt.Sprintf("Delete %s failed", name)
}

func (ResponseMessage) RequestFailed(name string) string {
	return fmt.Sprintf("Invalid %s request payload", name)
}
func (ResponseMessage) NotFound(name string) string {
	return fmt.Sprintf("Fetch %s not found", name)
}

package helper

import (
	"fmt"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Token   string      `json:"token,omitempty"`
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

func AuthResponse(message, token string) Response {
	return Response{
		Success: true,
		Message: message,
		Token:   token,
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
func (ResponseMessage) DeleteAllSuccessfully(name string) string {
	return fmt.Sprintf("Delete all %s successfully", name)
}

func (ResponseMessage) DeleteAllFailed(name string) string {
	return fmt.Sprintf("Delete all %s failed", name)
}

func (ResponseMessage) RequestFailed(name string) string {
	return fmt.Sprintf("Invalid %s request payload", name)
}
func (ResponseMessage) NotFound(name string) string {
	return fmt.Sprintf("Fetch %s not found", name)
}

func (ResponseMessage) LoginFailed() Response {
	return Response{
		Success: false,
		Message: "Login failed please check email and password!",
	}
}
func (ResponseMessage) LoginFailedEntity() Response {
	return Response{
		Success: false,
		Message: "Email and password is not valid!",
	}
}

func (ResponseMessage) LoginSuccessfully() string {
	return "Login successfully"
}
func (ResponseMessage) RegisterFailed() Response {
	return Response{
		Success: false,
		Message: "Please check email is valid and password minimum 8 character must contain at least one uppercase letter, one number, and one symbol",
	}
}
func (ResponseMessage) RegisterFailedEntity() Response {
	return Response{
		Success: false,
		Message: "Email and password is not valid!",
	}
}

func (ResponseMessage) RegisterSuccessfully() string {
	return "Register successfully"
}

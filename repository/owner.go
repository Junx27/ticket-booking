package repository

type HasUserID interface {
	GetUserID(id uint) (uint, error)
}

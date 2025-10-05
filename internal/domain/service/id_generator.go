package service

type IDGenerator interface {
	GenerateID() (int64, error)
}

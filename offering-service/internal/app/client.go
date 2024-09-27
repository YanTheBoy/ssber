package app

type FlatsService struct {
	storage FlatsStorage
}

func NewFlatsService(s FlatsStorage) *FlatsService {
	return &FlatsService{
		storage: s,
	}
}

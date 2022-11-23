package service

// Service ...
type Service struct{}

// New function
func New() (*Service, error) {
	return &Service{}, nil
}

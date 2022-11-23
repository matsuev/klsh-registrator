package service

// Service ...
type Service struct{}

// NewService function
func NewService() (*Service, error) {
	return &Service{}, nil
}

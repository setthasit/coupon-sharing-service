package services

type HealthCheckService interface {
	IsOk() error
}

type HealthCheckServiceInstance struct{}

func NewHealthCheckService() HealthCheckService {
	return &HealthCheckServiceInstance{}
}

func (*HealthCheckServiceInstance) IsOk() error {
	return nil
}

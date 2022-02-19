package domains

type HealthCheckService interface {
	IsOk() error
}

type HealthCheckServiceInstance struct{}

func NewHealthCheckService() *HealthCheckServiceInstance {
	return &HealthCheckServiceInstance{}
}

func (*HealthCheckServiceInstance) IsOk() error {
	return nil
}

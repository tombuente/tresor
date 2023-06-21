package health

import (
	"context"

	"github.com/tombuente/tresor/spec/healthspec"
)

var _ Service = (*serviceImpl)(nil)

type Service interface {
	GetHealth(ctx context.Context) healthspec.HealthRes
}

type serviceImpl struct{}

func NewService() Service {
	return serviceImpl{}
}

func (service serviceImpl) GetHealth(ctx context.Context) healthspec.HealthRes {
	return healthspec.HealthRes{Message: "API is healthy."}
}

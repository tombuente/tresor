package health

import (
	"context"

	"github.com/tombuente/tresor/service/base"
)

type Service struct{}

func NewService() Service {
	return Service{}
}

func (service Service) GetHealth(ctx context.Context) base.SimpleMessage {
	return base.SimpleMessage{Message: "API is healthy."}
}

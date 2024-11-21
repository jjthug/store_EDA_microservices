package baskets

import (
	"EDA_GO/internal/system"
	"context"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono system.Service) (err error) {
	return nil
}

func Root(ctx context.Context, svc system.Service) (err error) {
	return nil
}

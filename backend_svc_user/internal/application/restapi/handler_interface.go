package restapi

import "context"

type HandlerInterface interface {
	Run(ctx context.Context) error
	GetContext() *context.Context
}

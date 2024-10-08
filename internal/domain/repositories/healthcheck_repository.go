package repositories

import (
	"context"
)

type IHealthCheckRepository interface {
	CheckDbConnection(ctx context.Context) error
}

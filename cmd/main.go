package main

import (
	"context"

	"github.com/dmarins/student-api/internal/infrastructure/di"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/dmarins/student-api/internal/infrastructure/logger"
)

func main() {
	ctx := context.Background()

	env.LoadEnvironmentVariables(ctx)
	logger.NewLogger()

	di.StartCompositionRoot(ctx).Run()
}

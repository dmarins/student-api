package main

import (
	"context"

	"github.com/dmarins/student-api/internal/infrastructure/di"
	"github.com/dmarins/student-api/internal/infrastructure/env"
)

func main() {
	env.LoadEnvironmentVariables()
	di.StartCompositionRoot(context.Background()).Run()
}

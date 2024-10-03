package tests

import (
	"context"
	"testing"

	"github.com/dmarins/student-api/internal/domain/dtos"
	"github.com/dmarins/student-api/internal/domain/mocks"
	"github.com/dmarins/student-api/internal/infrastructure/env"
	"go.uber.org/mock/gomock"
)

type contextKey string

var requestContextEnvVarName = contextKey(env.ProvideRequestContextName())

func SetupTest(t *testing.T) (context.Context, *gomock.Controller, *mocks.MockITracer, *mocks.MockILogger) {
	ctrl := gomock.NewController(t)

	tracerMock := mocks.NewMockITracer(ctrl)
	loggerMock := mocks.NewMockILogger(ctrl)

	requestContext := dtos.RequestContext{
		TenantId: "x-tenant",
		Cid:      "cid",
	}

	ctx := context.WithValue(context.Background(), requestContextEnvVarName, requestContext)

	return ctx, ctrl, tracerMock, loggerMock
}

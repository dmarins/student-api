package repositories_test

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckRepository_CheckDbConnection_Success(t *testing.T) {
	sut := builder.BuildHealthCheckRepository()

	err := sut.CheckDbConnection(builder.Ctx)

	assert.NoError(t, err)
}

func TestHealthCheckRepository_CheckDbConnection_Fails(t *testing.T) {
	sut := failedBuilder.BuildHealthCheckRepository()

	err := sut.CheckDbConnection(builder.Ctx)

	assert.Error(t, err)
}

package tracer

const (
	StudentHandlerCreate  = "StudentHandler:Create"
	HealthCheckHandlerGet = "HealthCheckHandler:Get"

	StudentCreationUseCasePersistenceExecute = "StudentCreationUseCasePersistence:Execute"
	StudentCreationUseCaseValidationsExecute = "StudentCreationUseCaseValidations:Execute"
	HealthCheckExecute                       = "HealthCheck:Execute"

	StudentRepositoryAdd                   = "StudentRepository:Add"
	StudentRepositoryExistsByName          = "StudentRepository:ExistsByName"
	HealthCheckRepositoryCheckDbConnection = "HealthCheckRepository:CheckDbConnection"
)

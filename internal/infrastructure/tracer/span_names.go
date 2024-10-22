package tracer

const (
	HealthCheckHandlerGet = "HealthCheckHandler:Get"
	StudentHandlerCreate  = "StudentHandler:Create"
	StudentHandlerGet     = "StudentHandler:Get"

	HealthCheckExecute                       = "HealthCheck:Execute"
	StudentCreationUseCasePersistenceExecute = "StudentCreationUseCasePersistence:Execute"
	StudentCreationUseCaseValidationsExecute = "StudentCreationUseCaseValidations:Execute"
	StudentReadingUseCaseFindByIdExecute     = "StudentReadingUseCaseFindById:Execute"

	HealthCheckRepositoryCheckDbConnection = "HealthCheckRepository:CheckDbConnection"
	StudentRepositoryAdd                   = "StudentRepository:Add"
	StudentRepositoryExistsByName          = "StudentRepository:ExistsByName"
	StudentRepositoryFindById              = "StudentRepository:FindById"
)

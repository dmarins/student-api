package tracer

const (
	HealthCheckHandlerGet = "HealthCheckHandler:Get"
	StudentHandlerCreate  = "StudentHandler:Create"
	StudentHandlerGet     = "StudentHandler:Get"

	HealthCheckExecute                     = "HealthCheck:Execute"
	StudentCreateUseCasePersistenceExecute = "StudentCreateUseCasePersistence:Execute"
	StudentCreateUseCaseValidationsExecute = "StudentCreateUseCaseValidations:Execute"
	StudentReadUseCaseFindByIdExecute      = "StudentReadUseCaseFindById:Execute"
	StudentUpdateUseCasePersistenceExecute = "StudentUpdateUseCasePersistence:Execute"
	StudentUpdateUseCaseValidationsExecute = "StudentUpdateUseCaseValidations:Execute"

	HealthCheckRepositoryCheckDbConnection = "HealthCheckRepository:CheckDbConnection"
	StudentRepositoryAdd                   = "StudentRepository:Add"
	StudentRepositoryExistsByName          = "StudentRepository:ExistsByName"
	StudentRepositoryFindById              = "StudentRepository:FindById"
	StudentRepositoryUpdate                = "StudentRepository:Update"
	StudentRepositoryDelete                = "StudentRepository:Delete"
)

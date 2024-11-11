package tracer

const (
	HealthCheckHandlerGet = "HealthCheckHandler:Get"
	StudentHandlerPost    = "StudentHandler:Post"
	StudentHandlerGet     = "StudentHandler:Get"
	StudentHandlerPut     = "StudentHandler:Put"
	StudentHandlerDelete  = "StudentHandler:Delete"

	HealthCheckExecute                     = "HealthCheck:Execute"
	StudentCreateUseCasePersistenceExecute = "StudentCreateUseCasePersistence:Execute"
	StudentCreateUseCaseValidationsExecute = "StudentCreateUseCaseValidations:Execute"
	StudentReadUseCaseFindByIdExecute      = "StudentReadUseCaseFindById:Execute"
	StudentUpdateUseCasePersistenceExecute = "StudentUpdateUseCasePersistence:Execute"
	StudentUpdateUseCaseValidationsExecute = "StudentUpdateUseCaseValidations:Execute"
	StudentDeleteUseCaseFindByIdExecute    = "StudentDeleteUseCaseFindById:Execute"
	StudentDeleteUseCasePersistenceExecute = "StudentDeleteUseCasePersistence:Execute"

	HealthCheckRepositoryCheckDbConnection = "HealthCheckRepository:CheckDbConnection"
	StudentRepositoryAdd                   = "StudentRepository:Add"
	StudentRepositoryExistsByName          = "StudentRepository:ExistsByName"
	StudentRepositoryFindById              = "StudentRepository:FindById"
	StudentRepositoryUpdate                = "StudentRepository:Update"
	StudentRepositoryDelete                = "StudentRepository:Delete"
	StudentRepositorySearchBy              = "StudentRepository:SearchBy"
)

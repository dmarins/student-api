package tracer

const (
	HealthCheckHandlerGet = "HealthCheckHandler:Get"
	StudentHandlerCreate  = "StudentHandler:Create"
	StudentHandlerRead    = "StudentHandler:Read"
	StudentHandlerUpdate  = "StudentHandler:Update"
	StudentHandlerDelete  = "StudentHandler:Delete"
	StudentHandlerSearch  = "StudentHandler:Search"

	HealthCheckExecute                     = "HealthCheck:Execute"
	StudentCreateUseCasePersistenceExecute = "StudentCreateUseCasePersistence:Execute"
	StudentCreateUseCaseValidationsExecute = "StudentCreateUseCaseValidations:Execute"
	StudentReadUseCaseFindByIdExecute      = "StudentReadUseCaseFindById:Execute"
	StudentUpdateUseCasePersistenceExecute = "StudentUpdateUseCasePersistence:Execute"
	StudentUpdateUseCaseValidationsExecute = "StudentUpdateUseCaseValidations:Execute"
	StudentDeleteUseCaseFindByIdExecute    = "StudentDeleteUseCaseFindById:Execute"
	StudentDeleteUseCasePersistenceExecute = "StudentDeleteUseCasePersistence:Execute"
	StudentSearchUseCaseSearchByExecute    = "StudentSearchUseCaseSearchBy:Execute"

	HealthCheckRepositoryCheckDbConnection = "HealthCheckRepository:CheckDbConnection"
	StudentRepositoryAdd                   = "StudentRepository:Add"
	StudentRepositoryExistsByName          = "StudentRepository:ExistsByName"
	StudentRepositoryFindById              = "StudentRepository:FindById"
	StudentRepositoryUpdate                = "StudentRepository:Update"
	StudentRepositoryDelete                = "StudentRepository:Delete"
	StudentRepositorySearchBy              = "StudentRepository:SearchBy"
	StudentRepositoryCount                 = "StudentRepository:Count"
)

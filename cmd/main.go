package main

import (
	"github.com/dmarins/student-api/internal/infrastructure/di"
	"github.com/dmarins/student-api/internal/infrastructure/env"

	_ "github.com/dmarins/student-api/docs/openapi"
)

//	@title			Student Swagger Example API
//	@version		1.0
//	@description	This is an example API written in Go.
///	@termsOfService	http://swagger.io/terms/

///	@contact.name	API Support
///	@contact.url	http://www.swagger.io/support
///	@contact.email	support@swagger.io

///	@license.name	Apache 2.0
///	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/
// @schemes	http
func main() {
	env.LoadEnvironmentVariables()
	di.StartCompositionRoot().Run()
}

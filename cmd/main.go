package main

import (
	"paramount_school/cmd/server"
	_ "paramount_school/docs"
	"paramount_school/internal/repository"
)

// @title           Lunch Wallet Swagger API
// @version         1.0
// @description     This is a lunch wallet server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Lunch-wallet Team API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  info@lunchwallet.com

// @license.name  BSD
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
// @in header
// @name Authorization

func main() {
	//Gets the environment variables
	env := server.InitDBParams()

	//Initializes the database
	db, err := repository.Initialize(env.DbUrl)
	if err != nil {
		return
	}

	//Runs the app
	server.Run(db, env.Port)
}

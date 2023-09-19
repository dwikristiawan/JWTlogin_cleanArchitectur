package cmd

import (
	"fmt"
	"jwtLogin/app/config"
	"jwtLogin/internal/rest/user"
	service "jwtLogin/service/user"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var (
	EnvFilePath string
	rootCmd     = &cobra.Command{
		Use:   "cobra-cli",
		Short: "JWT_Login",
	}
)
var (
	rootConfig  config.Root
	database    *sqlx.DB
	userHandler user.UserHandler
)

// Execute executes the root command.
func Execute() {
	rootCmd.PersistentFlags().StringVarP(&EnvFilePath, "env", "e", ".env", ".env file to read from")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Cannot Run CLI. err > ", err)
		os.Exit(1)
	}
}

func initPostGresDB() {
	var err error
	log.Infof("Initialize Postgres Connection")
	database, err = config.OpenPostgresDatabaseConnection(config.Postgres{
		Host:     rootConfig.Postgres.Host,
		Port:     rootConfig.Postgres.Port,
		User:     rootConfig.Postgres.User,
		Password: rootConfig.Postgres.Password,
		Dbname:   rootConfig.Postgres.Dbname,
	})

	if err != nil {
		log.Errorf("FAILED CONNECT POSTGRES >>>> %v", err)
		os.Exit(1)
	}
}
func initConfigReader() {
	rootConfig = config.Load(EnvFilePath)
}

func initApp() {
	userHandler = *user.NewUserHandler(*service.NewUserService())
}

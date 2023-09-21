package cmd

import (
	"JWTLogin/app/config"
	"JWTLogin/internal/rest/user"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use:   "rest",
	Short: "Start REST server",
	Run:   restServer,
}

func init() {
	rootCmd.AddCommand(restCommand)
}
func restServer(cmd *cobra.Command, args []string) {
	props := config.LoadForServer(EnvFilePath)
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{""},
		AllowHeaders: []string{""},
	}))
	user.InitUserHandler(e, rootConfig, userHandler)

	err := e.Start(props.Address)
	if err != nil {
		log.Errorf("Cannot Start the application !!, Err > ", err)
		os.Exit(1)
	}
}

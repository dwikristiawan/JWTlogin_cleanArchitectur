package cmd

import "github.com/spf13/cobra"

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

	healthcheckGroup := e.Group("/healthcheck")
}
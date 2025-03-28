package commands

import (
	"fmt"

	"github.com/JscorpTech/jst-go/api/routes"
	"github.com/JscorpTech/jst-go/bootstrap"
	"github.com/spf13/cobra"
)

func RunServer() *cobra.Command {
	return &cobra.Command{
		Use: "runserver",
		Run: func(cmd *cobra.Command, args []string) {
			app := bootstrap.NewApp()
			defer bootstrap.CloseApp(app)
			routes.InitRoutes(app)
			app.Server.Logger.Fatal(app.Server.Start(fmt.Sprintf(":%s", app.Env.Port)))
		},
	}
}

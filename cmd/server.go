package cmd

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Runs the Server Component of the application",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, _ := os.Create("gin.log")
		gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

		router := gin.Default()

		router.Use(gin.Recovery())
		router.Use(cors.Default())

		router.Use(static.Serve("/", static.LocalFile("./views", true)))

		// Setup route group for the API
		api := router.Group("/api")
		{
			api.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})

			api.GET("/uptime", UptimeHandler)
			api.GET("/starttime", StartTimeHandler)
			api.GET("/ip", ClientIPHandler)
			api.GET("/uniqueip", IPAddressHandler)

			// Ok now lets get some drug api's set up
			apiDrug := api.Group("/drug")

			apiDrug.GET("/", RootDrugHandler)
			apiDrug.GET("/:drug", DrugHandler)
			apiDrug.GET("/:drug/classifications", DrugClassificationsHandler)
		}

		// router.Run(":80")
		router.Run(":8112")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

/*h
Copyright © 2022 Johannes Siebel <johannes.siebel@protonmail.com>

*/
package cmd

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"

	"github.com/therealjsie/resonance-backend/internals"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run http server on port 8000 running the resonance application.",
	Long:  `Run http server on port 8000 running the resonance application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Database
		// TODO: Run database migration in Kubernetes job
		// https://andrewlock.net/deploying-asp-net-core-applications-to-kubernetes-part-8-running-database-migrations-using-jobs-and-init-containers/
		// TODO: Health check
		// TODO: Integrate OIDC (Dex, Keycloak)

		router := gin.Default()
		router.Use(cors.Default())

		router.GET("/healthz", health_endpoint)

		v1 := router.Group("/api/v1")
		{
			v1.GET("/posts", posts_endpoint)
			v1.GET("/config", config_endpoint)
		}

		router.Run(":8000")
	},
}

func posts_endpoint(c *gin.Context) {
	config := internals.ReadConfig()
	posts, err := internals.Get_Posts(config)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"error": err,
		})
	}

	c.JSON(200, posts)
}

func config_endpoint(c *gin.Context) {
	config := internals.ReadConfig()
	c.JSON(200, config)
}

func health_endpoint(c *gin.Context) {
	c.JSON(200, "Service is up.")
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

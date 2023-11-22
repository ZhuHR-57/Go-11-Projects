/**
  @Go version: 1.17.6
  @project: elevenProject
  @ide: GoLand
  @file: routes.go
  @author: Lido
  @time: 2023-01-11 19:10
  @description:
*/
package routes

import (
	"forumProject/logger"
	"forumProject/settings"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, settings.Conf.Version)
	})

	return r
}

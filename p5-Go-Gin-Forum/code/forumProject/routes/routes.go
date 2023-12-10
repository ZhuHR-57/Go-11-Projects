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
	"forumProject/controller"
	"forumProject/logger"
	snowflake "forumProject/pkg/sonwflake"
	"forumProject/settings"
	"net/http"
	"strconv"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, settings.Conf.Version)
	})

	r.GET("/sf", func(c *gin.Context) {
		ID, err := snowflake.GetID()
		if err != nil {
			zap.L().Fatal("sf getID failed")
		}
		c.String(http.StatusOK, strconv.FormatUint(ID, 10))
	})

	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	return r
}

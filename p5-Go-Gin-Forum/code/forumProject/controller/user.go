package controller

import (
	"forumProject/logic"
	"forumProject/models"
	"net/http"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func SignUpHandler(c *gin.Context) {

	// 1. 获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))

		// 判断err类型是否是validator内置的类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, gin.H{
				"msg": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": errs.Translate(trans), // 使用翻译器
		})
		return
	}

	// 2. 业务逻辑
	logic.SignUp()

	// 3. 返回值
	c.JSON(http.StatusOK, gin.H{
		"msg": "successed",
	})
}

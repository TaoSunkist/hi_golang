package user

import (
	// "encoding/json"
	"hi_golang/service/thsunkistsrv/model"

	"github.com/gin-gonic/gin"
)

// 详细信息
func Info(c *gin.Context) {
	var userInfo = model.Fake()
	var apiResponse = new(model.ApiResponse)
	apiResponse.Data = userInfo
	c.JSON(200, apiResponse)
}

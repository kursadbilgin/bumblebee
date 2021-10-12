package Controllers

import (
	"bumblebee/Lib"
	"bumblebee/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

func ConverterDeepLinkToLink(c *gin.Context) {
	var request Models.ConverterDeepLinkToWebRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}

	if validation := validation(request.DeepLink, "deepLink"); validation != nil {
		c.JSON(400, &validation)
		return
	}
	cr := c.MustGet("redis").(*redis.Client)
	webLink, err := Lib.Get(cr, request.DeepLink)
	if err != nil {
		logError := fmt.Sprintf("[ConverterDeepLinkToLink] Lib.Get error %s", err.Error())
		LogError(logError, err.Error())
		c.JSON(400, Models.ErrorResponse{
			Status:       http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("[ConverterDeepLinkToLink] Lib.Get error %s", err.Error()),
		})
		return
	}
	if webLink == "no value found" {
		webLink, err = DeepToWeb(request.DeepLink)
		if err != nil {
			c.JSON(400, Models.ErrorResponse{
				Status:       http.StatusBadRequest,
				ErrorMessage: err.Error(),
			})
			return
		}
	}
	if err = Lib.Set(cr, request.DeepLink, webLink); err != nil {
		logError := fmt.Sprintf("[ConverterDeepLinkToLink] Lib.Set error %s", err.Error())
		LogError(logError, err.Error())
		c.JSON(400, Models.ErrorResponse{
			Status:       http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("[ConverterDeepLinkToLink] Lib.Set error %s", err.Error()),
		})
		return
	}

	c.JSON(200, Models.ResponseDeepToWeb{
		Status:  http.StatusOK,
		WebLink: webLink,
	})
}

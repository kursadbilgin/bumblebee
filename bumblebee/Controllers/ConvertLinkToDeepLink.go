package Controllers

import (
	"bumblebee/Lib"
	"bumblebee/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
)

func ConverterLinkToDeepLink(c *gin.Context) {
	var request Models.ConverterWebToDeepLinkRequest
	if err := c.BindJSON(&request); err != nil {
		return
	}

	if validation := validation(request.WebURL, "webLink"); validation != nil {
		c.JSON(400, &validation)
		return
	}
	cr := c.MustGet("redis").(*redis.Client)
	deepLink, err := Lib.Get(cr, request.WebURL)
	if err != nil {
		logError := fmt.Sprintf("[ConverterLinkToDeepLink] Lib.Get error %s", err.Error())
		LogError(logError, err.Error())
		c.JSON(400, Models.ErrorResponse{
			Status:       http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("[ConverterLinkToDeepLink] Lib.Get error %s", err.Error()),
		})
		return
	}
	if deepLink == "no value found" {
		deepLink, err = WebToDeep(request.WebURL)
		if err != nil {
			c.JSON(400, Models.ErrorResponse{
				Status:       http.StatusBadRequest,
				ErrorMessage: err.Error(),
			})
			return
		}
	}
	if err = Lib.Set(cr, request.WebURL, deepLink); err != nil {
		logError := fmt.Sprintf("[ConverterLinkToDeepLink] Lib.Set error %s", err.Error())
		LogError(logError, err.Error())
		c.JSON(400, Models.ErrorResponse{
			Status:       http.StatusBadRequest,
			ErrorMessage: fmt.Sprintf("[ConverterLinkToDeepLink] Lib.Set error %s", err.Error()),
		})
		return
	}

	c.JSON(200, Models.ResponseWebToDeep{
		Status:   http.StatusOK,
		DeepLink: deepLink,
	})
}

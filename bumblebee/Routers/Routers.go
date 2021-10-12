package Routers

import (
	"bumblebee/Controllers"
	"bumblebee/Lib"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func SetupRouter(redis *redis.Client) *gin.Engine {
	r := gin.Default()
	r.Use(Lib.RedisSet(redis))
	v1 := r.Group("v1/api")
	{
		v1.POST("convert-link-to-deeplink", Controllers.ConverterLinkToDeepLink)
		v1.POST("convert-deeplink-to-link", Controllers.ConverterDeepLinkToLink)
	}

	return r
}

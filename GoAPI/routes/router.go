package routers

import (
	"io"
	"mind-set/config"
	"mind-set/internal/middleware"
	"mind-set/internal/utils/errors"
	response "mind-set/internal/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRouters() *gin.Engine {
	var r *gin.Engine

	if config.Config.Debug == false {
		r = ReleaseRouter()
		r.Static("/uploads", "./uploads")
		r.Use(
			middleware.RequestCostHandler(),
			middleware.CustomLogger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	} else {
		r = gin.New()
		r.Static("/uploads", "./uploads")
		r.Use(
			middleware.RequestCostHandler(),
			gin.Logger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	}
	// set up trusted agents
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}

	SetWebStudentRoute(r)
	r.NoRoute(func(c *gin.Context) {
		response.Resp().SetHttpCode(http.StatusNotFound).FailCode(c, errors.NotFound)
	})
	return r
}

func ReleaseRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.Discard
	engine := gin.New()
	return engine
}

package http

import (
	"github.com/Depado/ginprom"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lines/internal"
)

// CreateEngine creates a new gin engine and sorts CORS out.
func CreateEngine(config *internal.MainConfig) *gin.Engine {
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.CORSOrigins
	corsConfig.AllowCredentials = true
	if config.EnablePrometheus {
		AddPrometheusMetrics(r)
	}
	// TODO: Sort this out
	//r.ForwardedByClientIP = true
	//err := r.SetTrustedProxies(config.CORSOrigins)
	//if err != nil {
	//	config.Logger.Error("main", "CreateEngine", fmt.Sprintf("Error setting trusted proxies: %s", err.Error()))
	//}
	r.Use(cors.New(corsConfig))
	return r
}

func AddPrometheusMetrics(e *gin.Engine) {
	p := ginprom.New(
		ginprom.Engine(e),
		ginprom.Subsystem("SpcTrack_Backend"),
		ginprom.Path("/metrics"),
	)
	e.Use(p.Instrument())
}

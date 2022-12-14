package infra

import (
	"net/http"
	"peanut/middleware"
	"time"

	"peanut/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Router *gin.Engine
	Store  *gorm.DB
}

func SetupServer(s *gorm.DB) Server {
	// Init router
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// CORS config
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PATCH", "PUT", "DELETE"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	// Config route
	v1 := r.Group("api/v1")
	{
		userCtrl := controller.NewUserController(s)

		v1.POST("login", userCtrl.Login)
		v1.POST("register", userCtrl.CreateUser)

		users := v1.Group("/users")
		{
			users.Use(middleware.JwtAuthMiddleware())

			users.GET("", userCtrl.GetUsers)
			users.GET("/:id", userCtrl.GetUser)
			users.POST("", userCtrl.CreateUser)
			// users.PATCH("/:id", userCtrl.UpdateUser)
			// users.DELETE("/:id", userCtrl.DeleteUserByID)
		}
	}

	// health check
	r.GET("api/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return Server{
		Store:  s,
		Router: r,
	}
}

package infra

import (
	"net/http"
	"time"

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
		users := v1.Group("/users")
		{
			users.GET("", func(c *gin.Context) { userCtrl.GetUsers(c) })
			users.GET("/:id", func(c *gin.Context) { userCtrl.GetUserByID(c) })
			users.POST("", func(c *gin.Context) { userCtrl.CreateUser(c) })
			users.PATCH("/:id", func(c *gin.Context) { userCtrl.UpdateUser(c) })
			users.DELETE("/:id", func(c *gin.Context) { userCtrl.DeleteUserByID(c) })
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

package infra

import (
	"net/http"
	"time"

	"peanut/config"
	"peanut/controller"
	"peanut/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	_ "peanut/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// Custom middleware
	r.Use(middleware.HandleError)
	r.NoRoute(middleware.HandleNoRoute)
	r.NoMethod(middleware.HandleNoMethod)

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
		users := v1.Group("/users")
		{
			users.POST("/signup", userCtrl.SignUp)
			users.POST("/login", userCtrl.Login)
		}

		bookCtrl := controller.NewBookController(s)
		books := v1.Group("/books")
		{
			books.Use(middleware.JwtAuth()).POST("", bookCtrl.CreateBook)
		}
	}

	// health check
	r.GET("api/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	if config.IsDevelopment() {
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return Server{
		Store:  s,
		Router: r,
	}
}

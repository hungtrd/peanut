package infra

import (
	"net/http"
	"time"

	"peanut/controller"
	"peanut/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Router *gin.Engine
	Store  *gorm.DB
}

func SetupServer(store *gorm.DB) Server {
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
		userCtrl := controller.NewUserController(store)
		bookCrtl := controller.NewBookController(store)
		auth := v1.Group("/auth")
		{
			auth.POST("/login", userCtrl.Login)
			auth.POST("/signup", userCtrl.CreateUser)
		}
		users := v1.Group("/users")
		{
			users.GET("", middleware.Authentication, userCtrl.GetUsers)
			users.POST("", middleware.Authentication, userCtrl.CreateUser)
			users.GET("/:id", middleware.Authentication, userCtrl.GetUser)
			// users.PATCH("/:id", userCtrl.UpdateUser)
			// users.DELETE("/:id", userCtrl.DeleteUserByID)
		}

		books := v1.Group("/books")
		{
			books.GET("", middleware.Authentication, bookCrtl.GetBooks)
			books.POST("", middleware.Authentication, bookCrtl.CreateBook)
			books.POST("/:id", middleware.Authentication, bookCrtl.GetBook)
			books.PUT("/:id", middleware.Authentication, bookCrtl.UpdateBook)
			books.DELETE("/:id", middleware.Authentication, bookCrtl.DeleteBook)
		}

	}

	// health check
	r.GET("api/health", middleware.Authentication, func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	return Server{
		Store:  store,
		Router: r,
	}
}

package server

import (
	"github.com/R3l3ntl3ss/CarJacked/controllers/admin"
	"github.com/R3l3ntl3ss/CarJacked/controllers/auth"
	"github.com/R3l3ntl3ss/CarJacked/controllers/casecnt"
	"github.com/gin-gonic/contrib/jwt"
	"os"

	"github.com/R3l3ntl3ss/CarJacked/libraries/mongo"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// NewRouter : Create the routes
func NewRouter() *gin.Engine {

	router := gin.New()

	// Add Logging and Recovery Middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Create CORS with authorization
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("authorization")

	router.Use(cors.New(config))

	// Get JWT Auth Secret
	secret := os.Getenv("SECRET")

	// Initialize MongoDB server
	m := &mongo.Mongo{}
	m.Init()

	// Routes for Authentication
	authRouter := router.Group("/auth")
	{
		a := auth.Controller{
			M: m,
			Secret: secret,
		}

		authRouter.POST("/login", a.Login)
		authRouter.POST("/signUp", a.SignUp)
	}

	// Routes for admin operations [Requires Authentication]
	adminRouter := router.Group("/admin")
	adminRouter.Use(jwt.Auth(secret))
	{
		a := admin.Controller{
			M: m,
			Secret: secret,
		}

		adminRouter.GET("/getCase", a.GetCase)
		adminRouter.GET("/solveCase", a.CaseSolved)
		adminRouter.GET("/checkForNewCase", a.CheckForNewCase)
	}

	// Routes for creating and seeing cases [Public]
	caseRouter := router.Group("/case")
	{
		c := casecnt.Controller{
			M: m,
		}

		caseRouter.POST("", c.CreateCase)
		caseRouter.GET("/id/:caseID", c.GetCaseWithID)
	}

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return router
}

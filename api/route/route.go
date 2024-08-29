package route

import (
	"github.com/KhanbalaRashidov/GoCleanArchitecture/api/middleware"
	"github.com/KhanbalaRashidov/GoCleanArchitecture/envs"
	"github.com/gin-gonic/gin"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"time"
)

func Setup(env *envs.Env, timeout time.Duration, db *mongo2.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
}

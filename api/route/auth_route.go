package route

import (
	"github.com/KhanbalaRashidov/GoCleanArchitecture/api/controller"
	"github.com/KhanbalaRashidov/GoCleanArchitecture/domain"
	"github.com/KhanbalaRashidov/GoCleanArchitecture/envs"
	"github.com/KhanbalaRashidov/GoCleanArchitecture/mongo"
	"github.com/KhanbalaRashidov/GoCleanArchitecture/repository"
	"github.com/KhanbalaRashidov/GoCleanArchitecture/usecases"
	"github.com/gin-gonic/gin"
	"time"
)

func NewLoginRouter(env *envs.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	lc := &controller.LoginController{
		LoginUsecase: usecases.NewLoginUsecase(ur, timeout),
		Env:          env,
	}
	group.POST("/login", lc.Login)
}

func NewSignupRouter(env *envs.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecases.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}

func NewRefreshTokenRouter(env *envs.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecases.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}

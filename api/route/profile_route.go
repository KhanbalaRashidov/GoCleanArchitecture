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

func NewProfileRouter(env *envs.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecases.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}

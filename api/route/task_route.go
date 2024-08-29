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

func NewTaskRouter(env *envs.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecases.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}

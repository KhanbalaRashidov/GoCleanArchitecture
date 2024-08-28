package db

import (
	"github.com/KhanbalaRashidov/GoCleanArchitecture/envs"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	Env   *envs.Env
	Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = envs.NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	CloseMongoDBConnection(app.Mongo)
}

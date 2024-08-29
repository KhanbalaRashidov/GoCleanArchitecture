package main

import (
	"github.com/KhanbalaRashidov/GoCleanArchitecture/api/route"
	"github.com/KhanbalaRashidov/GoCleanArchitecture/db"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	app := db.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}

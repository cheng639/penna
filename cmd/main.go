package main

import (
	"errors"
	"fmt"
	"github.com/danielkov/gin-helmet/ginhelmet"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"os"
	"path"
	"penna/common"
	"penna/config"
	"penna/middleware"
	"penna/model"
	"penna/router"
)

func main() {
	defer common.RecoverAndLogStack()
	defer func() {
		err := config.RedisClient().Close()
		if err != nil {
			panic(err)
		}

		sqldb, err := config.GormDB().DB()
		if err != nil {
			panic(err)
		}
		err = sqldb.Close()
		if err != nil {
			panic(err)
		}
	}()

	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	println("The current working directory is ", workDir)
	config.Config().WorkDir = workDir
	config.ParseConfig(path.Join(workDir, "config.yaml"))
	config.InitLogger()
	config.InitRedisClient()
	config.InitGormDB()

	engine := gin.New()
	engine.Use(middleware.RecoveryMiddleware())
	engine.Use(middleware.RequestIDMiddleware())
	engine.Use(middleware.LoggerMiddleware())
	engine.Use(ginhelmet.Default())
	engine.Use(middleware.CORSMiddleware())

	apiV1 := engine.Group("/api/v1")
	router.RegisterRouters(engine, apiV1.BasePath())

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册 model.LocalTime 类型的自定义校验规则
		v.RegisterCustomTypeFunc(model.ValidateJSONDateType, model.LocalTime{})
	}
	config.Logger().Info().Msgf("%s is running on %s port.", config.Config().Server.Name, config.Config().Server.Port)
	err = engine.Run(config.Config().Server.IP + ":" + config.Config().Server.Port) // listens on 127.0.0.1:8090 by default
	if err != nil {
		config.Logger().Error().Err(errors.New(fmt.Sprintf("%v", err))).Msgf("%s start failed !", config.Config().Server.Name)
	}
}

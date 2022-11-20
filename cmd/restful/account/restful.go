package main

import (
	"time"

	"github.com/blackhorseya/irent/cmd"
	"github.com/blackhorseya/irent/pkg/contextx"
	"github.com/blackhorseya/irent/pkg/cors"
	ab "github.com/blackhorseya/irent/pkg/entity/domain/account/biz"
	"github.com/blackhorseya/irent/pkg/errorx"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type restful struct {
	router *gin.Engine
	biz    ab.IBiz
}

func NewRestful(logger *zap.Logger, router *gin.Engine, biz ab.IBiz) cmd.Restful {
	router.Use(cors.AddAllowAll())
	router.Use(ginzap.RecoveryWithZap(logger, true))
	router.Use(ginzap.GinzapWithConfig(logger, &ginzap.Config{
		TimeFormat: time.RFC3339,
		UTC:        true,
		SkipPaths:  []string{"/metrics"},
	}))
	router.Use(contextx.AddContextxWitLoggerMiddleware(logger))
	router.Use(errorx.AddErrorHandlingMiddleware())

	return &restful{
		router: router,
		biz:    biz,
	}
}

func (i *restful) InitRouting() error {
	// todo: 2022/11/21|sean|impl me
	return nil
}

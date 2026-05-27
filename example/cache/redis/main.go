package main

import (
	"context"

	"github.com/gin-gonic/gin"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
)

var cacheEntry *rkcache.CacheEntry

func main() {
	boot := rkboot.NewBoot()

	boot.Bootstrap(context.TODO())

	// assign cache
	cacheEntry = rkcache.GetCacheEntry("redis-cache")

	// assign router
	ginEntry := rkgin.GetGinEntry("cache-service")
	ginEntry.Router.GET("/v1/get", Get)
	ginEntry.Router.GET("/v1/set", Set)

	boot.WaitForShutdownSig(context.TODO())
}

func Get(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func Set(ctx *gin.Context) { _ = "STUB: not implemented"; return }

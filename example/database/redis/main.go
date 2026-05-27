// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-db/redis"
)

var redisClient *redis.Client

func main() {
	boot := rkboot.NewBoot()

	boot.Bootstrap(context.TODO())

	// Auto migrate database and init global userDb variable
	redisEntry := rkredis.GetRedisEntry("redis")
	redisClient, _ = redisEntry.GetClient()

	// Register APIs
	ginEntry := rkgin.GetGinEntry("server")
	ginEntry.Router.GET("/v1/get", Get)
	ginEntry.Router.POST("/v1/set", Set)

	boot.WaitForShutdownSig(context.TODO())
}

type KV struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Set(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func Get(ctx *gin.Context) { _ = "STUB: not implemented"; return }

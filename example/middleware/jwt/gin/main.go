// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-gin/v2/middleware/context"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample rk-boot server.
// @termsOfService http://swagger.io/terms/

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Create a new boot instance.
	boot := rkboot.NewBoot()

	// Register handler
	entry := rkgin.GetGinEntry("greeter")
	entry.Router.GET("/v1/login", Login)
	entry.Router.GET("/v1/whoami", WhoAmI)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	boot.WaitForShutdownSig(context.TODO())
}

// CustomClaims defines JWT claims
type CustomClaims struct {
	UserName string `json:"uname"`
	jwt.RegisteredClaims
}

// Login handler
// @Summary Login
// @Id 1
// @Tags JWT
// @version 1.0
// @Security  JWT
// @Param name query string true "name"
// @produce application/json
// @Router /v1/login [get]
func Login(ctx *gin.Context) {
	_ = "STUB: not implemented"
	// Simply generate JWT token from user provided name for demo
	return
}

// By default, JWT middleware will create a new SignerEntry with the same name of Gin Entry
// default signer entry will use symmetric algorithm (HS256) with token of (rk jwt key)
// refer rkmidjwt.NewOptionSet

// WhoAmI handler
// @Summary WhoAmI
// @Id 2
// @Tags JWT
// @version 1.0
// @Security  JWT
// @produce application/json
// @Router /v1/whoami [get]
func WhoAmI(ctx *gin.Context) {
	_ = "STUB: not implemented"
	// 1: get JWT token from context which injected into context by middleware
	return
}

// convert claim to custom claim

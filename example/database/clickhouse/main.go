// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	"gorm.io/gorm"
)

var userDb *gorm.DB

func main() {
	boot := rkboot.NewBoot()

	boot.Bootstrap(context.TODO())

	// Auto migrate database and init global userDb variable
	clickHouseEntry := rkclickhouse.GetClickHouseEntry("user-db")
	userDb = clickHouseEntry.GetDB("user")
	if !userDb.DryRun {
		userDb.AutoMigrate(&User{})
	}

	// Register APIs
	ginEntry := rkgin.GetGinEntry("user-service")
	ginEntry.Router.GET("/v1/user", ListUsers)
	ginEntry.Router.GET("/v1/user/:id", GetUser)
	ginEntry.Router.PUT("/v1/user", CreateUser)
	ginEntry.Router.POST("/v1/user/:id", UpdateUser)
	ginEntry.Router.DELETE("/v1/user/:id", DeleteUser)

	boot.WaitForShutdownSig(context.TODO())
}

// *************************************
// *************** Model ***************
// *************************************

type Base struct {
	CreatedAt time.Time `yaml:"-" json:"-"`
	UpdatedAt time.Time `yaml:"-" json:"-"`
}

type User struct {
	Base
	Id   string `yaml:"id" json:"id"`
	Name string `yaml:"name" json:"name"`
}

func ListUsers(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func GetUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func CreateUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func UpdateUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func DeleteUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

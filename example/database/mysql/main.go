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

var model interface{}

var userDb *gorm.DB

func main() {
	boot := rkboot.NewBoot()

	boot.Bootstrap(context.TODO())

	// Auto migrate database and init global userDb variable
	mysqlEntry := rkmysql.GetMySqlEntry("user-db")
	userDb = mysqlEntry.GetDB("demo")
	model = User{}

	if !userDb.DryRun {
		userDb.AutoMigrate(&User{})
	}

	db, _ := userDb.DB()
	res, _ := db.Query("SHOW TABLES")

	var table string

	for res.Next() {
		res.Scan(&table)
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
	CreatedAt time.Time      `yaml:"-" json:"-"`
	UpdatedAt time.Time      `yaml:"-" json:"-"`
	DeletedAt gorm.DeletedAt `yaml:"-" json:"-" gorm:"index"`
}

type User struct {
	Base
	Id   int    `yaml:"id" json:"id" gorm:"primaryKey"`
	Name string `yaml:"name" json:"name"`
}

func ListUsers(ctx *gin.Context) { _ = "STUB: not implemented"; return }

// inject logger with requestId and traceId

func GetUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func CreateUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func UpdateUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

// get user

func DeleteUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

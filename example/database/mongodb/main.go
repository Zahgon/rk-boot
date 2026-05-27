package main

import (
	"context"

	"github.com/gin-gonic/gin"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	userCollection *mongo.Collection
)

func createCollection(db *mongo.Database, name string) { _ = "STUB: not implemented"; return }

func main() {
	boot := rkboot.NewBoot()

	boot.Bootstrap(context.TODO())

	// Auto migrate database and init global userDb variable
	db := rkmongo.GetMongoDB("my-mongo", "users")
	createCollection(db, "meta")

	userCollection = db.Collection("meta")

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

type User struct {
	Id   string `bson:"id" yaml:"id" json:"id"`
	Name string `bson:"name" yaml:"name" json:"name"`
}

func ListUsers(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func GetUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func CreateUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func UpdateUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

func DeleteUser(ctx *gin.Context) { _ = "STUB: not implemented"; return }

package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkmongo "github.com/rookie-ninja/rk-db/mongodb"
	rkgin "github.com/rookie-ninja/rk-gin/v2/boot"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	todoCollection *mongo.Collection
)

// @title Swagger API With MongoDB
// @version 1.0
// @description This is Example MongoDB.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	boot := rkboot.NewBoot()
	boot.Bootstrap(context.TODO())

	db := rkmongo.GetMongoDB("todo-mongo", "tododb")
	todoCollection = db.Collection("todos")

	// Register APIs
	ginEntry := rkgin.GetGinEntry("todo-service")
	ginEntry.Router.GET("/", HelloWorld)
	ginEntry.Router.GET("/v1/todos", ListTodos)
	ginEntry.Router.GET("/v1/todo/:id", GetTodo)
	ginEntry.Router.POST("/v1/todo", CreateTodo)
	ginEntry.Router.PUT("/v1/todo/:id", UpdateTodo)
	ginEntry.Router.DELETE("/v1/todo/:id", DeleteTodo)

	boot.WaitForShutdownSig(context.TODO())
}

// *************************************
// *************** Model ***************
// *************************************

type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Body      string             `json:"body" bson:"body"`
	Completed bool               `json:"completed" bson:"completed"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// @Summary Hello world
// @Tags Hello world
// @version 1.0
// @produce application/json
// @Success 200
// @Router / [get]
func HelloWorld(ctx *gin.Context) { _ = "STUB: not implemented"; return }

// @Summary Get List Todo
// @ID get-todo-list
// @Tags Todos
// @version 1.0
// @produce application/json
// @Success 200
// @Failure 500
// @Router /v1/todos [get]
func ListTodos(ctx *gin.Context) { _ = "STUB: not implemented"; return }

// @Summary Get Todo
// @ID get-todo-by-id
// @Param id path string true "todo ID"
// @Tags Todos
// @version 1.0
// @produce application/json
// @Success 200
// @Failure 500
// @Router /v1/todo/{id} [get]
func GetTodo(ctx *gin.Context) { _ = "STUB: not implemented"; return }

// @Summary Post Todo
// @Tags Todos
// @version 1.0
// @produce application/json
// @Success 201
// @Failure 500
// @Router /v1/todo [post]
func CreateTodo(ctx *gin.Context) { _ = "STUB: not implemented"; return }

// default value if not input

// @Summary Update Todo
// @ID update-todo-by-id
// @Param id path string true "todo ID"
// @Tags Todos
// @version 1.0
// @produce application/json
// @Success 200
// @Failure 404
// @Router /v1/todo/{id} [put]
func UpdateTodo(ctx *gin.Context) { _ = "STUB: not implemented"; return }

// default value if not input

// @Summary Delete Todo
// @ID delete-todo-by-id
// @Param id path string true "todo ID"
// @Tags Todos
// @version 1.0
// @produce application/json
// @Success 200
// @Failure 404
// @Router /v1/todo/{id} [delete]
func DeleteTodo(ctx *gin.Context) { _ = "STUB: not implemented"; return }

package main

import (
	"api/src/controllers"
	"context"
	"log"
	"time"

	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	os.Setenv("PORT", "3333")

	app := gin.Default()

	app.Use(cors.Default())

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://user:pass@localhost:27017/test?authSource=test"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	app.GET("/users", func(c *gin.Context) {
		controllers.GetAllUsers(c, client)
	})

	app.Run()

}

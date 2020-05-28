package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"net/http"
)

// GetAllUsers handle the route users in http GET method
func GetAllUsers(ctx *gin.Context, db *mongo.Client) {
	usersCollection := db.Database("go").Collection("users")
	ctxBack := context.Background()
	cursor, err := usersCollection.Find(ctxBack, bson.D{})
	if err != nil {
		ctx.String(http.StatusBadRequest, "")
	}

	var users []bson.M

	if err = cursor.All(ctxBack, &users); err != nil {
		ctx.String(http.StatusBadRequest, "")
	}

	ctx.JSON(http.StatusOK, users)

}

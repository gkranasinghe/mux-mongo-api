
package controllers

import (
    "context"
    "encoding/json"
    "mux-mongo-api/configs"
    "mux-mongo-api/models"
    "mux-mongo-api/responses"
    "net/http"
    "time"
    "github.com/go-playground/validator/v10"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        var user models.User
        defer cancel()

        //validate the request body
        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            rw.WriteHeader(http.StatusBadRequest)
            response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        //use the validator library to validate required fields
        if validationErr := validate.Struct(&user); validationErr != nil {
            rw.WriteHeader(http.StatusBadRequest)
            response := responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        newUser := models.User{
            Id:       primitive.NewObjectID(),
            Name:     user.Name,
            Location: user.Location,
            Title:    user.Title,
        }
        result, err := userCollection.InsertOne(ctx, newUser)
        if err != nil {
            rw.WriteHeader(http.StatusInternalServerError)
            response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        rw.WriteHeader(http.StatusCreated)
        response := responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
        json.NewEncoder(rw).Encode(response)
    }
}
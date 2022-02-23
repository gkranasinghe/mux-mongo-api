
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
	"github.com/gorilla/mux" //add this
    "go.mongodb.org/mongo-driver/bson" //add this
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

func GetAUser() http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        params := mux.Vars(r)
        userId := params["userId"]
        var user models.User
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(userId)

        err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)
        if err != nil {
            rw.WriteHeader(http.StatusInternalServerError)
            response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        rw.WriteHeader(http.StatusOK)
        response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}}
        json.NewEncoder(rw).Encode(response)
    }
}

func EditAUser() http.HandlerFunc {
    return func(rw http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
        params := mux.Vars(r)
        userId := params["userId"]
        var user models.User
        defer cancel()

        objId, _ := primitive.ObjectIDFromHex(userId)

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

        update := bson.M{"name": user.Name, "location": user.Location, "title": user.Title}

        result, err := userCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})
        if err != nil {
            rw.WriteHeader(http.StatusInternalServerError)
            response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
            json.NewEncoder(rw).Encode(response)
            return
        }

        //get updated user details
        var updatedUser models.User
        if result.MatchedCount == 1 {
            err := userCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)

            if err != nil {
                rw.WriteHeader(http.StatusInternalServerError)
                response := responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
                json.NewEncoder(rw).Encode(response)
                return
            }
        }

        rw.WriteHeader(http.StatusOK)
        response := responses.UserResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}}
        json.NewEncoder(rw).Encode(response)
    }
}
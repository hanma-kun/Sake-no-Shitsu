package database

import (
	"log"
	"net/http"
	"sake/types"
	"sake/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gitlab.com/avarf/getenvs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Usercollections = ClientConnection.Database("sakeno").Collection("users")
var TokenCollection = ClientConnection.Database("sakeno").Collection("tokens")
var Validate = validator.New()
var Secret string = getenvs.GetEnvString("SECRET", "qfwtwXWUayQKDsooCEC9RzbA3yYlXTua")

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := CreateContext()
		defer cancel()

		var user types.Users
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := Validate.Struct(user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}

		if count, err := Usercollections.CountDocuments(ctx, bson.M{"email": user.Email}); err != nil || count > 0 {
			defer cancel()
			if count > 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Email already is registered with other account"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if count, err := Usercollections.CountDocuments(ctx, bson.M{"phone": user.Phone}); err != nil || count > 0 {
			defer cancel()
			if count > 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Mobile number is already registered with other account"})
				return
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user.Id = primitive.NewObjectID()
		user.Password = utils.HashPassword(user.Password)
		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Address_Details = make([]types.Address, 0)
		user.UserCart = make([]types.Products, 0)
		user.Order_Detials = make([]types.Order, 0)
		tokendata, err := utils.TokenGenerator(user.Email, user.FirstName, user.LastName, user.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if _, err = Usercollections.InsertOne(ctx, user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create the record"})
			log.Fatal(err.Error())
			return
		}
		if _, err = TokenCollection.InsertOne(ctx, tokendata); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create the record"})
			log.Fatal(err.Error())
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, tokendata)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := CreateContext()
		defer cancel()
		var user types.LoginId
		var founduser types.Users
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		if err := Validate.Struct(user); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		err := Usercollections.FindOne(ctx, bson.M{"email": user.Email}).Decode(&founduser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to authenticate the given credentials . Please try again later "})
			return
		}
		if !utils.VerifyPassword(user.Password, founduser.Password) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to authenticate the given credentials . Please try again later "})
			return
		}
		token, err := utils.TokenGenerator(founduser.Email, founduser.FirstName, founduser.LastName, founduser.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var updateObject primitive.D
		updateObject = append(updateObject, bson.E{Key: "access_token", Value: token.Token})
		updateObject = append(updateObject, bson.E{Key: "access_token_exp", Value: token.TokenExp})
		updateObject = append(updateObject, bson.E{Key: "refresh_token", Value: token.Refresh_Token})
		updateObject = append(updateObject, bson.E{Key: "refresh_token_exp", Value: token.Refresh_Token_exp})
		uprest := true
		if _, err = TokenCollection.UpdateOne(ctx, bson.M{"user_id": founduser.Id}, bson.M{"$set": updateObject}, &options.UpdateOptions{Upsert: &uprest}); err != nil {
			defer cancel()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, token)
	}
}

func Refresh_Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := CreateContext()
		defer cancel()
		var refresh_token types.RefreshToken

		if err := c.BindJSON(&refresh_token); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := jwt.ParseWithClaims(refresh_token.Refreshtoken, &types.SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(Secret), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, ok := token.Claims.(*types.SignedDetails)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token invalid"})
			return
		}

		var prevtoken types.Tokens
		obj, err := primitive.ObjectIDFromHex(claims.Uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := TokenCollection.FindOne(ctx, bson.M{"user_id": obj}).Decode(&prevtoken); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if prevtoken.Refresh_Token != refresh_token.Refreshtoken {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
			return
		}
		var founduser types.Users
		err = Usercollections.FindOne(ctx, bson.M{"_id": obj}).Decode(&founduser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to authenticate the given credentials . Please try again later "})
			return
		}
		access_token, err := utils.TokenGenerator(founduser.Email, founduser.FirstName, founduser.LastName, founduser.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var updateObject primitive.D
		updateObject = append(updateObject, bson.E{Key: "access_token", Value: access_token.Token})
		updateObject = append(updateObject, bson.E{Key: "access_token_exp", Value: access_token.TokenExp})
		updateObject = append(updateObject, bson.E{Key: "refresh_token", Value: access_token.Refresh_Token})
		updateObject = append(updateObject, bson.E{Key: "refresh_token_exp", Value: access_token.Refresh_Token_exp})
		uprest := true
		if _, err = TokenCollection.UpdateOne(ctx, bson.M{"user_id": founduser.Id}, bson.M{"$set": updateObject}, &options.UpdateOptions{Upsert: &uprest}); err != nil {
			defer cancel()
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, access_token)
	}
}

package database

import (
	"net/http"
	"sake/types"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ProductCollection = ClientConnection.Database("sakeno").Collection("products")

func PlaceOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := CreateContext()
		defer cancel()
		uid, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "The token doesnot exists"})
			return
		}
		userId, ok := uid.(string)
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User doesnot exists"})
			return
		}
		uid, err := primitive.ObjectIDFromHex(userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User doesnot exists"})
			return
		}

		var userData types.Users
		if err := Usercollections.FindOne(ctx, bson.M{"_id": uid}).Decode(&userData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		var OrderData types.Order
		if err := c.Bind(&OrderData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		OrderData.Id = primitive.NewObjectID()
		OrderData.Ordered_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		if _, err := Usercollections.UpdateOne(ctx, bson.M{"_id": uid}, bson.M{"$push": bson.M{"orders": OrderData}}); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, OrderData)
	}
}

func ProductFetch() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := CreateContext()
		defer cancel()
		var products []types.Products
		var filter bson.M
		category := c.Query("category")
		if category != "" {
			filter = bson.M{"category": category}
		} else {
			filter = bson.M{}
		}

		cursor, err := ProductCollection.Find(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		for cursor.Next(ctx) {
			var product types.Products
			if err := cursor.Decode(&product); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			products = append(products, product)
		}
		c.JSON(http.StatusOK, products)
	}
}

func ProductDetials() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := CreateContext()
		defer cancel()
		id := c.Query("id")
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No id was sent through"})
			return
		}
		uid, _ := primitive.ObjectIDFromHex(id)
		var product types.Products
		if err := ProductCollection.FindOne(ctx, bson.M{"_id": uid}).Decode(&product); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unable to find the record"})
			return
		}

		c.JSON(http.StatusOK, product)
	}
}

func AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := CreateContext()
		defer cancel()
		cid := c.Query("id")
		if cid == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No id was sent through"})
			return
		}
		userid, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No uid was sent through"})
			return
		}

		pid, err := primitive.ObjectIDFromHex(cid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No id was sent through"})
			return
		}

		search, err := ProductCollection.Find(ctx, bson.M{"_id": pid})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No id was sent through"})
			return
		}
		var products []types.Products
		if err = search.All(ctx, &products); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		suid, ok := userid.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to convert the token "})
			return
		}
		uid, err := primitive.ObjectIDFromHex(suid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		filter := bson.D{primitive.E{Key: "_id", Value: uid}}
		update := bson.D{{Key: "$push", Value: bson.D{primitive.E{Key: "usercart", Value: bson.D{{Key: "$each", Value: products}}}}}}

		if _, err := Usercollections.UpdateOne(ctx, filter, update); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
	}
}

func GetCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := CreateContext()
		defer cancel()
		id, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unable to find an userid"})
			return
		}
		uid, ok := id.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse the id"})
			return
		}
		k, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		var user types.Users
		if err := Usercollections.FindOne(ctx, bson.M{"_id": k}).Decode(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"cart": user.UserCart, "address": user.Address_Details})
	}
}

func RemoveCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := CreateContext()
		defer cancel()
		puids := c.Query("id")
		kid, exists := c.Get("uid")
		if !exists {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unable to find an userid"})
			return
		}
		uid, ok := kid.(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse the id"})
			return
		}
		productID, err := primitive.ObjectIDFromHex(puids)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		id, err := primitive.ObjectIDFromHex(uid)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		filter := bson.D{primitive.E{Key: "_id", Value: id}}
		update := bson.M{"$pull": bson.M{"usercart": bson.M{"_id": productID}}}
		f, err := Usercollections.UpdateMany(ctx, filter, update)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, f)
	}
}

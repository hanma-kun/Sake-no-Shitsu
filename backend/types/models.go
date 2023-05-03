package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)
type LoginId struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type Tokens struct {
	Id                primitive.ObjectID `bson:"_id" json:"session_id"`
	User_Id           primitive.ObjectID `bson:"user_id" json:"user_id"`
	Token             string             `bson:"access_token" json:"access_token"`
	TokenExp          time.Time          `bson:"access_token_exp" json:"access_token_exp"`
	Refresh_Token     string             `bson:"refresh_token" json:"refresh_token"`
	Refresh_Token_exp time.Time          `bson:"refresh_token_exp" json:"refresh_token_exp"`
}

type Users struct {
	Id              primitive.ObjectID `bson:"_id" json:"_id"`
	FirstName       string             `bson:"firstname" json:"firstname" validate:"required,omitempty,min=1,max=30" `
	LastName        string             `bson:"lastname" json:"lastname" validate:"required,omitempty,min=1,max=30"`
	Password        string             `bson:"password" json:"password" validate:"required,omitempty,min=6"`
	Email           string             `bson:"email" json:"email" validate:"required,email"`
	Phone           string             `bson:"phone" json:"phone" validate:"required"`
	Created_At      time.Time          `bson:"created_at" json:"created_at"`
	Updated_At      time.Time          `bson:"updated_at" json:"updated_at"`
	UserCart        []Products         `bson:"usercart" json:"usercart"`
	Address_Details []Address          `bson:"address" json:"address"`
	Order_Detials   []Order            `bson:"orders" json:"orders"`
}

type Products struct {
	Id           primitive.ObjectID `bson:"_id" json:"_id"`
	Category     string             `bson:"category" json:"category"`
	Product_Name string             `bson:"product_name" json:"product_name"`
	Price        uint               `bson:"price" json:"price"`
	Rating       uint               `bson:"rating" json:"rating"`
	Image        string             `bson:"image" json:"image"`
}

type Address struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id"`
	House   string             `bson:"house_name" json:"house_name"`
	Street  string             `bson:"street_name" json:"street_name"`
	City    string             `bson:"city_name" json:"city_name"`
	Pincode string             `bson:"pincode" json:"pincode"`
}

type Order struct {
	Id             primitive.ObjectID `bson:"_id" json:"_id"`
	Cart           []Products         `bson:"order_cart" json:"order_cart"`
	Ordered_At     time.Time          `bson:"ordered_at" json:"ordered_at"`
	Price          uint               `bson:"price" json:"price"`
	Payment_method string             `bson:"payment" json:"payment"`
}

type UploadProducts struct {
	Category     string `json:"category" binding:"required"`
	Product_Name string `json:"productname" binding:"required"`
	Price        uint   `json:"price" binding:"required"`
	Rating       uint   `json:"rating" binding:"required"`
	Image        string `json:"image" binding:"required"`
}

type UploadAddress struct {
	House   string `bson:"house_name" json:"house_name"`
	Street  string `bson:"street_name" json:"street_name"`
	City    string `bson:"city_name" json:"city_name"`
	Pincode string `bson:"pincode" json:"pincode"`
}

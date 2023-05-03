package utils

import (
	"fmt"
	"log"
	"sake/types"
	"time"

	"github.com/dgrijalva/jwt-go"
	"gitlab.com/avarf/getenvs"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Secret string = getenvs.GetEnvString("SECRET", "qfwtwXWUayQKDsooCEC9RzbA3yYlXTua")

func TokenGenerator(email, firstname, lastname string, uid primitive.ObjectID) (types.Tokens, error) {
	access_token_exp := time.Now().Local().Add(time.Hour * time.Duration(24))
	refresh_token_exp := time.Now().Local().Add(time.Hour * time.Duration(168))
	claims := &types.SignedDetails{
		Email:     email,
		Firstname: firstname,
		Lastname:  lastname,
		Uid:       uid.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: access_token_exp.Unix(),
		},
	}

	refreshClaims := &types.SignedDetails{
		Uid: uid.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refresh_token_exp.Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(Secret))
	if err != nil {
		log.Fatal(err)
	}
	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(Secret))
	if err != nil {
		log.Fatal(err)
	}
	tokenStruct := types.Tokens{
		Id:                primitive.NewObjectID(),
		User_Id:           uid,
		Token:             token,
		TokenExp:          access_token_exp,
		Refresh_Token:     refreshtoken,
		Refresh_Token_exp: refresh_token_exp,
	}

	return tokenStruct, err

}

func ValidateToken(signedToken string) (*types.SignedDetails, error) {
	token, err := jwt.ParseWithClaims(signedToken, &types.SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return &types.SignedDetails{}, err
	}
	claims, ok := token.Claims.(*types.SignedDetails)
	if !ok {
		return &types.SignedDetails{}, fmt.Errorf("token is invalid")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return &types.SignedDetails{}, fmt.Errorf("token have expired")
	}
	return claims, nil
}

type Claims struct {
	Uid string
	jwt.StandardClaims
}

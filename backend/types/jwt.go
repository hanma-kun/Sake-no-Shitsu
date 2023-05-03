package types

import jwt "github.com/dgrijalva/jwt-go"

type SignedDetails struct {
	Email     string
	Firstname string
	Lastname  string
	Uid       string
	jwt.StandardClaims
}

type RefreshToken struct {
	Refreshtoken string `json:"refresh_token"`
}

package define

import "github.com/golang-jwt/jwt/v5"

type UserClaim struct {
	Id       uint   `json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var JwtKey = "gozero"

type M map[string]interface{}

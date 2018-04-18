package controllers

import (
	// "errors"
	"github.com/astaxie/beego"
	// "github.com/dgrijalva/jwt-go"
	// "shop/libs/jwt-go"
	// "strings"
	// "time"
)

type BaseController struct {
	beego.Controller
}

// ParseToken parse JWT token in http header.
// func (c *BaseController) ParseToken() (t *jwt.Token, e error) {

// 	authString := c.Ctx.Input.Header("Authorization")
// 	beego.Debug("AuthString:", authString)

// 	kv := strings.Split(authString, " ")
// 	if len(kv) != 2 || kv[0] != "Bearer" {
// 		beego.Error("AuthString invalid:", authString)
// 		return nil, errors.New("errInputData")
// 	}
// 	tokenString := kv[1]

// 	// Parse token
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return []byte("yooplus"), nil
// 	})
// 	if err != nil {
// 		beego.Error("Parse token:", err)
// 		if ve, ok := err.(*jwt.ValidationError); ok {
// 			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
// 				// That‘s not even a token
// 				return nil, errors.New("errInputData")
// 			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
// 				// Token is either expired or not active yet
// 				return nil, errors.New("errExpired")
// 			} else {
// 				// Couldn‘t handle this token
// 				return nil, errors.New("errInputData")
// 			}
// 		} else {
// 			// Couldn‘t handle this token
// 			return nil, errors.New("errInputData")
// 		}
// 	}
// 	if !token.Valid {
// 		beego.Error("Token invalid:", tokenString)
// 		return nil, errors.New("errInputData")
// 	}
// 	beego.Debug("Token:", token)

// 	return token, nil
// }

// //验证token是否有效
// func (this *BaseController) indicateToken() (bool, int64, error) {

// 	token, err := this.ParseToken()
// 	if err != nil {
// 		return false, -1, err
// 	}
// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return false, -1, nil
// 	}
// 	var userId int64 = claims["userId"].(int64)
// 	return true, userId, nil
// }

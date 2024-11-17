package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com./golang-jwt/jwt/v5"
	"github.com/ChanchalS7/practice_api/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

//retrieve JWT key from .env file

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

//Generate JWT token
func GenerateJWT(user model.User)(string, error){
	tokenTTL,_ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id" : user.ID,
		"role":user.RoleID,
		"iat":time.Now().Unix(),
		"eat":time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),

	})
	return token.SignedString(privateKey)
}

//validate JWT key
func ValidateJWT(context *gin.Context) error{
	token, err := getToken(context)
	if err!=nil{
		return err
	}
	_,ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("Invalid token provided")
}

//Validate admin role
func ValidateAdminRoleJWT(context *gin.Context) error{
	token, err := getToken(context)
	if err!=nil{
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 1 {
		return nil
	}
	return errors.New("Invalid admin token provided")
}

//Validate customer role
func ValidateCustomerRoleJWT(context *gin.Context) error {
	token, err:= getToken(context)
	if err !=nil{
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["role"].(float64))
	if ok && token.Valid && userRole == 2 || userRole == 1 {
		return nil
	}
	return errors.New("Invalid author token provided")
}
//fetch User details from the token
func CurrentUser(context *gin.Context) model.User{
	err := ValidateJWT(context)
	if err !=nil{
		return model.User{}
	}
	token, _ := getToken(context)
	claims, _:= token.Claims.(jwt.MapClaims)
	userId :=uint(claims["id"].(float64))
	user, err:= model.GetUserById(userId)
	if err!=nil{
		return model.User{}
	}
	return user 
}
//check token validity
func getToken(context *gin.Context) (*jwt.Token, error){
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)(interface{},error){
		if _,ok:= token.Method(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("Unexpected signing method:%v", token.Header["alg"])
		}
		return privateKey,nil
	})
	return token,err
}
//extract token from request Authorization header
func getTokenFromRequest(context *gin.Context)string{
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken :=strings.Split(bearerToken," ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
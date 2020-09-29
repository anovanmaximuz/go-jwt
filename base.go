package main

import (
	"fmt"
	"net/http"

	"./config"
	"./controllers"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


func main() {
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/login", loginHandler)
	router.GET("/person/:id", auth, inDB.GetPerson)
	router.GET("/persons", auth, inDB.GetPersons)
	router.POST("/person", auth, inDB.CreatePerson)
	router.PUT("/person", auth, inDB.UpdatePerson)
	router.DELETE("/person/:id", auth, inDB.DeletePerson)
	router.Run(":8080")
}

func response(code,message,data string) string{
     return "ok";
}

func loginHandler(c *gin.Context) {
	var user Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	if user.Username != "anovan" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"message": "wrong username or password",
		})
	} else 	if user.Password != "ano123" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  http.StatusUnauthorized,
				"message": "wrong username or password",
			})
		
	}else{
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"code":http.StatusOK,"message":"success","data": token,
	})
       }
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	// if token.Valid && err == nil {
	if token != nil && err == nil {
		fmt.Println("token verified")
	} else {
		result := gin.H{
			"code":http.StatusUnauthorized,
			"message": "not authorized",
			"error":   err.Error(),
		}
		c.JSON(http.StatusUnauthorized, result)
		c.Abort()
	}
}

package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthorizeJWT authorizes or refuses a JWT
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"Error": "Authentication failed - missing authorization header",
			})
		} else {
			tokenString := strings.TrimSpace(authHeader[len(BearerSchema):])
			token, err := verifyToken(tokenString)
			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)
				fmt.Println("claims: ", claims)

			} else {
				fmt.Println(err.Error())
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"Error": "Authentication failed - " + err.Error(),
				})
			}
		}
	}
}

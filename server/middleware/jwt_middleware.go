package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	utils "github.com/marijakljestan/golang-web-app/server/util"
	"net/http"
)

func AuthorizeJWT(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const BearerSchema string = "Bearer "
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" || authHeader == "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No Authorization header provided"})
		}

		tokenString := authHeader[len(BearerSchema):]
		if token, err := utils.ValidateToken(tokenString); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Provided token is not valid"})
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				ctx.AbortWithStatus(http.StatusUnauthorized)
			} else {
				if token.Valid {
					if claims["role"] != requiredRole {
						ctx.AbortWithStatus(http.StatusUnauthorized)
					} else {
						ctx.Set("username", claims["username"])
						ctx.Set("role", claims["role"])
					}
				} else {
					ctx.AbortWithStatus(http.StatusUnauthorized)
				}
			}
		}
	}
}

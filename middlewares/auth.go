package middlewares

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secretkey"), nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		ctx.Set("claims", claims)

		ctx.Next()
	}
}

// func JWTMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		authHeader := c.GetHeader("Authorization")
// 		if authHeader == "" {
// 			c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Missing authorization header"))
// 			c.Abort()
// 			return
// 		}

// 		// Ambil token dari header
// 		splitToken := strings.Split(authHeader, "Bearer ")
// 		if len(splitToken) != 2 {
// 			c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid authorization header"))
// 			c.Abort()
// 			return
// 		}

// 		tokenString := splitToken[1]

// 		// Verifikasi token JWT
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, utils.ErrorResponse("Invalid token signing method")
// 			}
// 			return config.JWTKey, nil
// 		})
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, utils.ErrorResponse(err.Error()))
// 			c.Abort()
// 			return
// 		}

// 		claims, ok := token.Claims.(jwt.MapClaims)
// 		if !ok || !token.Valid {
// 			c.JSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid token"))
// 			c.Abort()
// 			return
// 		}

// 		// Set nilai pengguna pada konteks Gin
// 		c.Set("user_id", claims["user_id"])

// 		c.Next()
// 	}
// }

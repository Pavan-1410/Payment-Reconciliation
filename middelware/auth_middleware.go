package middelware

import (
	"net/http"
	"os"
	"payment_reconciliation/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {		//gin.HandlerFunc means a function that takes *gin.Context as its argument.
	return func(c *gin.Context){	// this is that function
		// get headers
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			c.Abort()
			return
		}

		// check bearer

		parts := strings.Split(authHeader, " ")	// split by space

		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization format or Bearer is missing",
			})
			c.Abort()
			return
		}

		tokenString :=parts[1]

		// 3. Parse and validate JWT
		token, err := jwt.ParseWithClaims(
			tokenString,
			&utils.Claims{},
			func(token *jwt.Token) (interface{}, error) {

				if token.Method != jwt.SigningMethodHS256 {
					return nil, jwt.ErrSignatureInvalid
				}

				return []byte(os.Getenv("JWT_SECRET")), nil
			},
		)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			c.Abort()
			return
		}

		// 4. Get our custom claims
		claims, ok := token.Claims.(*utils.Claims)

		if !ok{
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			return 
		}
		// 5. Store user information in Gin context
		c.Set("user_id", claims.UserID)
		c.Set("is_admin", claims.IS_Admin)

		// 6. Continue to the next handler
		c.Next()


	}		
}
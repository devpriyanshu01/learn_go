package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"restapi/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
	fmt.Println("--------------- JWT MIDDLEWARE ---------------")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("--------------- INSIDE JWT MIDDLEWARE ----------------")

		token, err := r.Cookie("Bearer")
		if err != nil {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		} else {
			fmt.Println("token received", token)
			fmt.Println("token.value", token.Value)
		}

		jwtSecret := os.Getenv("JWT_SECRET")

		parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			//hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(jwtSecret), nil
		})

		fmt.Println("parsed token claims:", parsedToken.Claims)

		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				http.Error(w, "Token Expired", http.StatusUnauthorized)
				return
			} else if errors.Is(err, jwt.ErrTokenMalformed) {
				http.Error(w, "Token Malformed", http.StatusUnauthorized)
				return
			}
			utils.ErrorHandler(err, "")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if parsedToken.Valid {
			log.Println("valid JWT")
		} else {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
			log.Println("Invalid Login Token:", token.Value)
			return
		}

		type Contextkey string

		ctx := context.WithValue(r.Context(), Contextkey("role"), claims["role"])
		ctx = context.WithValue(ctx, Contextkey("expiredAt"), claims["exp"])
		ctx = context.WithValue(ctx, Contextkey("username"), claims["user"])
		ctx = context.WithValue(ctx, Contextkey("userId"), claims["uid"])

		next.ServeHTTP(w, r.WithContext(ctx))
		fmt.Println("Sent Response from JWT Middleware")

	})
}

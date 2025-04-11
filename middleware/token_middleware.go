package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

// 👇 custom context key type banaya
type contextKey string

const userContextKey = contextKey("user")

func verifyToken(tokenStr string, secret string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))

		claims, err := verifyToken(tokenStr, os.Getenv("ACCESS_TOKEN_SECRET"))
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		fmt.Println("✅ Token verified successfully for user:", claims["email"]) // assuming your claims contain "email"
		ctx := context.WithValue(r.Context(), userContextKey, claims)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/contextkey"
	 role "github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/roles"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/response"
	"github.com/Yash-Watchguard/MovieTicketBooking/utills"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader  == "" {
			response.ErrorResponse(w, "Authorization token in not present", http.StatusBadRequest)
			return
		}

		if !strings.HasPrefix(authorizationHeader , "Bearer ") {
			response.ErrorResponse(w, "invalid token", http.StatusBadRequest)
			return
		}

		tokenString := strings.TrimPrefix(authorizationHeader , "Bearer ")

		token, err := utills.VarifyJwt(tokenString)
		if err != nil {
			response.ErrorResponse(w, "invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response.ErrorResponse(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userId := claims["userId"].(string)
		rolestring := claims["role"].(string)

		role := role.Role(rolestring)

		ctx := context.WithValue(r.Context(), contextkey.UserId, userId)
		ctx = context.WithValue(ctx, contextkey.UserRole, role)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/models/contextkey"
	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/models/roles"
	// role "github.com/Yash-Watchguard/MovieTicketBooking/internal/models/roles"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/response"
	"github.com/Yash-Watchguard/MovieTicketBooking/utills"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authoriaztionHeader:=r.Header.Get("Authorization")

        if(authoriaztionHeader==""){
			response.ErrorResponse(w,"Authorization token in not present",http.StatusBadRequest,1000)
			return
		}
     
		// we know that in the beareer token first part is bearer
		if !strings.HasPrefix(authoriaztionHeader,"Bearer "){
			response.ErrorResponse(w,"invalid token",http.StatusBadRequest,1000)
			return
		}

		tokenString:=strings.TrimPrefix(authoriaztionHeader,"Bearer ")

		token,err:=utills.VarifyJwt(tokenString)
		if err!=nil{
			response.ErrorResponse(w,"invalid token",http.StatusUnauthorized,1000)
			return
		}

		claims,ok:=token.Claims.(jwt.MapClaims)

		if !ok ||!token.Valid{
             response.ErrorResponse(w,  "Invalid token",http.StatusUnauthorized, 1002)
			return
		}

		userId:=claims["userId"].(string)
		rolestring:=claims["role"].(string)

		role:=role.Role(rolestring)

		ctx:=context.WithValue(r.Context(),contextkey.UserId,userId)
		ctx=context.WithValue(ctx,contextkey.UserRole,role)

		next.ServeHTTP(w,r.WithContext(ctx))
	})
}
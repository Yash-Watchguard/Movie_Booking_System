package utills

import (
	"testing"
	"time"

	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/roles"
	"github.com/golang-jwt/jwt/v5"
)

// unit test for jwt
func TestGenerateJwt(t *testing.T) {
	userId := "12345"
	userRole := role.Admin

	tokenString, err := GenerateJwt(userId, userRole)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})

	if err != nil {
		t.Fatalf("Error parsing token: %v", err)
	}

	if !token.Valid {
		t.Fatalf("Token is not valid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		t.Fatalf("Expected MapClaims, got %T", token.Claims)
	}

	if claims["userId"] != userId {
		t.Errorf("Expected userId %v, got %v", userId, claims["userId"])
	}

	if claims["role"] != string(userRole) {
		t.Errorf("Expected role %v, got %v", userRole, claims["role"])
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		t.Fatalf("Expected exp claim to be float64, got %T", claims["exp"])
	}

	expirationTime := time.Unix(int64(exp), 0)
	if time.Until(expirationTime) < time.Hour*23 || time.Until(expirationTime) > time.Hour*25 {
		t.Errorf("Expiration time is not approximately 24 hours from now")
	}
}

func TestVarifyJwt_ValidToken(t *testing.T) {
	userId := "12345"
	userRole := role.Customer

	tokenString, err := GenerateJwt(userId, userRole)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	token, err := VarifyJwt(tokenString)
	if err != nil {
		t.Fatalf("Expected no error verifying token, got %v", err)
	}

	if !token.Valid {
		t.Errorf("Expected token to be valid")
	}
}

func TestVarifyJwt_InvalidToken(t *testing.T) {
	invalidToken := "invalid.token.string"

	token, err := VarifyJwt(invalidToken)
	if err == nil {
		t.Errorf("Expected error verifying invalid token, got none")
	}

	if token != nil {
		t.Errorf("Expected token to be nil for invalid token")
	}
}

func TestVarifyJwt_ExpiredToken(t *testing.T) {
	// Create a token with expired time
	claims := jwt.MapClaims{
		"userId": "12345",
		"role":   string(role.Admin),
		"exp":    time.Now().Add(-time.Hour).Unix(), // Expired 1 hour ago
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		t.Fatalf("Error signing token: %v", err)
	}

	_, err = VarifyJwt(tokenString)
	if err == nil {
		t.Errorf("Expected error for expired token, got none")
	}
}

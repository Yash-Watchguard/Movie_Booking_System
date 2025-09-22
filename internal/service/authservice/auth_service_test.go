package authservice

import (
	"errors"
	"testing"

	repomock "github.com/Yash-Watchguard/MovieTicketBooking/internal/mocks/repoaitorymock"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	//role "github.com/Yash-Watchguard/MovieTicketBooking/internal/models/roles"

	"golang.org/x/crypto/bcrypt"
)

// unit testing for auth service
func TestSignUp(t *testing.T) {
	tests := []struct {
		name           string
		email          string
		password       string
		phoneNumber    string
		existingUser   bool
		shouldFailSave bool
		wantError      error
	}{
		{
			name:         "User already exists",
			email:        "yashgoyal@gmail.com",
			password:     "sjdjndj",
			phoneNumber:  "9999999999",
			existingUser: true,
			wantError:    errors.New("user with this email is already available"),
		},
		{
			name:           "Save user fails",
			email:          "sjdsa",
			password:       "passsndsad",
			phoneNumber:    "9876543210",
			shouldFailSave: true,
			wantError:      errors.New("error to save user"),
		},
		{
			name:        "Successful signup",
			email:       "sjdjsad",
			password:    "pjsdsjs",
			phoneNumber: "1112223333",
			wantError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := repomock.NewMockUserRepo()
			mockUserRepo.ShouldFailSave = tt.shouldFailSave

			if tt.existingUser {
				mockUserRepo.UserByEmail[tt.email] = &model.User{}
			}

			authService := NewAuthService(mockUserRepo)

			_, err := authService.SignUp(tt.name, tt.email, tt.phoneNumber, tt.password)

			if tt.wantError == nil {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
			} else {
				if err == nil || err.Error() != tt.wantError.Error() {
					t.Errorf("expected error %v, got %v", tt.wantError, err)
				}
			}
		})
	}
}

func TestLogin(t *testing.T) {
	tests := []struct {
		name        string
		email       string
		password    string
		isUserExist bool
		wrongPass   bool
		wantError   error
	}{
		{
			name:        "User does not exist",
			email:       "nouser@example.com",
			password:    "password",
			isUserExist: false,
			wantError:   errors.New("invalid credentials"),
		},
		{
			name:        "Password mismatch",
			email:       "test@example.com",
			password:    "wrongpassword",
			isUserExist: true,
			wrongPass:   true,
			wantError:   errors.New("invalid credentials"),
		},
		{
			name:        "Successful login",
			email:       "test@example.com",
			password:    "correctpassword",
			isUserExist: true,
			wantError:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepo := repomock.NewMockUserRepo()

			if tt.isUserExist {
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("correctpassword"), bcrypt.DefaultCost)

				user := &model.User{
					Id:       "user-123",
					Email:    tt.email,
					Password: string(hashedPassword),
					Role:     "user",
				}
				mockUserRepo.UserByEmail[tt.email] = user
			}

			authService := NewAuthService(mockUserRepo)

			_, token, err := authService.Login(tt.name, tt.email, tt.password)

			if tt.wantError == nil {
				if err != nil {
					t.Errorf("expected no error, got %v", err)
				}
				if token == "" {
					t.Errorf("expected jwt token, got empty")
				}
			} else {
				if err == nil || err.Error() != tt.wantError.Error() {
					t.Errorf("expected error %v, got %v", tt.wantError, err)
				}
			}
		})
	}
}

// Extra test for JWT failure (if utills.GenerateJwt is mockable)



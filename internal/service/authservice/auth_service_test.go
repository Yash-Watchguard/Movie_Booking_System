package authservice

import (
	"errors"
	"testing"

	repomock "github.com/Yash-Watchguard/MovieTicketBooking/internal/mocks/repoaitorymock"
	// model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

func TestSignUp(t *testing.T) {
    tests := []struct {
        name           string
        email          string
        password       string
        phoneNumber    string
        existingUser   bool
        shouldFailSave bool
        wantError error
	   
    }{
        {
            name:         "yahs",
            email:        "yashgoyal@gmail.com",
            password:     "sjdjndj",
            phoneNumber:  "sdjsjd",
            existingUser: true,
            wantError:    errors.New("user with this email is already available"),
			// ExpectedError: errors.New("user with this email is already available"),
        },
        {
            name:           "Save user fails",
            email:          "sjdsa",
            password:       "passsndsad",
            phoneNumber:    "9876543210",
            shouldFailSave: true,
            wantError:      errors.New("error to save user"),
			// ExpectedError: errors.New("error to save user"),
        },
        {
            name:        "Successful signup",
            email:       "sjdjsad",
            password:    "pjsdsjs",
            phoneNumber: "1112223333",
            wantError:   nil,
			// ExpectedError: nil,
        },
    }

    for _, tt := range tests {
        
            mockUserRepo := repomock.NewMockUserRepo()
            mockUserRepo.ShouldFailSave = tt.shouldFailSave

            if tt.existingUser {
                mockUserRepo.UserByEmail[tt.email] = nil
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
        
    }
}

func TestLogin(t *testing.T){
	tests:=[]struct{
       name string
	   email string
	   password string
	   IsUserExist bool
	   wantError error
	}{
		{
           name: "yash",
		   email: "yasb",
		   password: "sjd",
		   IsUserExist: false,
		   wantError: errors.New("invalid credentials"),
		   
		},
		
	}

	for _,test:=range tests{
		var err error
		mockUserrepo:=repomock.NewMockUserRepo()

		if test.IsUserExist{
			mockUserrepo.UserByEmail[test.email]=nil
		}
        
		AuthService:=NewAuthService(mockUserrepo)

		_,_,err=AuthService.Login(test.name,test.email,test.password)

		if err==nil ||err.Error()!=test.wantError.Error(){
            t.Errorf("jdndjcjdnc")
		}
	}
}

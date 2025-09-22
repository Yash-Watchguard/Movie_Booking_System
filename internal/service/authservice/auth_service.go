package authservice

import (
	"errors"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	userrepo "github.com/Yash-Watchguard/MovieTicketBooking/internal/repository/user_repo"
	"github.com/Yash-Watchguard/MovieTicketBooking/utills"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo userrepo.UserRepoInterface
}

func NewAuthService(userRepo userrepo.UserRepoInterface) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (authService *AuthService) SignUp(name string, userEmail string, phoneNumber string, Password string) (string, error) {

	_, err := authService.userRepo.GetUserByEmail(userEmail)
	if err == nil {
		return "", errors.New("user with this email is already available")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing password")
	}

	userId := utills.GenerateUuid()
	err = authService.userRepo.SaveUser(userId, name, userEmail, phoneNumber, string(hashedPassword))
	if err != nil {
		return "", err
	}

	return userId, nil
}

func (authService *AuthService) Login(name string, email string, password string) (*model.User, string, error) {

	// check user is present or not
	user, err := authService.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	jwtToken, err := utills.GenerateJwt(user.Id, user.Role)
	if err != nil {
		return nil, "", errors.New("error in generating jwt token")

	}
	return user, jwtToken, nil
}

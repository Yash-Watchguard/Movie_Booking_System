package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Yash-Watchguard/MovieTicketBooking/internal/response"
	"github.com/Yash-Watchguard/MovieTicketBooking/internal/service/authservice"
	"github.com/Yash-Watchguard/MovieTicketBooking/utills"
	
)


type AuthHandler struct {
	authService authservice.AuthServiceInterface
}

func NewAuthHandler(authServive authservice.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService: authServive}
}

func (authHandler *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var err error
	//   first check for the invalid credientials
	type userData struct {
		Name  string `json:"name"`
		Email string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Password string `json:"password"`
	}
	var user userData
	err= json.NewDecoder(r.Body).Decode(&user)

	if err!=nil{
		response.ErrorResponse(w,"Invalid input",http.StatusBadRequest)
		return
	}

	// check user email
	err=utills.CheckEmail(user.Email)
	if err!=nil{
		response.ErrorResponse(w,"Invalid Email",http.StatusBadRequest)
		return
	}
	// check phone number
	err=utills.CheckPhoneNumber(user.PhoneNumber)
	if err!=nil{
		response.ErrorResponse(w,"Invalid phonenumber",http.StatusBadRequest)
		return
	}

	userId,err:=authHandler.authService.SignUp(user.Name,user.Email,user.PhoneNumber,user.Password)

	if err!=nil{
		response.ErrorResponse(w,"Signup Failed",http.StatusInternalServerError)
		return
	}
	// make the user dto 

	response.SuccessResponse(w,map[string]interface{}{"Id":userId},"User Created Successfully",http.StatusCreated)

}
func (authHandler *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    //   first get the input from the request
    var err error
	type UserData struct{
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
	}
	var user UserData

	err=json.NewDecoder(r.Body).Decode(&user)
	if err!=nil{
		response.ErrorResponse(w,"Inavalid input",http.StatusBadRequest)
		return
	}

	err=utills.CheckEmail(user.Email)
	if err!=nil{
		response.ErrorResponse(w,"Invalid Email",http.StatusBadRequest)
		return
	}
	
	NewUser,JwtToken,err:=authHandler.authService.Login(user.Name,user.Email,user.Password)
    if err!=nil{
		response.ErrorResponse(w,"login failed",http.StatusInternalServerError)
		return
	}

	response.SuccessResponse(w,map[string]any{"token":JwtToken,"UserId":NewUser.Id},"Token Generted Successfully",http.StatusCreated)

}

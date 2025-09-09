package userrepo

import (
	"database/sql"
	"errors"
	"sync"

	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/models/roles"
)

type UserRepo struct {
	db *sql.DB
	mu *sync.Mutex
}
func NewUserRepo(db *sql.DB)*UserRepo{
	return &UserRepo{db: db}
}

func(userRepo *UserRepo)SaveUser(userId,name,email,phoneNumber,password string)(error){
	query:=`INSERT INTO users(user_id, name, email, phone_number, password, role) VALUES(?, ?, ?, ?, ?, ?)`

	userRepo.mu.Lock()
	_,err:=userRepo.db.Exec(query, userId, name, email, phoneNumber, password, role.Customer)
    userRepo.mu.Unlock()

	if err!=nil{
        // Improved error handling
        return errors.New("failed to save user")
	}
	return nil
}

func(userRepo *UserRepo)GetUserByEmail(email string)(*model.User,error){

	userRepo.mu.Lock()
	query:=`SELECT user_id, name, email, phone_number, password, role FROM users WHERE email = ?`
    userRepo.mu.Unlock()
	row:=userRepo.db.QueryRow(query,email)

	var user model.User

	err:=row.Scan(&user.Id,&user.Name,&user.Email,&user.PhoneNumber,&user.Password,&user.Role)
	if err!=nil{
		if err!=sql.ErrNoRows{
			return nil,errors.New("user not found")
		}
		return nil,err
	}
	return &user,nil
}

func(userRepo *UserRepo)GetUserById(userId string)(*model.User,error){
	query:=`SELECT user_id, name, email, phone_number, password, role FROM users where userid=?`
    
	userRepo.mu.Lock()
	row:=userRepo.db.QueryRow(query,userId)
	userRepo.mu.Unlock()

	var user model.User

	err:=row.Scan(&user.Id,&user.Name,&user.Email,&user.PhoneNumber,&user.Role)
	if err!=nil{
		if err!=sql.ErrNoRows{
			return nil,errors.New("user not found")
		}
		return nil,err
	}
	return &user,nil
}
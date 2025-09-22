package userrepo

import (
	"database/sql"
	"errors"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
	role "github.com/Yash-Watchguard/MovieTicketBooking/internal/constants/roles"
)

type UserRepo struct {
	db *sql.DB
}
func NewUserRepo(db *sql.DB)*UserRepo{
	return &UserRepo{db: db}
}

func(userRepo *UserRepo)SaveUser(userId,name,email,phoneNumber,password string)(error){
	query:=`INSERT INTO users(user_id, name, email, phone_number, password, role) VALUES(?, ?, ?, ?, ?, ?)`

	_,err:=userRepo.db.Exec(query, userId, name, email, phoneNumber, password, role.Customer)

	if err!=nil{
        return errors.New("failed to save user")
	}
	return nil
}

func(userRepo *UserRepo)GetUserByEmail(email string)(*model.User,error){
	query:=`SELECT user_id, name, email, phone_number, password, role FROM users WHERE email = ?`

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
	query:=`SELECT user_id, name, email, phone_number, password, role FROM users where user_id=?`
	row:=userRepo.db.QueryRow(query,userId)

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

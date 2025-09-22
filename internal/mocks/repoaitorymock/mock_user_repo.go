package repomock

import (
	"errors"
	model "github.com/Yash-Watchguard/MovieTicketBooking/internal/models"
)

// SaveUser(userId,name,email,phoneNumber,password string)error
// 	GetUserById(userId string) (*model.User,error)
// 	GetUserByEmail(email string)(*model.User,error)


type MockUserRepo struct {
    UserById     map[string]*model.User
    UserByEmail  map[string]*model.User
    ShouldFailSave bool
}



func NewMockUserRepo()*MockUserRepo{
  return &MockUserRepo{UserById: make(map[string]*model.User),
	UserByEmail:make(map[string]*model.User),
}
}

func (mockUserRepo *MockUserRepo) SaveUser(userId, name, email, phoneNumber, password string) error {
    if mockUserRepo.ShouldFailSave {
        return errors.New("error to save user")
    }
    mockUserRepo.UserById[userId] = nil
    mockUserRepo.UserByEmail[email] = mockUserRepo.UserById[userId]
    return nil
}


func(mockUserRepo * MockUserRepo)GetUserById(userId string) (*model.User,error){
	user,exits:=mockUserRepo.UserById[userId]

	if !exits{
		return nil,nil
	}
	return user,nil
}
func(mockUserRepo *MockUserRepo)GetUserByEmail(email string)(*model.User,error){
    user,exits:=mockUserRepo.UserByEmail[email]

	if !exits{
		return nil,errors.New("user is not present")
	}
	return user,nil
}




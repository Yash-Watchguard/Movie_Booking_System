package utills

import (
	"errors"
	"regexp"
	"strings"
)

func CheckPhoneNumber(number string) error {
	re := regexp.MustCompile(`^[6-9]\d{9}$`)

	if !re.MatchString(number){
		return errors.New("invalid phone number: must be 10 digits and start with 6-9")
	}
	return nil
}

func CheckEmail(email string)error{
	Email:=strings.ToLower(email)
	re:=regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)

	if !re.MatchString(Email){
		 return errors.New("invalid email address")
	}
	return nil
}

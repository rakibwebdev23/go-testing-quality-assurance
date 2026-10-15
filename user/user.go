package user

import (
	"errors"
	"strings"
)

func CheckUserName(username string) bool {
	if len(username) < 6 || strings.Contains(username, "admin") {
		return false
	}
	return true
}

func Login (username string, password string)(error, bool){
	if !CheckUserName(username){
		return errors.New("invalid username"),false
	}
	return nil,true
}
package entity

import (
	"mywallet/application/apperror"
	"strings"
)

type User struct {
    ID string `` //
}

type UserRequest struct {
	Name string
}

func NewUser(req UserRequest) (*User, error) {

	if req.Name == "" {
		return nil, apperror.UserNameMustNotEmpty
	}

	id := strings.ToLower(strings.ReplaceAll(req.Name, " ", ""))

	var obj User
	obj.ID = id

	return &obj, nil
}


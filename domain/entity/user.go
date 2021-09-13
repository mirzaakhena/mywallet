package entity

import (
  "mywallet/application/apperror"
)

type User struct {
  ID   string `` //
  Name string
}

type UserRequest struct {
  ID   string `` //
  Name string
}

func NewUser(req UserRequest) (*User, error) {

  if req.Name == "" {
    return nil, apperror.UserNameMustNotEmpty
  }

  var obj User
  obj.Name = req.Name
  obj.ID = req.ID

  return &obj, nil
}

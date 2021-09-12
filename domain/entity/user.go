package entity

import (
  "mywallet/application/apperror"
)

type User struct {
  ID   string `` //
  Name string
}

func NewUser(name string) (*User, error) {

  if name == "" {
    return nil, apperror.UserNameMustNotEmpty
  }

  var obj User
  obj.Name = name

  return &obj, nil
}

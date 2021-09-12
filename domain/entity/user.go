package entity

import (
  "mywallet/application/apperror"
)

type User struct {
  ID   string `` //
  Name string
}


func NewUser(id, name string) (*User, error) {

  if id == "" {
    return nil, apperror.UserIDMustNotEmpty
  }

  if name == "" {
    return nil, apperror.UserNameMustNotEmpty
  }

  var obj User
  obj.ID = id
  obj.Name = name

  return &obj, nil
}

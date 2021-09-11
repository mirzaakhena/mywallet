package vo

import (
	"mywallet/application/apperror"
	"strings"
)

type UserTeamRole string

const (
	AdminUserTeamRoleEnum  UserTeamRole = "ADMIN"
	MemberUserTeamRoleEnum UserTeamRole = "MEMBER"
)

var enumUserTeamRole = map[UserTeamRole]UserTeamRoleDetail{
	AdminUserTeamRoleEnum:  {},
	MemberUserTeamRoleEnum: {},
}

type UserTeamRoleDetail struct { //
}

func NewUserTeamRole(name string) (UserTeamRole, error) {
	name = strings.ToUpper(name)

	if _, exist := enumUserTeamRole[UserTeamRole(name)]; !exist {
		return "", apperror.UnrecognizedEnum.Var(name, "UserTeamRole")
	}

	return UserTeamRole(name), nil
}

func (r UserTeamRole) GetDetail() UserTeamRoleDetail {
	return enumUserTeamRole[r]
}

func (r UserTeamRole) PossibleValues() []UserTeamRole {
	res := make([]UserTeamRole, len(enumUserTeamRole))
	for key := range enumUserTeamRole {
		res = append(res, key)
	}
	return res
}

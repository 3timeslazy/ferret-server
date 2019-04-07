package auth

import (
	"github.com/MontFerret/ferret-server/pkg/common/dal"
)

type UserEntity struct {
	dal.Entity
	User
}

type User struct {
	Credentials
	Role string `json:"role"`
}

type Credentials struct {
	Name     string `json:"name"`
	Password []byte `json:"password"`
}

package auth

import (
	"context"

	"github.com/MontFerret/ferret-server/pkg/common/dal"
)

type Repository interface {
	GetByUsername(ctx context.Context, username string) (*UserEntity, error)

	CreateUser(ctx context.Context, user User) (dal.Entity, error)
}

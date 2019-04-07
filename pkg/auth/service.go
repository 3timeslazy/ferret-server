package auth

import (
	"context"
	"time"

	"github.com/MontFerret/ferret-server/pkg/common"
	"github.com/MontFerret/ferret-server/pkg/common/dal"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type (
	DBContext interface {
		GetAuthRepository() (Repository, error)
	}

	Service struct {
		db DBContext
	}
)

func NewService(db DBContext) (*Service, error) {
	if db == nil {
		return nil, errors.Wrap(common.ErrMissedArgument, "db context")
	}

	return &Service{db}, nil
}

func (service *Service) VerifyCredentials(ctx context.Context, cred Credentials) (*User, error) {
	repo, err := service.db.GetAuthRepository()
	if err != nil {
		return nil, errors.Wrap(err, "resolve auth repository")
	}

	entity, err := repo.GetByUsername(ctx, cred.Name)
	if err != nil {
		return nil, errors.Wrap(err, "get by username")
	}

	err = bcrypt.CompareHashAndPassword(entity.Password, cred.Password)
	if err != nil && err != bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.Wrap(err, "check hash")
	}

	// on this step err can be `nil` or `bcrypt.ErrMismatchedHashAndPassword`
	isVerified := err != bcrypt.ErrMismatchedHashAndPassword

	if !isVerified {
		return nil, nil
	}

	return &entity.User, nil
}

func (service *Service) BuildToken(user *User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role":     user.Role,
		"username": user.Name,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	})

	sk, err := service.signingKey()
	if err != nil {
		return "", errors.Wrap(err, "get signing key")
	}

	signed, err := token.SignedString(sk)
	if err != nil {
		return "", errors.Wrap(err, "sign token")
	}

	return signed, nil
}

func (service *Service) ParseToken(token string) (*jwt.Token, error) {
	sk, err := service.signingKey()
	if err != nil {
		return nil, errors.Wrap(err, "get signing key")
	}

	parsed, err := jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) {
		if _, ok := tkn.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return sk, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "parse token")
	}

	if parsed == nil || (parsed != nil && !parsed.Valid) {
		return nil, ErrInvalidToken
	}

	return parsed, nil
}

func (service *Service) GetByUsername(ctx context.Context, username string) (*UserEntity, error) {
	repo, err := service.db.GetAuthRepository()
	if err != nil {
		return nil, errors.Wrap(err, "resolve auth repository")
	}

	return repo.GetByUsername(ctx, username)
}

func (service *Service) CreateUser(ctx context.Context, user User) (dal.Entity, error) {
	repo, err := service.db.GetAuthRepository()
	if err != nil {
		return dal.Entity{}, errors.Wrap(err, "resolve auth repository")
	}

	hashed, err := service.hash(user.Password)
	if err != nil {
		return dal.Entity{}, errors.Wrap(err, "hash password")
	}

	user.Password = hashed

	return repo.CreateUser(ctx, user)
}

func (service *Service) hash(bs []byte) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword(bs, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hashed, nil
}

func (service *Service) signingKey() ([]byte, error) {
	return []byte("secret"), nil
}

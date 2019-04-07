package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/MontFerret/ferret-server/pkg/auth"
	"github.com/MontFerret/ferret-server/pkg/common"
	"github.com/MontFerret/ferret-server/pkg/common/dal"
	"github.com/MontFerret/ferret-server/server/db/repositories/queries"

	"github.com/arangodb/go-driver"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
)

type (
	userRecord struct {
		Key string `json:"_key"`
		dal.Metadata
		auth.User
	}

	AuthRepository struct {
		client     driver.Client
		collection driver.Collection
	}
)

func NewAuthRepository(client driver.Client, db driver.Database, collectionName string) (*AuthRepository, error) {
	ctx := context.Background()

	collection, err := initCollection(ctx, db, collectionName)

	if err != nil {
		return nil, err
	}

	return &AuthRepository{client, collection}, nil
}

func (repo *AuthRepository) GetByUsername(ctx context.Context, username string) (*auth.UserEntity, error) {
	cursor, err := repo.collection.Database().Query(ctx,
		fmt.Sprintf(queries.FindOneByName, repo.collection.Name()),
		map[string]interface{}{
			"name": username,
		},
	)

	if err != nil {
		return nil, err
	}

	defer cursor.Close()

	user := new(auth.UserEntity)

	if cursor.HasMore() {
		rec := userRecord{}

		meta, err := cursor.ReadDocument(ctx, &rec)
		if err != nil {
			return nil, err
		}

		user = fromUser(meta, rec)
	}

	return user, nil
}

func (repo *AuthRepository) CreateUser(ctx context.Context, user auth.User) (dal.Entity, error) {
	if user.Name == "" {
		return dal.Entity{}, common.Error(common.ErrInvalidArgument, "empty name")
	}

	cursor, err := repo.collection.Database().Query(
		ctx,
		fmt.Sprintf(queries.FindOneByName, repo.collection.Name()),
		map[string]interface{}{
			"name": user.Name,
		},
	)
	if err != nil {
		return dal.Entity{}, err
	}

	defer cursor.Close()

	if cursor.HasMore() {
		return dal.Entity{}, errors.Wrapf(common.ErrNotUnique, "user name '%s'", user.Name)
	}

	id, err := uuid.NewV4()
	if err != nil {
		return dal.Entity{}, errors.Wrap(err, "new id")
	}

	createdAt := time.Now()

	rec := userRecord{
		Key: id.String(),
		Metadata: dal.Metadata{
			CreatedAt: createdAt,
		},
		User: user,
	}

	meta, err := repo.collection.CreateDocument(ctx, rec)
	if err != nil {
		return dal.Entity{}, errors.Wrap(err, "create user")
	}

	return createdEntity(meta, createdAt), nil
}

func fromUser(meta driver.DocumentMeta, record userRecord) *auth.UserEntity {
	return &auth.UserEntity{
		Entity: dal.Entity{
			ID:       meta.Key,
			Rev:      meta.Rev,
			Metadata: record.Metadata,
		},
		User: auth.User{
			Credentials: auth.Credentials{
				Name:     record.Name,
				Password: record.Password,
			},
			Role: record.Role,
		},
	}
}

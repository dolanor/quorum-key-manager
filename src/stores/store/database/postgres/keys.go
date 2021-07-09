package postgres

import (
	"context"

	"github.com/consensys/quorum-key-manager/src/infra/log"
	postgres2 "github.com/consensys/quorum-key-manager/src/infra/postgres"

	"github.com/consensys/quorum-key-manager/src/stores/store/database"
	"github.com/consensys/quorum-key-manager/src/stores/store/entities"

	"github.com/consensys/quorum-key-manager/pkg/errors"
)

type Keys struct {
	logger log.Logger
	db     postgres2.Client
}

var _ database.Keys = &Keys{}

func NewKeys(logger log.Logger, db postgres2.Client) *Keys {
	return &Keys{
		logger: logger,
		db:     db,
	}
}

func (d *Keys) Get(_ context.Context, id string) (*entities.Key, error) {
	return nil, errors.ErrNotImplemented
}

func (d *Keys) GetDeleted(_ context.Context, id string) (*entities.Key, error) {
	return nil, errors.ErrNotImplemented

}

func (d *Keys) GetAll(_ context.Context) ([]*entities.Key, error) {
	return nil, errors.ErrNotImplemented

}

func (d *Keys) GetAllDeleted(_ context.Context) ([]*entities.Key, error) {
	return nil, errors.ErrNotImplemented

}

func (d *Keys) Add(_ context.Context, key *entities.Key) error {
	return errors.ErrNotImplemented
}

func (d *Keys) Update(_ context.Context, key *entities.Key) error {
	return errors.ErrNotImplemented

}

func (d *Keys) AddDeleted(_ context.Context, key *entities.Key) error {
	return errors.ErrNotImplemented

}

func (d *Keys) Remove(_ context.Context, id string) error {
	return errors.ErrNotImplemented
}

func (d *Keys) RemoveDeleted(_ context.Context, id string) error {
	return errors.ErrNotImplemented
}

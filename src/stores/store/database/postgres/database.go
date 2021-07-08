package postgres

import (
	"github.com/consensys/quorum-key-manager/pkg/log"
	"github.com/consensys/quorum-key-manager/src/stores/infra/postgres"
	"github.com/consensys/quorum-key-manager/src/stores/store/database"
)

type Database struct {
	logger       log.Logger
	db           postgres.Client
	eth1Accounts *ETH1Accounts
	keys         *Keys
}

var _ database.Database = &Database{}

func New(logger log.Logger, db postgres.Client) *Database {
	return &Database{
		logger:       logger,
		db:           db,
		eth1Accounts: NewETH1Accounts(logger), // TODO: Implement ETH1Accounts using Postgres and not in-memory
		keys:         NewKeys(logger, db),
	}
}

func (db *Database) ETH1Accounts() database.ETH1Accounts {
	return db.eth1Accounts
}

func (db *Database) Keys() database.Keys {
	return db.keys
}

package svc

import (
	"context"

	"github.com/jmoiron/sqlx"

	"go.uber.org/zap"
)

var _ Store = (*PostgresStore)(nil)

type PostgresStore struct {
	logger *zap.Logger
	db     *sqlx.DB
}

func NewPostgresStore(logger *zap.Logger, db *sqlx.DB) *PostgresStore {
	return &PostgresStore{
		logger: logger,
		db:     db,
	}
}

func (p *PostgresStore) Create(ctx context.Context, user *StoreUserModel) error {
	q := `INSERT INTO users (
	id,
	name,
	email_id,
	phone_number
	) VALUES (
	:id,
	:name,
	:email_id,
	:phone_number) `
	_, err := p.db.NamedExecContext(ctx, q, user)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresStore) Update(ctx context.Context, user *StoreUserModel) error {
	panic("implement me")
}

func (p *PostgresStore) Delete(ctx context.Context, user *StoreUserModel) error {
	panic("implement me")
}

func (p *PostgresStore) Get(ctx context.Context, id string) (*StoreUserModel, error) {
	panic("implement me")
}

package svc

import (
	"context"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

var _ Store = (*MemStore)(nil)

type MemStore struct {
	logger  *zap.Logger
	userMap map[string]string
}

func NewMemStore(logger *zap.Logger) *MemStore {
	return &MemStore{
		logger:  logger,
		userMap: map[string]string{},
	}
}

func (m *MemStore) Create(ctx context.Context, user *StoreUserModel) error {
	m.userMap["id"] = user.ID.String()
	m.userMap["name"] = user.Name
	m.userMap["email"] = user.EmailID
	m.userMap["phone"] = user.PhoneNumber
	return nil
}

func (m *MemStore) Update(ctx context.Context, user *StoreUserModel) error {
	panic("implement me")
}

func (m *MemStore) Delete(ctx context.Context, user *StoreUserModel) error {
	panic("implement me")
}

func (m *MemStore) Get(ctx context.Context, id string) (*StoreUserModel, error) {

	user := &StoreUserModel{
		ID:          uuid.MustParse(m.userMap["id"]),
		Name:        m.userMap["name"],
		EmailID:     m.userMap["email"],
		PhoneNumber: m.userMap["phone"],
	}
	return user, nil
}

package svc

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"go.uber.org/zap"
)

type UserService interface {
	CreateUser(ctx context.Context, p Payload) error
	UpdateUser(ctx context.Context, p Payload) error
	GetUser(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) error
}

var _ UserService = (*Service)(nil)

type Service struct {
	logger *zap.Logger
	store  Store
}

func NewService(logger *zap.Logger, store Store) *Service {
	return &Service{
		logger: logger,
		store:  store,
	}
}

func (s *Service) CreateUser(ctx context.Context, p Payload) error {
	if validateUser(p.Name) {
		return errors.New("empty user name")
	}
	if validateUser(p.EmailID) {
		return errors.New("empty email id")
	}
	userID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	user := StoreUserModel{
		ID:          userID,
		Name:        p.Name,
		EmailID:     p.EmailID,
		PhoneNumber: p.PhoneNumber,
	}
	s.logger.Info("user", zap.Any("user", user))
	if err := s.store.Create(ctx, &user); err != nil {
		s.logger.Error("unable to create user to db", zap.Error(err))
		return err
	}
	return nil
}

func (s *Service) UpdateUser(ctx context.Context, p Payload) error {
	panic("implement me")
}

func (s *Service) GetUser(ctx context.Context, id string) (*User, error) {

	_, err := uuid.Parse(id)
	if err != nil {
		s.logger.Error("error parsing uuid", zap.Error(err))
		return nil, err
	}
	model, _ := s.store.Get(ctx, id)
	u := &User{
		ID:          model.ID,
		Name:        model.Name,
		EmailID:     model.EmailID,
		PhoneNumber: model.PhoneNumber,
	}
	return u, nil
}

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	panic("implement me")
}

func validateUser(str string) bool {
	return len(str) == 0
}

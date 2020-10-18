package svc

import "context"

type Store interface {
	Create(ctx context.Context, user *StoreUserModel) error
	Update(ctx context.Context, user *StoreUserModel) error
	Delete(ctx context.Context, user *StoreUserModel) error
	Get(ctx context.Context, id string) (*StoreUserModel, error)
}

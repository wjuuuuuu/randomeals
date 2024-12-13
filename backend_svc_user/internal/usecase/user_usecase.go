package usecase

import (
	"backend_svc_user/internal/adapters/repository"
	"backend_svc_user/internal/entitiy"
	"context"
)

// 비지니스 로직!
type UserUseCase struct {
	userRepo repository.UserRepository
}

func NewUseCase(repo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: repo}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, user *entitiy.User) error {
	return nil
}

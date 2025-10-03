package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/you/myapp-backend/domain"
	"github.com/you/myapp-backend/repositories"
)

type UserUsecase interface {
	GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

type userUsecaseImpl struct {
	userRepository repositories.UserRepository
}

func NewUserUsecase(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepository: userRepository}
}

func (u *userUsecaseImpl) GetUserById(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := u.userRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, domain.ErrUserNotFound
	}
	return user, nil
}

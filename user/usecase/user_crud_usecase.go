package usecase

import (
	"go-tinder/domain"
	"context"
)

type userUsecase struct {
	userRepository domain.UserRepository
	
}

func NewUserUsecase(ur domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (u *userUsecase) GetByID(ctx context.Context, id int64) (domain.User, error) {
	res, err := u.userRepository.GetByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return res, nil
}

func (u *userUsecase) Store(ctx context.Context, user *domain.User) error {
	return u.userRepository.Store(ctx, user)
}

func (u *userUsecase) Delete(ctx context.Context, id int64) error {
	return u.userRepository.Delete(ctx, id)
}
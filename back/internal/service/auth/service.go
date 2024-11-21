package auth

import (
	"context"
	"errors"

	"github.com/pttrulez/activitypeople/internal/domain"
	"github.com/pttrulez/activitypeople/internal/infra/store"
	"github.com/pttrulez/activitypeople/internal/service"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, email, password string) (domain.User, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword),
		[]byte(password))
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (s *Service) Register(ctx context.Context, email, password, name string) (
	domain.User, error) {
	_, err := s.userRepo.GetByEmail(ctx, email)
	if !errors.Is(err, store.ErrNotFound) {
		return domain.User{}, service.ErrAlreadyExists
	}

	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	u, err := s.userRepo.Insert(ctx, email, string(encpw), name)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func NewService(userRepo UserRepo) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

type UserRepo interface {
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Insert(ctx context.Context, email, hashedPassword, name string) (domain.User, error)
}

type Service struct {
	userRepo UserRepo
}

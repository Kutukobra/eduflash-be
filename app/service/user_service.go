package service

import (
	"context"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/Kutukobra/eduflash-be/app/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repository.UserRepository
	roomRepo repository.RoomRepository
}

func NewUserService(userRepo repository.UserRepository, roomRepo repository.RoomRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		roomRepo: roomRepo,
	}
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	userData, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (s *UserService) GetRoomsByOwnerId(ctx context.Context, owner_id string) ([]model.Room, error) {
	rooms, err := s.roomRepo.GetRoomsByOwnerId(ctx, owner_id)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (s *UserService) RegisterUser(ctx context.Context, email string, username string, password string) (*model.User, error) {
	if err := ValidateRole(role); err != nil {
		return nil, err
	}

	userData, err := s.userRepo.RegisterUser(ctx, username, email, password, role)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (s *UserService) LoginUser(ctx context.Context, email string, password string) (*model.User, error) {
	userData, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	userData.Password = "" // clean so no exposed hash

	return userData, nil
}

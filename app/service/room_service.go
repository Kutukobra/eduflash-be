package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/Kutukobra/eduflash-be/app/repository"
)

type RoomService struct {
	room_count int32
	repo       repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) *RoomService {
	return &RoomService{
		repo:       repo,
		room_count: 0,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, owner_id string) (*model.Room, error) {
	id_string := fmt.Sprintf("%06d", s.room_count)
	s.room_count++

	room, err := s.repo.CreateRoom(ctx, id_string, owner_id)
	if err != nil {
		return nil, err
	}
	s.room_count += 7
	return room, nil
}

func (s *RoomService) JoinRoom(ctx context.Context, id string) error {
	roomData, err := s.repo.GetRoomById(ctx, id)
	if err != nil {
		return err
	}

	log.Println(roomData.ID)
	if roomData == nil {
		return errors.New("invalid room")
	}

	return nil
}

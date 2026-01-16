package service

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/Kutukobra/eduflash-be/app/repository"
	"github.com/jackc/pgx/v5/pgconn"
)

type RoomService struct {
	repo repository.RoomRepository
}

func NewRoomService(repo repository.RoomRepository) *RoomService {
	return &RoomService{repo: repo}
}

func (s *RoomService) CreateRoom(ctx context.Context, roomName string, ownerId string) (*model.Room, error) {
	for {
		id := fmt.Sprintf("%06d", rand.Intn(1000000))

		room, err := s.repo.CreateRoom(ctx, id, roomName, ownerId)
		if err == nil {
			return room, nil
		}

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			continue
		}

		return nil, err
	}
}

func (s *RoomService) GetRoomById(ctx context.Context, room_id string) (*model.Room, error) {
	roomData, err := s.repo.GetRoomById(ctx, room_id)
	if err != nil {
		return nil, err
	}
	return roomData, nil
}

func (s *RoomService) JoinRoom(ctx context.Context, roomId string, studentName string) (*model.Room, error) {
	roomData, err := s.repo.GetRoomById(ctx, roomId)
	if err != nil {
		return nil, err
	}

	err = s.repo.JoinRoom(ctx, roomId, studentName)
	if err != nil {
		return nil, err
	}

	return roomData, nil
}

func (s *RoomService) GetStudentsByRoomId(ctx context.Context, roomId string) ([]string, error) {
	students, err := s.repo.GetStudentsByRoomId(ctx, roomId)
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (s *RoomService) GetQuizzesByRoomId(ctx context.Context, roomId string) ([]string, error) {
	rooms, err := s.repo.GetQuizzesByRoomId(ctx, roomId)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

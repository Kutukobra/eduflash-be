package service

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/Kutukobra/eduflash-be/app/repository"
)

type RoomService struct {
	room_count int32
	repo       repository.RoomRepository
	checkId    map[int]bool
}

func NewRoomService(repo repository.RoomRepository) *RoomService {
	return &RoomService{
		repo:       repo,
		room_count: 0,
		checkId:    make(map[int]bool),
	}
}

func (s *RoomService) generateRoomId() string {
	r := rand.New(rand.NewSource(int64(s.room_count)))

	var room_id int
	for {
		room_id = r.Intn(999999)
		if !s.checkId[room_id] {
			s.checkId[room_id] = true
			break
		}
	}

	id_string := fmt.Sprintf("%06d", room_id)
	s.room_count++
	return id_string
}

func (s *RoomService) CreateRoom(ctx context.Context, roomName string, owner_id string) (*model.Room, error) {
	id_string := s.generateRoomId()

	room, err := s.repo.CreateRoom(ctx, id_string, roomName, owner_id)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (s *RoomService) JoinRoom(ctx context.Context, room_id string, student_name string) (*model.RoomStudent, error) {
	_, err := s.repo.GetRoomById(ctx, room_id)
	if err != nil {
		return nil, err
	}

	roomData, err := s.repo.JoinRoom(ctx, room_id, student_name)
	if err != nil {
		return nil, err
	}

	return roomData, nil
}

func (s *RoomService) GetStudentsByRoomId(ctx context.Context, room_id string) ([]string, error) {
	students, err := s.repo.GetStudentsByRoomId(ctx, room_id)
	if err != nil {
		return nil, err
	}

	return students, nil
}

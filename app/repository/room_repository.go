package repository

import (
	"context"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/jackc/pgx/v5"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, id string, roomName string, ownerId string) (*model.Room, error)
	GetRoomById(ctx context.Context, id string) (*model.Room, error)
	GetRoomsByOwnerId(ctx context.Context, ownerId string) ([]model.Room, error)
	JoinRoom(ctx context.Context, room_id string, student_name string) error
	GetStudentsByRoomId(ctx context.Context, room_id string) ([]string, error)
	AddQuiz(ctx context.Context, roomId string, quizId string) error
}

type PGRoomRepository struct {
	driver *pgx.Conn
}

func rowToRoom(row pgx.Row) (*model.Room, error) {
	var room model.Room
	err := row.Scan(&room.ID, &room.RoomName, &room.CreatedAt, &room.OwnerId)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func NewPGRoomRepository(driver *pgx.Conn) *PGRoomRepository {
	return &PGRoomRepository{driver: driver}
}

func (r *PGRoomRepository) CreateRoom(
	ctx context.Context,
	id string, roomName string, ownerId string,
) (*model.Room, error) {
	query := `
		INSERT INTO rooms (id, room_name, owner_id)
		VALUES ($1, $2, $3) RETURNING id, room_name, created_at, owner_id`

	row := r.driver.QueryRow(
		ctx, query, id, roomName, ownerId,
	)
	return rowToRoom(row)
}

func (r *PGRoomRepository) GetRoomById(ctx context.Context, id string) (*model.Room, error) {
	query := `SELECT * FROM rooms WHERE id = $1`

	row := r.driver.QueryRow(ctx, query, id)

	return rowToRoom(row)
}

func (r *PGRoomRepository) GetRoomsByOwnerId(ctx context.Context, ownerId string) ([]model.Room, error) {
	query := `SELECT id, room_name, owner_id FROM Rooms WHERE owner_id = $1 ORDER BY created_at DESC`

	rows, err := r.driver.Query(ctx, query, ownerId)
	if err != nil {
		return nil, err
	}

	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		if err := rows.Scan(&room.ID, &room.RoomName, &room.OwnerId); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *PGRoomRepository) JoinRoom(ctx context.Context, roomId string, studentName string) error {
	query := `INSERT INTO Room_Student (Room_Id, Student_Name) VALUES ($1, $2)`

	_, err := r.driver.Exec(
		ctx, query, roomId, studentName,
	)

	return err
}

func (r *PGRoomRepository) GetStudentsByRoomId(ctx context.Context, roomId string) ([]string, error) {
	query := `SELECT student_name FROM Room_Student WHERE room_id = $1`

	rows, err := r.driver.Query(ctx, query, roomId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		students = append(students, name)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}

func (r *PGRoomRepository) AddQuiz(ctx context.Context, roomId string, quizId string) error {
	query := `INSERT INTO room_quiz (room_id, quiz_id) VALUES ($1, $2)`

	_, err := r.driver.Exec(ctx, query, roomId, quizId)
	if err != nil {
		return err
	}

	return nil
}

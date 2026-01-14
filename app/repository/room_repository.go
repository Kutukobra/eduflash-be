package repository

import (
	"context"
	"log"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/jackc/pgx/v5"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, id string, owner_id string) (*model.Room, error)
	GetRoomById(ctx context.Context, id string) (*model.Room, error)
	GetRoomsByOwnerId(ctx context.Context, owner_id string) ([]model.Room, error)
	JoinRoom(ctx context.Context, room_id string, student_name string) (*model.RoomStudent, error)
	GetStudentsByRoomId(ctx context.Context, room_id string) ([]string, error)
}

type PGRoomRepository struct {
	driver *pgx.Conn
}

func rowToRoom(row pgx.Row) (*model.Room, error) {
	var room model.Room
	err := row.Scan(&room.ID, &room.Owner_ID)
	if err != nil {
		return nil, err
	}

	return &room, nil
}

func rowToRoomStudent(row pgx.Row) (*model.RoomStudent, error) {
	var roomStudent model.RoomStudent
	err := row.Scan(&roomStudent.Room_ID, &roomStudent.Student_Name)
	if err != nil {
		return nil, err
	}

	return &roomStudent, nil
}

func NewPGRoomRepository(driver *pgx.Conn) *PGRoomRepository {
	return &PGRoomRepository{driver: driver}
}

func (r *PGRoomRepository) CreateRoom(
	ctx context.Context,
	id string, owner_id string,
) (*model.Room, error) {
	query := `
		INSERT INTO Rooms (Id, Owner_id)
		VALUES ($1, $2) RETURNING Id, Owner_Id`

	row := r.driver.QueryRow(
		ctx, query, id, owner_id,
	)
	return rowToRoom(row)
}

func (r *PGRoomRepository) GetRoomById(ctx context.Context, id string) (*model.Room, error) {
	query := `SELECT * FROM Rooms WHERE Id = $1`

	row := r.driver.QueryRow(ctx, query, id)

	return rowToRoom(row)
}

func (r *PGRoomRepository) GetRoomsByOwnerId(ctx context.Context, owner_id string) ([]model.Room, error) {
	log.Println(owner_id)

	query := `SELECT id, owner_id FROM Rooms WHERE Owner_id = $1`

	rows, err := r.driver.Query(ctx, query, owner_id)
	if err != nil {
		return nil, err
	}

	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		if err := rows.Scan(&room.ID, &room.Owner_ID); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r *PGRoomRepository) JoinRoom(ctx context.Context, room_id string, student_name string) (*model.RoomStudent, error) {
	query := `
		INSERT INTO Room_Student (Room_Id, Student_Name) VALUES ($1, $2)
		RETURNING Room_Id, Student_Name`

	row := r.driver.QueryRow(
		ctx, query, room_id, student_name,
	)

	return rowToRoomStudent(row)
}

func (r *PGRoomRepository) GetStudentsByRoomId(ctx context.Context, room_id string) ([]string, error) {
	query := `SELECT student_name FROM Room_Student WHERE room_id = $1`

	rows, err := r.driver.Query(ctx, query, room_id)
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

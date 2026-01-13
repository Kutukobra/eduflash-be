package repository

import (
	"context"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/jackc/pgx/v5"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, id string, owner_id string) (*model.Room, error)
	GetRoomById(ctx context.Context, id string) (*model.Room, error)
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

	row := r.driver.QueryRow(ctx, query)

	return rowToRoom(row)
}

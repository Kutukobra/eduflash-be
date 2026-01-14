package app

import (
	"github.com/Kutukobra/eduflash-be/app/config"
	"github.com/Kutukobra/eduflash-be/app/database"
	"github.com/Kutukobra/eduflash-be/app/handler"
	"github.com/Kutukobra/eduflash-be/app/repository"
	"github.com/Kutukobra/eduflash-be/app/service"
)

type App struct {
	userHandler *handler.UserHandler
	roomHandler *handler.RoomHandler
}

func New(cfg *config.Config) (*App, error) {
	PGDatabase, err := database.NewPostgresDatabase(cfg.PGString)
	if err != nil {
		return nil, err
	}

	userRepository := repository.NewPGUserRepository(PGDatabase)
	roomRepository := repository.NewPGRoomRepository(PGDatabase)

	userService := service.NewUserService(userRepository, roomRepository)
	roomService := service.NewRoomService(roomRepository)

	userHandler := handler.NewUserHandler(userService)
	roomHandler := handler.NewRoomHandler(roomService)

	return &App{
		userHandler: userHandler,
		roomHandler: roomHandler,
	}, nil
}

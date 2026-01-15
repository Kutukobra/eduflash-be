package service

import (
	"context"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/Kutukobra/eduflash-be/app/repository"
)

type QuizService struct {
	quizRepo repository.QuizRepository
	roomRepo repository.RoomRepository
}

func NewQuizService(quizRepo repository.QuizRepository, roomRepo repository.RoomRepository) *QuizService {
	return &QuizService{quizRepo: quizRepo, roomRepo: roomRepo}
}

func (s *QuizService) CreateQuiz(ctx context.Context, roomId string, quiz []model.QuizContent) (string, error) {
	id, err := s.quizRepo.CreateQuiz(ctx, quiz)
	if err != nil {
		return "", err
	}

	err = s.roomRepo.AddQuiz(ctx, roomId, id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (s *QuizService) GetQuizById(ctx context.Context, id string) ([]model.QuizContent, error) {
	return s.quizRepo.GetQuizById(ctx, id)
}

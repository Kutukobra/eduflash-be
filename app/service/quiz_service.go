package service

import (
	"context"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/Kutukobra/eduflash-be/app/repository"
)

type QuizService struct {
	repo repository.QuizRepository
}

func NewQuizService(repo repository.QuizRepository) *QuizService {
	return &QuizService{repo: repo}
}

func (s *QuizService) CreateQuiz(ctx context.Context, quiz []model.QuizContent) (string, error) {
	id, err := s.repo.CreateQuiz(ctx, quiz)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s *QuizService) GetQuizById(ctx context.Context, id string) ([]model.QuizContent, error) {
	return s.repo.GetQuizById(ctx, id)
}

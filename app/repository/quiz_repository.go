package repository

import (
	"context"
	"encoding/json"

	"github.com/Kutukobra/eduflash-be/app/model"
	"github.com/jackc/pgx/v5"
)

type QuizRepository interface {
	CreateQuiz(ctx context.Context, quiz []model.QuizContent) (string, error)
	GetQuizById(ctx context.Context, id string) ([]model.QuizContent, error)
	SubmitScore(ctx context.Context, quizId string, studentName string, score float32) error
	GetQuizScores(ctx context.Context, quizId string) ([]model.StudentScores, error)
}

type PGQuizRepository struct {
	driver *pgx.Conn
}

func NewPGQuiRepository(driver *pgx.Conn) *PGQuizRepository {
	return &PGQuizRepository{driver: driver}
}

func (r *PGQuizRepository) CreateQuiz(
	ctx context.Context, quiz []model.QuizContent,
) (string, error) {
	data, err := json.Marshal(quiz)
	if err != nil {
		return "", err
	}

	var resultId string
	query := `INSERT INTO quizzes (content) VALUES ($1) RETURNING id`
	err = r.driver.QueryRow(ctx, query, data).Scan(&resultId)

	if err != nil {
		return "", err
	}

	return resultId, nil
}

func (r *PGQuizRepository) GetQuizById(ctx context.Context, id string) ([]model.QuizContent, error) {
	var raw []byte

	query := `SELECT content FROM quizzes WHERE id = $1`
	err := r.driver.QueryRow(ctx, query, id).Scan(&raw)
	if err != nil {
		return nil, err
	}

	var quiz []model.QuizContent
	err = json.Unmarshal(raw, &quiz)

	return quiz, err
}

func (r *PGQuizRepository) SubmitScore(ctx context.Context, quizId string, studentName string, score float32) error {
	query := `
		INSERT INTO student_scores (quiz_id, student_name, score)
		VALUES ($1, $2, $3)`

	_, err := r.driver.Exec(ctx, query, quizId, studentName, score)

	return err
}

func (r *PGQuizRepository) GetQuizScores(ctx context.Context, quizId string) ([]model.StudentScores, error) {
	query := `SELECT student_name, score FROM student_scores WHERE quiz_id = $1`

	rows, err := r.driver.Query(ctx, query, quizId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []model.StudentScores
	for rows.Next() {
		var score model.StudentScores
		if err := rows.Scan(&score.StudentName, &score.Score); err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return scores, nil
}

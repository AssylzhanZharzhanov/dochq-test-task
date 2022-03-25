package grpc

import (
	"context"
	"fmt"
	"github.com/AssylzhanZharzhanov/dochq-test-task/models"
	"github.com/jmoiron/sqlx"
)

const (
	answersTable = "answers"
)

type AnswerRepository struct {
	db *sqlx.DB
}

func (r *AnswerRepository) CreateAnswer(ctx context.Context, answer models.Answer) (models.Answer, error) {
	output := models.Answer{}

	query := fmt.Sprintf(`
		INSERT INTO %s (key, value) 
		VALUES ($1, $2)
		RETURNING key, value
	`, answersTable)

	row := r.db.QueryRowContext(ctx, query, answer.Key, answer.Val)

	if err := row.Scan(&output.Key, &output.Val); err != nil {
		return output, err
	}

	return output, nil
}

func (r *AnswerRepository) GetAnswer(ctx context.Context, key string) (models.Answer, error) {
	output := models.Answer{}

	query := fmt.Sprintf(`
		SELECT 
			key, 
			value 
		FROM %s
		WHERE key = $1
		ORDER BY id DESC
		LIMIT 1
	`, answersTable)

	err := r.db.GetContext(ctx, &output, query, key)
	return output, err
}

func (r *AnswerRepository) UpdateAnswer(ctx context.Context, key string, value string) (models.Answer, error) {
	output := models.Answer{}

	query := fmt.Sprintf(`
		INSERT INTO %s (key, value) 
		VALUES ($1, $2)
		RETURNING key, value
	`, answersTable)

	row := r.db.QueryRowContext(ctx, query, key, value)

	if err := row.Scan(&output.Key, &output.Val); err != nil {
		return output, err
	}

	return output, nil
}

func (r *AnswerRepository) DeleteAnswer(ctx context.Context, key string) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE key = $1
	`, answersTable)

	_, err := r.db.ExecContext(ctx, query, key)
	return err
}

func NewAnswerRepository(db *sqlx.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}

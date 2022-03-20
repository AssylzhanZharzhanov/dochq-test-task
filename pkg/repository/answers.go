package repository

import (
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

func (r *AnswerRepository) CreateAnswer(answer models.Answer) (models.Answer, error) {
	output := models.Answer{}

	query := fmt.Sprintf(`
		INSERT INTO %s (key, value) 
		VALUES ($1, $2)
		RETURNING key, value
	`, answersTable)

	row := r.db.QueryRow(query, answer.Key, answer.Val)

	if err := row.Scan(&output.Key, &output.Val); err != nil {
		return output, err
	}

	return output, nil
}

func (r *AnswerRepository) GetAnswer(key string) (models.Answer, error) {
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

	err := r.db.Get(&output, query, key)
	return output, err
}

func (r *AnswerRepository) UpdateAnswer(key string, value string) (models.Answer, error) {
	output := models.Answer{}

	query := fmt.Sprintf(`
		INSERT INTO %s (key, value) 
		VALUES ($1, $2)
		RETURNING key, value
	`, answersTable)

	row := r.db.QueryRow(query, key, value)

	if err := row.Scan(&output.Key, &output.Val); err != nil {
		return output, err
	}

	//query := fmt.Sprintf(`
	//	UPDATE %s
	//	SET value = $2
	//	WHERE key = $1
	//	RETURNING key, value
	//`, answersTable)
	//
	//row := r.db.QueryRow(query, key, value)
	//if err := row.Scan(&output.Key, &output.Val); err != nil {
	//	return output, err
	//}

	return output, nil
}

func (r *AnswerRepository) DeleteAnswer(key string) error {
	query := fmt.Sprintf(`
		DELETE FROM %s
		WHERE key = $1
	`, answersTable)

	_, err := r.db.Exec(query, key)
	return err
}

func NewAnswerRepository(db *sqlx.DB) *AnswerRepository {
	return &AnswerRepository{db: db}
}
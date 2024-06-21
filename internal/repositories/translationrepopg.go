package repositories

import (
	"context"
	"fmt"
	"leonardodelira/go-clean-template/internal/core/domain"
	"leonardodelira/go-clean-template/internal/core/ports"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepo struct {
	conn    *pgxpool.Pool
	timeout time.Duration
}

func NewTranslationPGRepository(conn *pgxpool.Pool) ports.TranslationRepository {
	envTimeout, _ := strconv.Atoi(os.Getenv("PG_TIMEOUT"))
	timeout := time.Duration(envTimeout) * time.Second

	return &postgresRepo{
		conn:    conn,
		timeout: timeout,
	}
}

func (r *postgresRepo) SaveTranslation(c context.Context, translation *domain.Translation) (int32, error) {
	ctx, cancel := context.WithTimeout(c, r.timeout)
	defer cancel()

	sql := "INSERT INTO history (origin_text, translation, language_destination, language_origin_detected) VALUES ($1, $2, $3, $4) RETURNING id"

	var row int32
	err := r.conn.QueryRow(ctx, sql, translation.Text, translation.Translation,
		translation.LanguageDestination, translation.LanguageOriginDetected).Scan(&row)

	if err != nil {
		return -1, fmt.Errorf("TranslationRepo - SaveTranslation - Query: %w", err)
	}

	return row, nil
}

func (r *postgresRepo) GetTranslations(c context.Context) ([]domain.Translation, error) {
	ctx, cancel := context.WithTimeout(c, r.timeout)
	defer cancel()

	sql := "SELECT id, origin_text, translation, language_destination, language_origin_detected FROM history"

	rows, err := r.conn.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetTranslation - Query: %w", err)
	}

	translations := make([]domain.Translation, 0)
	for rows.Next() {
		tl := domain.Translation{}

		err = rows.Scan(&tl.ID, &tl.Text, &tl.Translation, &tl.LanguageDestination, &tl.LanguageOriginDetected)
		if err != nil {
			return nil, fmt.Errorf("TranslationRepo - GetTranslation - rows.Scan: %w", err)
		}

		translations = append(translations, tl)
	}

	return translations, nil
}

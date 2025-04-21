package utils

import (
	"finhub-go/internal/core/errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func ToUUID(str string) (uuid.UUID, error) {
	if str == "" {
		return uuid.UUID{}, errors.ErrEmptyField
	}

	parsedUUID, err := uuid.Parse(str)
	if err != nil {
		return uuid.UUID{}, err
	}

	return parsedUUID, nil
}

func ToUUIDSlice(strs []string) []uuid.UUID {
	var result []uuid.UUID
	for _, s := range strs {
		if id, err := uuid.Parse(s); err == nil {
			result = append(result, id)
		}
	}
	return result
}

func ToNillableUUID(str string) (*uuid.UUID, error) {
	if str == "" {
		return nil, nil
	}

	parsedUUID, err := uuid.Parse(str)
	if err != nil {
		return nil, err
	}

	return &parsedUUID, nil
}

func ToDateTime(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, errors.ErrEmptyField
	}

	formats := []string{
		time.RFC3339,
		"2006-01-02",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return t, nil
		}
	}

	// TODO: usar um error mais padrao
	return time.Time{}, fmt.Errorf("formato de data inválido: %s", dateStr)
}

// TODO: rever as funçoes para usar ponteiro e formato q ela usa, depois de fazer o front sugiu mudanças
func ToDateUnsafe(dateStr *string) *time.Time {
	if dateStr == nil || *dateStr == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02", *dateStr)
	if err != nil {
		return nil
	}
	return &t
}

func ToNillableDateTime(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}
	formats := []string{
		time.RFC3339,
		"2006-01-02",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, dateStr); err == nil {
			return &t, nil
		}
	}
	// TODO: usar um error mais padrao
	return nil, fmt.Errorf("formato de data inválido: %s", dateStr)
}

func ToDateTimeString(date time.Time) string {
	return date.Format("2006-01-02 15:04:05")
}

// TODO: rever o seu uso
func SafeToNillableDateTimeString(date *time.Time) *string {
	if date == nil || date.IsZero() {
		return nil
	}
	formatted := date.Format("2006-01-02")
	return &formatted
}

package utils

import (
	"finhub-go/internal/core/errors"
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

func ToDate(dateStr string) (time.Time, error) {
	if dateStr == "" {
		return time.Time{}, errors.ErrEmptyField

	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// TODO: rever as funçoes para usar ponteiro
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

func ToNillableDate(dateStr string) (*time.Time, error) {
	if dateStr == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, err
	}
	return &t, nil
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

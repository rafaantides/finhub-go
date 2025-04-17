package notifier

import "context"

type Notifier interface {
	SendMessage(ctx context.Context, content string) error
	NotifyError(ctx context.Context, location string, err error) error
	NotifyImportResult(ctx context.Context, jobID string, filename string, successCount int, failCount int) error
}

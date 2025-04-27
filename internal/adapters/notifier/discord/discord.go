package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"finhub-go/internal/core/ports/outbound/cachestorage"
	"finhub-go/internal/utils/logger"
	"fmt"
	"net/http"
)

type Discord struct {
	webhookURL string
	cch        cachestorage.CacheStorage
	log        *logger.Logger
}

func NewDiscord(cch cachestorage.CacheStorage, webhookURL string) *Discord {
	log := logger.NewLogger("Discord")
	if webhookURL != "" {
		log.Start("WebhookID: ...%s", webhookURL[len(webhookURL)-8:])
	}
	return &Discord{
		webhookURL: webhookURL,
		cch:        cch,
		log:        log,
	}
}

func (d *Discord) send(ctx context.Context, content string) error {
	key := "discord:rate_limit"

	// Check if rate limit is active
	if _, err := d.cch.Get(ctx, key); err == nil {
		d.log.Warn("Rate limit active. Skipping send.")
		return nil
	}

	payload := map[string]string{
		"content": content,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		d.log.Error("Failed to serialize payload: %v", err)
		return err
	}

	resp, err := http.Post(d.webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		if _, err := d.cch.Set(ctx, key, "1", 60); err != nil {
			d.log.Error("Failed to set rate limit in cache: %v", err)
		}
		d.log.Warn("Rate limit reached. Message skipped.")
		return nil
	}

	if resp.StatusCode >= 300 {
		d.log.Error("Failed to send message: %s", resp.Status)
		return fmt.Errorf("failed to send message, status: %s", resp.Status)
	}

	return nil
}

func (d *Discord) SendMessage(ctx context.Context, content string) error {
	return d.send(ctx, content)
}

func (d *Discord) NotifyError(ctx context.Context, location string, err error) error {
	msg := fmt.Sprintf("Error at `%s`\nDetails: ```%s```", location, err.Error())
	return d.send(ctx, msg)
}

func (d *Discord) NotifyImportResult(ctx context.Context, jobID string, filename string, successCount int, failCount int) error {
	msg := fmt.Sprintf("Import result: `%s`\nJob ID: `%s`\nSuccess: `%d`\nFailure: `%d`", filename, jobID, successCount, failCount)
	return d.send(ctx, msg)
}

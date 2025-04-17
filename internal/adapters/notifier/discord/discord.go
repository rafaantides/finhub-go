package discord

// TODO: rever esse package as vezes n precisa ter um adapter, e mudar os logs para ingles

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

	// Verifica se está em rate limit
	if _, err := d.cch.Get(ctx, key); err == nil {
		d.log.Warn("Rate limit ativo. Ignorando envio.")
		return nil
	}

	payload := map[string]string{
		"content": content,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		d.log.Error("Erro ao serializar payload: %v", err)
		return err
	}

	resp, err := http.Post(d.webhookURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		if _, err := d.cch.Set(ctx, key, "1", 60); err != nil {
			d.log.Error("falha ao setar rate limit no cache: %v", err)
		}
		d.log.Warn("Rate limit. Mensagem ignorada.")
		return nil
	}

	if resp.StatusCode >= 300 {
		d.log.Error("Falha ao enviar mensagem: %s", resp.Status)
		return fmt.Errorf("failed to send message, status: %s", resp.Status)
	}

	return nil
}

func (d *Discord) SendMessage(ctx context.Context, content string) error {
	return d.send(ctx, content)
}

func (d *Discord) NotifyError(ctx context.Context, location string, err error) error {
	msg := fmt.Sprintf("Erro em `%s`\nDetalhes: ```%s```", location, err.Error())
	return d.send(ctx, msg)
}

func (d *Discord) NotifyImportResult(ctx context.Context, jobID string, filename string, successCount int, failCount int) error {
	msg := fmt.Sprintf("Resultado da importação: `%s`\njob_id: `%s`\nSucesso: `%d`\nFalha: `%d`", filename, jobID, successCount, failCount)
	return d.send(ctx, msg)
}

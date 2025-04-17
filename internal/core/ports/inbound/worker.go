package inbound

import "finhub-go/internal/core/dto"

type MessageProcessor func(messageBody []byte, timeoutSeconds int) (*dto.NotifierMessage, error)

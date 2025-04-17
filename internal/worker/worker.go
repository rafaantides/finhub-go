package worker

import (
	"context"
	"finhub-go/internal/core/ports/inbound"
	"finhub-go/internal/core/ports/outbound/cachestorage"
	"finhub-go/internal/core/ports/outbound/messagebus"
	"finhub-go/internal/core/ports/outbound/notifier"
	"finhub-go/internal/utils/logger"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

type Worker struct {
	ctx            context.Context
	queue          string
	processor      inbound.MessageProcessor
	prefetchCount  int
	timeoutSeconds int

	log  *logger.Logger
	mbus messagebus.MessageBus
	cch  cachestorage.CacheStorage
	noti notifier.Notifier

	stopChan chan struct{} // Canal para sinalizar parada segura
	mu       sync.Mutex
}

func NewWorker(
	queue string,
	processor inbound.MessageProcessor,
	prefetchCount, timeoutSeconds int,
	log *logger.Logger,
	mbus messagebus.MessageBus,
	cch cachestorage.CacheStorage,
	noti notifier.Notifier,
) *Worker {
	ctx := context.Background()
	return &Worker{
		ctx:            ctx,
		queue:          queue,
		processor:      processor,
		prefetchCount:  prefetchCount,
		timeoutSeconds: timeoutSeconds,
		log:            log,
		mbus:           mbus,
		cch:            cch,
		noti:           noti,
		stopChan:       make(chan struct{}), // Inicializa o canal de parada
		mu:             sync.Mutex{},
	}
}
func (w *Worker) Start() {
	w.log.Start(
		"%s | Queue: %s | Concurrency: %d messages | Timeout: %ds",
		getFuncName(w.processor),
		w.queue,
		w.prefetchCount,
		w.timeoutSeconds,
	)

	msgs, err := w.mbus.ConsumeMessages(w.queue)
	if err != nil {
		w.log.Fatal("Error starting queue consumption %s: %v", w.queue, err)
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, w.prefetchCount) // Controla workers simultâneos

	for {
		select {
		case <-w.stopChan:
			w.log.Warn("Stop signal received. Shutting down worker...")
			goto cleanup

		case msg, ok := <-msgs:
			if !ok {
				w.log.Warn("Message channel closed. Shutting down worker...")
				goto cleanup
			}

			wg.Add(1)
			semaphore <- struct{}{} // Bloqueia se já houver prefetchCount workers em execução

			go func(msg messagebus.Message) {
				defer wg.Done()
				defer func() { <-semaphore }() // Libera um slot ao final

				if err := w.processMessage(msg.Body()); err != nil {
					w.Stop() // Para o worker
				}

				if err := msg.Ack(); err != nil {
					w.log.Error("Failed to ACK message: %v", err)
				}
			}(msg)
		}
	}

cleanup:
	wg.Wait() // Aguarda todas as goroutines antes de encerrar
	close(semaphore)
	w.log.Success("Worker successfully stopped.")
}

func (w *Worker) processMessage(messageBody []byte) error {
	nmsg, err := w.processor(messageBody, w.timeoutSeconds)
	if err != nil {
		if nmsg != nil {
			if _, incrErr := w.cch.Incr(w.ctx, fmt.Sprintf("notify:%s:errors", nmsg.JobID)); incrErr != nil {
				w.log.Error("Failed to increment error counter in cache: %v", incrErr)
			}
		}

		w.log.Error("ctx: %s | %v", w.ctx, err)
		return err
	}

	if nmsg != nil {
		if _, incrOk := w.cch.Incr(w.ctx, fmt.Sprintf("notify:%s:success", nmsg.JobID)); incrOk != nil {
			w.log.Error("Failed to increment success counter in cache: %v", incrOk)
		}

		if nmsg.IsLastChunk {
			w.log.Info("Finished processing job_id=%s. Sending notification...", nmsg.JobID)

			// Recupera os contadores
			successCount := 0
			failCount := 0

			successStr, err := w.cch.Get(w.ctx, fmt.Sprintf("notify:%s:success", nmsg.JobID))
			if err == nil {
				successCount, _ = strconv.Atoi(successStr)
			}
			failStr, err := w.cch.Get(w.ctx, fmt.Sprintf("notify:%s:errors", nmsg.JobID))
			if err == nil {
				failCount, _ = strconv.Atoi(failStr)
			}

			if err := w.noti.NotifyImportResult(w.ctx, nmsg.JobID, nmsg.Filename, successCount, failCount); err != nil {
				w.log.Error("Failed to send notification for job_id=%s: %v", nmsg.JobID, err)
			}
		}
	}

	return nil
}

// Método para parar o worker com segurança
func (w *Worker) Stop() {
	w.mu.Lock()
	defer w.mu.Unlock()

	select {
	case <-w.stopChan:
		// Se já foi fechado, não faz nada
	default:
		close(w.stopChan)
	}
}

func getFuncName(i interface{}) string {
	ptr := reflect.ValueOf(i).Pointer()
	fn := runtime.FuncForPC(ptr)
	if fn == nil {
		return "unknown"
	}
	name := fn.Name()
	parts := strings.Split(name, ".")
	short := parts[len(parts)-1]
	return strings.TrimSuffix(short, "-fm")

}

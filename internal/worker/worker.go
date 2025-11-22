package worker

import (
	"context"
	"encoding/json"
	"fmt"

	queue "github.com/SunilKividor/donela/internal/queuq"
	"github.com/google/uuid"
)

type Worker struct {
	q queue.Queue
	p *Processor
}

func NewWorker(q queue.Queue, p *Processor) *Worker {
	return &Worker{
		q: q,
		p: p,
	}
}

func (w *Worker) Start(ctx context.Context) {
	fmt.Println("[WORKER] Started")

	for {
		select {
		case <-ctx.Done():
			fmt.Println("[WORKER] SHUTTING DOWN...")
			return
		default:
			w.processNext(ctx)
		}

	}
}

func (w *Worker) processNext(ctx context.Context) {
	msg, err := w.q.Receive(ctx)
	if err != nil {
		fmt.Println("[WORKER] error receiving :", err)
		return
	}

	if msg == nil {
		return
	}

	job, err := parseMessage(msg.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = w.p.Process(ctx, job)
	if err != nil {
		fmt.Println("[WORKER] error processing :", err)
	}
	w.q.Delete(ctx, msg)
}

type S3Event struct {
	Records []struct {
		S3 struct {
			Object struct {
				Key string `json:"key"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

func parseMessage(body []byte) (*Job, error) {
	var s3Event S3Event
	err := json.Unmarshal(body, &s3Event)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal s3 event: %w", err)
	}

	if len(s3Event.Records) == 0 {
		return nil, fmt.Errorf("no records found in s3 event")
	}

	objectKey := s3Event.Records[0].S3.Object.Key
	if objectKey == "" {
		return nil, fmt.Errorf("object key is empty")
	}

	job := &Job{
		Key:     objectKey,
		TrackID: uuid.NewString(),
	}

	return job, nil
}

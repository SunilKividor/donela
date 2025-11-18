package worker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type Worker struct {
	sqs      *sqs.Client
	queueURL string
}

func NewWorker(sqs *sqs.Client, queueURL string) *Worker {
	return &Worker{
		sqs:      sqs,
		queueURL: queueURL,
	}
}

func (w *Worker) Start(ctx context.Context) {

}

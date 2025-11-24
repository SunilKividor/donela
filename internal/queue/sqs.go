package queue

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSQueue struct {
	client   *sqs.Client
	queueURL string
}

func NewSQSQueue(client *sqs.Client, queueURL string) *SQSQueue {
	return &SQSQueue{
		client:   client,
		queueURL: queueURL,
	}
}

func (q *SQSQueue) Receive(ctx context.Context) (*Message, error) {
	resp, err := q.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &q.queueURL,
		MaxNumberOfMessages: 1,
		WaitTimeSeconds:     20,
		VisibilityTimeout:   60,
	})

	if err != nil {
		fmt.Println("[Queue] Error receiving message:", err)
		time.Sleep(3 * time.Second)
		return nil, err
	}

	if len(resp.Messages) == 0 {
		return nil, nil
	}

	msg := resp.Messages[0]

	return &Message{
		ID:            *msg.MessageId,
		Body:          []byte(*msg.Body),
		ReceiptHandle: msg.ReceiptHandle,
	}, nil
}

func (q *SQSQueue) Delete(ctx context.Context, msg *Message) error {
	_, err := q.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &q.queueURL,
		ReceiptHandle: msg.ReceiptHandle,
	})

	return err
}

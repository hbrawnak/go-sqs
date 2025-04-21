package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"os"
)

func receiveMessage(queueURL string) {
	region := os.Getenv("AWS_REGION")
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	svc := sqs.New(sess)

	for {
		output, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL),
			MaxNumberOfMessages: aws.Int64(5),
			WaitTimeSeconds:     aws.Int64(10),
		})

		if err != nil {
			log.Println("Error receiving messages:", err)
			continue
		}

		for _, message := range output.Messages {
			log.Printf("Received message: %s\n", *message.Body)

			svc.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueURL),
				ReceiptHandle: message.ReceiptHandle,
			})
		}
	}
}

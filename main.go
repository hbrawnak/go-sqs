package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	queueURL := os.Getenv("SQS_URL")

	go func() {
		var wg sync.WaitGroup

		for i := 1; i <= 10; i++ {
			wg.Add(1)

			go func(i int) {
				defer wg.Done()
				message := fmt.Sprintf("Concurrent message id: #%d", i)
				err := sendMessage(queueURL, message)
				if err != nil {
					log.Printf("Error sending message: %v", err)
				}

				fmt.Printf("Sent: %s\n", message)
			}(i)
		}

		wg.Wait()
	}()

	receiveMessage(queueURL)
}

#### GO-SQS

A simple project where I experimented with AWS SQS (Simple Queue Service) and Golang.
It demonstrates basic concepts like message production, consumption, concurrency, and error handling with SQS queues.

Additionally, Terraform is used to automate the creation of the SQS infrastructure on AWS.


### 📚 Features

* Send messages to SQS queues using a Producer.
* Receive and process messages using a Consumer.
* Concurrent message handling using Go's goroutines and WaitGroups.
* Retry logic for failed message sends.
* Dockerized services for easy setup and deployment.
* Terraform scripts for creating and managing the SQS queue.


#### 🛠️ Tech Stack

* Golang (Go 1.23.1)
* AWS SQS (Amazon Simple Queue Service)
* Terraform (for Infrastructure as Code)
* Docker & Docker Compose (for containerization)


### Project Structure
```
go-sqs/
├── consumer/
│   ├── main.go
│   └── consumer.go
│   └── Dockerfile
├── producer/
│   ├── main.go
│   └── producer.go
│   └── Dockerfile
├── terraform/
│   ├── main.tf
│   ├── variables.tf
│   └── outputs.tf
├── docker-compose.yml
├── go.mod
├── .env
└── README.md
```

### How to Run Locally

#### 1. Create SQS Queue (using Terraform)
```
cd terraform/
terraform init
terraform apply
```

#### 2. Prepare Environment Variables
Create a `.env` file in the root:
```
AWS_REGION=us-east-1
SQS_URL=https://sqs.us-east-1.amazonaws.com/your-account-id/your-queue-name
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
```

#### 3. Run Using Docker Compose
```
docker-compose up --build
```

#### 4. Manually Run (without Docker)
```
cd producer/
go run main.go producer.go
```
In another terminal
```
cd consumer/
go run main.go consumer.go
```

### Author
[Habibur Rahman](https://habib.im)

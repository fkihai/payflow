## Project Structure

```
payflow/
├── cmd/
│   ├── api/
│   │   └── main.go
│   └── worker/
│       └── main.go
│
├── config/
│   └── config.yaml
│
├── internal/
│   ├── domain/
│   │   ├── entity/
│   │   │   ├── transaction.go
│   │   │   ├── payment.go
│   │   │   └── device.go
│   │   │
│   │   └── repository/
│   │       ├── transaction_repository.go
│   │       └── payment_repository.go
│   │
│   ├── usecase/
│   │   ├── payment/
│   │   │   ├── create_qris.go
│   │   │   ├── confirm_payment.go
│   │   │   └── timeout_payment.go
│   │   │
│   │   └── device/
│   │       └── activate_device.go
│   │
│   ├── delivery/
│   │   ├── http/
│   │   │   ├── handler/
│   │   │   │   ├── payment_handler.go
│   │   │   │   └── webhook_handler.go
│   │   │   └── router.go
│   │   │
│   │   └── worker/
│   │       └── payment_worker.go
│   │
│   └── infrastructure/
│       ├── config/
│       │   └── loader.go
│       │
│       ├── db/
│       │   └── postgres/
│       │       └── transaction_repo.go
│       │
│       ├── mq/
│       │   └── rabbitmq/
│       │       └── publisher.go
│       │
│       └── payment_gateway/
│           └── mock/
│               └── client.go
│
├── pkg/
│   └── logger/
│       └── logger.go
│
├── migrations/
│   └── 001_create_transactions.sql
│
├── go.mod
└── go.sum

```

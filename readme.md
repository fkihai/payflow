## Project Structure

```
payflow/
├── .env                     # Environment variables (fill according to .env.example)
├── .env.example             # Template environment variables
├── go.mod                   # Go module file
├── go.sum                   # Go module dependencies checksum
│
├── cmd/                     # Entry points / main applications
│   ├── api/                 # HTTP API server
│   │   └── main.go
│   ├── worker/              # Background worker applications
│   │   └── main.go
│   └── migration/           # Database migration CLI
│       └── main.go
│
├── config/                  # Configuration files
│   └── config.yaml
│
├── internal/                # Core layers (Clean Architecture)
│   ├── domain/              # Business logic & entities
│   │   ├── entity/          # Domain models
│   │   ├── repository/      # Repository interfaces
│   │   ├── payment/         # Payment-specific domain logic
│   │   └── database/        # Database-specific logic
│   │
│   ├── usecase/             # Application/business rules
│   │   └── <feature>/       # Organized per feature (payment, device, etc.)
│   │
│   ├── delivery/            # Interface adapters
│   │   ├── httpx/           # HTTP handlers & routers (httpx framework)
│   │   └── worker/          # Background job adapters
│   │
│   └── infrastructure/      # External frameworks & services
│       ├── config/          # Config loader & helpers
│       ├── db/              # Database implementations
│       ├── mq/              # Message queue implementations
│       └── payment_gateway/ # External payment clients
│
├── pkg/                     # Shared libraries / helpers
│   └── logger/              # Logging utility
│
└── migrations/              # Database migration files


```

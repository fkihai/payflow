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
├── config/                  # Static configuration files
│   └── config.yaml
│
├── internal/                # Core layers (Clean Architecture)
│   ├── entity/              # Domain entities
│   │   └── *.go
│   │
│   ├── valueobject/         # Domain value objects
│   │   └── *.go
│   │
│   ├── usecase/             # Application / business rules
│   │   └── <feature>/       # Organized per feature (payment, device, etc.)
│   │       ├── interface.go
│   │       └── service.go
│   │
│   ├── delivery/            # Interface adapters
│   │   ├── http/            # HTTP handlers & routers
│   │   └── worker/          # Background job adapters
│   │
│   └── infrastructure/      # External frameworks & services
│       ├── config/          # Config loader & helpers
│       ├── db/              # Database implementations (repository)
│       ├── mq/              # Message queue implementations
│       └── gateway/         # External payment gateway clients
│
├── pkg/                     # Shared libraries (cross-layer safe)
│   └── logger/              # Logging utility
│
└── migrations/              # Database migration files
```

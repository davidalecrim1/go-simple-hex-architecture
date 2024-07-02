# Simple Hexagonal Architecture in Go

A simple sample project in Go implementing Hexagonal Architecture.

## Prerequisites

- [Docker](https://www.docker.com/)

## Getting Started

1. **Build the Project**

   To build the project, run:

   ```bash
   make build
   ```

2. **Run the Project Locally**

   To run the project locally, execute:

   ```bash
   make local_run
   ```

## Project Structure

```bash
.
├── cmd
│   └── app
└── internal
│   ├── domain.go
│   ├── repository.go
│   └── web.go
└── scripts
│   └── schema.sql
└── tests
    └── unit
```

- `cmd/app`: Initialization of the application.
- `internal/domain.go`: Domain logic of the application and repository interface.
- `internal/repository.go`: Repository implementation.
- `internal/web.go`: Web controller and handler functions for HTTP.
- `scripts/schema.sql`: Initialization scripts for SQLite.
- `tests/unit/`: Unit tests for the application.

## Makefile Targets

- `build`: Builds the application with Docker.
- `local_run`: Runs the application locally using Docker.
- `local_stop`: Stops and removes the Docker container.

## Contributing

Feel free to contribute to this project by submitting issues or pull requests. 

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
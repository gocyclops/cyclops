# Cyclops 🚀

A lightning-fast Go backend boilerplate that helps you ship production-ready APIs in minutes. Choose your favorite framework and add the features you need.

## Features ⚡

- **Multiple Framework Support**:
  - 🌟 Gin - High performance HTTP web framework
  - 🚀 Echo - High performance, minimalist web framework
  - ⚡ Fiber - Express-inspired web framework
  - 🦍 Gorilla Mux - Powerful HTTP routing and URL matcher

- **Core Features**:
  - 📦 GORM with PostgreSQL
  - 🔒 JWT Authentication
  - 🔄 Migration System
  - 📝 Structured Logging
  - ⚙️ Environment Configuration
  - 🐳 Docker Support
  - 🔄 Live Reload with Air

- **Optional Modules**:
  - 📨 Mail Service Integration
  - 💾 Redis Cache
  - 📁 AWS S3 Storage
  - 🔐 Authentication System

## Quick Start 🏃‍♂️

### Installation

```bash
go install github.com/gocyclops/cyclops@latest
cyclops
```

### Choose Your Stack

```bash
cd myproject
# Follow the interactive CLI prompts to:
# 1. Select your framework (Gin/Echo/Fiber/Gorilla)
# 2. Choose additional features
# 3. Configure your project
```

### Run Your Project

```bash
go mod tidy
go run main.go
```

## Project Structure 📁

```
├── .env                # Environment variables
├── .air.toml          # Live reload configuration
├── Dockerfile         # Container configuration
├── main.go            # Application entry point
├── controllers/       # Request handlers
├── models/           # Database models
├── repository/       # Data access layer
├── routes/           # API routes
├── database/         # Database configuration
├── migrations/       # Database migrations
└── utils/           # Helper functions
```

## Configuration ⚙️

1. Copy `.env.example` to `.env`
2. Update environment variables:
```env
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=myapp
DB_PORT=5432
```

## Development 💻

### Live Reload
```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with live reload
air
```

### Docker
```bash
# Build and run with Docker
docker-compose up
```

## Documentation 📚

- [Framework Guides](docs/frameworks.md)
  - [Gin](https://gin-gonic.com/docs/)
  - [Echo](https://echo.labstack.com/guide/)
  - [Fiber](https://docs.gofiber.io/)
  - [Gorilla](https://github.com/gorilla/mux#examples)

## Contributing 🤝

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing`)
5. Open a Pull Request

## Support 💖

- ⭐ Star this repo
- 🐛 Report bugs via [Issues](https://github.com/gocyclops/cyclops/issues)
- 💡 Request features via [Discussions](https://github.com/gocyclops/cyclops/discussions)

## License 📄

Copyright © 2024 - present [TeddyMuli](https://github.com/TeddyMuli).
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

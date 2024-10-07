Cyclops - Go Backend Boilerplate

# Cyclops 🚀

Cyclops is a boilerplate for building a backend application in Go. It includes a setup for database connections, Redis, S3, and routing with Fiber. This template is designed to help you quickly get started with your Go backend project.

## Features ✨

- **Fiber**: An Express-inspired web framework for Go.
- **GORM**: The fantastic ORM library for Golang.
- **Redis**: In-memory data structure store, used as a database, cache, and message broker.
- **S3**: Integration with AWS S3 for file storage.
- **SendGrid**: Email sending service.
- **Cron**: Job scheduling library.
- **Air**: Live reloading for Go applications.
- **JWT**: JSON Web Tokens for secure authentication.
- **Docker**: Containerization for consistent development and deployment environments.
- **GitHub Actions**: CI/CD workflows for automated testing and deployment.

## Getting Started 🚀

### Prerequisites 📋

- [Go 1.23 or higher](https://go.dev/dl/)
- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/) (optional, for running services like PostgreSQL and Redis)

### Installation 🛠️

Clone the repository:

```sh
git clone https://github.com/gocyclops/cyclops.git your_custom_directory --depth=1
cd your_custom_directory
rm -rf .git
git init
```

Replace placeholders with your project name:

```sh
find . -type f -exec sed -i 's/cyclops/yourprojectname/g' {} +
```

Set up environment variables:

```sh
cp .env.example .env
go mod tidy
```

Edit the `.env` file with your configuration:

```env
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_PORT=your_db_port
DB_HOST=your_db_host
REDIS_URL=your_redis_url
JWT_SECRET=your_jwt_secret
SECURE_COOKIE=true
MAIL_FROM=your_email@example.com
S3_URL=your_s3_url
S3_KEY_ID=your_s3_key_id
S3_SECRET_ACCESS_KEY=your_s3_secret_access_key
SENDGRID_API_KEY=your_sendgrid_api_key
ALLOWED_ORIGINS=http://localhost:8080
```

Run the application:

```sh
go run main.go
```

## Optionally with Docker 🐳

```
docker-compose up --build
```

## Documentation 📚

- [Fiber Documentation](https://docs.gofiber.io/)
- [GORM Documentation](https://gorm.io/docs/)
- [Redis Go Client Documentation](https://github.com/go-redis/redis)
- [AWS SDK for Go Documentation](https://aws.github.io/aws-sdk-go-v2/)
- [SendGrid Go Library Documentation](https://github.com/sendgrid/sendgrid-go)
- [Cron Documentation](https://github.com/robfig/cron)
- [Air Documentation](https://github.com/cosmtrek/air)
- [Docker Documentation](https://docs.docker.com/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
- [JWT Documentation](https://jwt.io/introduction/)

## Access the Application 🌐

Open your browser and navigate to [http://localhost:8080](http://localhost:8080).

## Using Air for Hot Reloading 🔄

Install Air:

```sh
go install github.com/cosmtrek/air@latest
```

Edit the `.air.toml` file in the root of your project.

Run Air:

```sh
air
```

## Contributing 🤝

If you want to say Thank You and/or support the active development of Cyclops:

Add a GitHub Star to the project: [Star on GitHub](https://github.com/gocyclops/cyclops)

Tweet about the project on your 𝕏 (Twitter): [Tweet about Cyclops](https://twitter.com/intent/tweet?text=Check%20out%20Cyclops%20-%20a%20boilerplate%20for%20building%20a%20backend%20application%20in%20Go!%20https://github.com/gocyclops/cyclops)

Write a review or tutorial on Medium, Dev.to or personal blog: [Write on Medium](https://medium.com/), [Write on Dev.to](https://dev.to/)

Support the project by donating a cup of coffee: [Buy Me a Coffee](https://www.buymeacoffee.com/gocyclops)

## License 📄

Copyright © 2024 - present [TeddyMuli](https://github.com/TeddyMuli) and [Contributors](https://github.com/gocyclops/cyclops/graphs/contributors).
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

# DCSS-Inspired Roguelike Game

A modern web-based roguelike game inspired by Dungeon Crawl Stone Soup (DCSS), built with Go, React, and Google Cloud Platform.

## Project Setup

This project uses Google Cloud Platform for infrastructure and GitHub for source control and CI/CD.

### Prerequisites

- [Google Cloud SDK](https://cloud.google.com/sdk/docs/install)
- [Go](https://golang.org/doc/install) (version 1.20 or later)
- [Docker](https://docs.docker.com/get-docker/)
- [Git](https://git-scm.com/downloads)

### Local Development

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/dcss-roguelike.git
   cd dcss-roguelike
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Run the server locally:
   ```
   go run cmd/server/main.go
   ```

4. Access the server at http://localhost:8080/health

### Building with Docker

Build the Docker image:
```
docker build -t roguelike-service .
```

Run the container:
```
docker run -p 8080:8080 roguelike-service
```

### Deployment

The project is automatically deployed to Google Cloud Run when changes are pushed to the main branch, using Cloud Build.

## Project Structure

```
.
├── cmd/
│   └── server/          # Server entry point
├── internal/            # Internal packages
├── pkg/                 # Public packages
├── web/                 # Frontend code (React)
├── cloudbuild.yaml      # Cloud Build configuration
├── Dockerfile           # Docker configuration
├── go.mod               # Go module definition
└── README.md            # This file
```

## Architecture

This project follows the architecture outlined in the Technical Architecture document, implementing:

- Go backend with microservices
- React frontend (coming soon)
- Entity Component System for game logic
- Google Cloud Platform for infrastructure

## License

[MIT License](LICENSE) # Testing Cloud Build integration

# Go Hello World Web Application

A simple Go web application demonstrating HTML templating, static file serving, and Kubernetes deployment. This project serves as a tutorial for building and deploying Go applications in a Kubernetes environment.

## Features

- Go web server with HTML template engine
- Static file serving (CSS, JS, etc.)
- Multiple page routing (Home and About pages)
- Health check endpoint (/ping)
- Docker containerization
- Kubernetes deployment configuration

## Project Structure

```plaintext
.
├── main.go
├── Dockerfile
├── templates/
│   ├── home.html
│   └── about.html
├── static/
│   └── css/
│       └── style.css
└── k8s/
    ├── deployment.yaml
    └── service.yaml
```

## Local Development

### Prerequisites

- Go 1.21 or later
- Docker (for containerization)
- kubectl (for Kubernetes deployment)

### Running Locally

1. Clone the repository
2. Run the application:
```bash
go run main.go
```
3. Visit http://localhost:8080 in your browser

### Building and Running with Docker

Build the Docker image:
```bash
docker build -t hello-world .
```

Run the container:
```bash
docker run -p 8080:8080 hello-world
```

## Kubernetes Deployment

### Prerequisites

- Kubernetes cluster
- kubectl configured to access your cluster
- Container registry access

### Deployment Steps

1. Push the Docker image to your container registry:
```bash
docker tag hello-world your-registry/hello-world:latest
docker push your-registry/hello-world:latest
```

2. Update the image reference in `k8s/deployment.yaml`

3. Apply the Kubernetes configurations:
```bash
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

4. Verify the deployment:
```bash
kubectl get pods
kubectl get services
```

## Application Endpoints

- `/` - Home page
- `/about` - About page
- `/ping` - Health check endpoint (returns JSON)

## Environment Configuration

The application uses environment variables for configuration. Create a `.env` file in the root directory with the following variables:

```env
APP_NAME=Go Hello World
APP_VERSION=1.0.0
APP_PORT=8080
APP_ENV=development
```

## Configuration

The application runs on port 8080 by default. The Kubernetes deployment includes:
- 2 replicas for high availability
- Resource limits and requests
- Health check probes
- Load balancing via Service

## Monitoring and Health Checks

The `/ping` endpoint can be used for:
- Kubernetes liveness probes
- Kubernetes readiness probes
- Load balancer health checks

## Contributing

This is a tutorial project. Feel free to fork and modify for your learning purposes.

## License

MIT License
```

This README provides:
1. Clear project overview
2. Setup instructions
3. Deployment steps
4. Project structure
5. Configuration details
6. Monitoring information
7. Development guidelines

Users can easily understand the project's purpose and how to deploy it in both local and Kubernetes environments.
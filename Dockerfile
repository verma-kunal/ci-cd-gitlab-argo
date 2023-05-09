FROM golang:1.20.3-bullseye

# Set the Working Directory inside the container
WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app files
COPY . .

# This container exposes port 3000 to the outside world
EXPOSE 3000

# Run the executable
CMD ["go", "run", "main.go"]
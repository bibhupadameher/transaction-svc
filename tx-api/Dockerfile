# Use official Go image
FROM golang:1.22-alpine

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN go build -o server .

# Expose API port
EXPOSE 8080


# Run
ENV APP_ENV=docker
CMD ["./server"]

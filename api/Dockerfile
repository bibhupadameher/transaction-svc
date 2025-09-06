# Use official Go image
FROM golang:1.22-alpine

WORKDIR /app

# Install dependencies
#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

# Copy source
COPY . .

# Build
RUN go build -o server .

# Expose API port
EXPOSE 8080

# Run
CMD ["./server"]

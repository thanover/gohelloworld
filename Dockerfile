FROM golang:1.24-alpine AS build

WORKDIR /app

# Copy go.mod and go.sum files (if they exist)
COPY go.mod go.sum* ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server .

# Use a smaller image for the final container
FROM alpine:latest

WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/server .

# Expose port 8080 for the application
EXPOSE 8080

# Run the application
CMD ["./server"]
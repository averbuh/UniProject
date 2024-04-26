# Use alpine linux as the base image (lightweight)
FROM golang:alpine AS builder

# Set working directory
WORKDIR /go/src/recipes

# Copy your Go source code (replace "app" with your project name)
COPY . .

# Install dependencies
RUN go mod download

# Build the Go binary (replace "main.go" with your main file)
RUN go build -o ./cmd/gin/main 

# Use a slimmer image for the final container
FROM alpine:latest

# Copy the built binary
COPY --from=builder /go/src/recipes/main /recipes

EXPOSE 8080
EXPOSE 443

# Set the entrypoint to run your application
ENTRYPOINT [ "/app/main" ]

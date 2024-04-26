# Use alpine linux as the base image (lightweight)
FROM golang:alpine AS builder

ARG CGO_ENABLED=0
# Set working directory
WORKDIR /recipes

COPY go.mod go.sum ./

# Install dependencies
RUN go mod download
COPY . .

EXPOSE 8080


# Build the Go binary (replace "main.go" with your main file)
RUN go build -o cmd/gin cmd/gin/main.go 

# Use a slimmer image for the final container
FROM alpine:latest

WORKDIR /
# Copy the built binary
COPY --from=builder recipes/cmd/gin/main /cmd/gin


EXPOSE 8080

# Set the entrypoint to run your application
ENTRYPOINT [ "/cmd/gin" ]
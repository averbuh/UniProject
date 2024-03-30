FROM golang:1.21.1-alpine

# Set destination for COPY
WORKDIR /app

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . ./

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# Build
RUN go build -o ./build/gin_api ./cmd/gin

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080
EXPOSE 443

# Run
CMD ["./build/gin_api"]

# golang image where workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/kobeld/golabs/docker/golang-docker

# Build the golang-docker command inside the container.
RUN go install github.com/kobeld/golabs/docker/golang-docker

# Run the golang-docker command when the container starts.
ENTRYPOINT /go/bin/golang-docker

# http server listens on port 8080.
EXPOSE 8080
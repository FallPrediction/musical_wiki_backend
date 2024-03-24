# Go Version
FROM golang:1.21.4-alpine

# Environment variables which CompileDaemon requires to run
ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

# Basic setup of the container
RUN mkdir /app
COPY .. /app
WORKDIR /app

# Get CompileDaemon
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon
RUN go install github.com/go-delve/delve/cmd/dlv@master

RUN mkdir /.aws

# The build flag sets how to build after a change has been detected in the source code
# The command flag sets how to run the app after it has been built
ENTRYPOINT CompileDaemon -build="go build -o musical_wiki" -command="sh ./deployments/local/kill_and_debug.sh"

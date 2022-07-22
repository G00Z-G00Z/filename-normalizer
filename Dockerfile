FROM golang:1.18 as dev

WORKDIR /app

# Installs go utils for testing
RUN go install -v golang.org/x/tools/cmd/godoc@latest

# Installs the cobra-cli
RUN go install github.com/spf13/cobra-cli@latest


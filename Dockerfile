############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

ARG PB_VERSION=0.19.0

RUN apk add --no-cache \
    # Git is required for fetching the dependencies.
    unzip

# Copy your custom PocketBase and build
COPY . /pb
WORKDIR /pb

# Note: This will pull the latest version of pocketbase. If you are just doing 
# simple customizations and don't have a local build environment for Go, 
# leave this line in. 
# For more complex builds that include other dependencies, remove this 
# line and rely on the go.sum lockfile.
#RUN go get github.com/pocketbase/pocketbase
RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" ./cmd/server/main.go
RUN ls
WORKDIR /

############################
# STEP 2 build a small image
############################
FROM alpine:latest

RUN apk add --no-cache \
    ca-certificates \
    # this is needed only if you want to use scp to copy later your pb_data locally
    openssh

# Copy our static executable.
COPY --from=builder /pb/main /pb/main

COPY /data /pb/data

# uncomment to copy the local pb_migrations dir into the image
COPY /pb_migrations /pb/pb_migrations

# uncomment to copy the local pb_hooks dir into the image
COPY /pb_hooks /pb/pb_hooks

EXPOSE 8080

# start PocketBase
CMD ["/pb/main", "serve", "--http=0.0.0.0:8080"]

FROM golang:1.16-alpine AS base

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY src/go.mod ./
COPY src/go.sum ./

RUN go mod download

# Build app
COPY src .
RUN CGO_ENABLED=0 GOOS=linux go build -o johnerator *.go

FROM alpine:latest as runner
RUN apk add ca-certificates
COPY --from=base /app/johnerator .
RUN adduser -D -g '' johnerator
RUN chmod +x johnerator
USER johnerator

ENTRYPOINT ["./johnerator"]
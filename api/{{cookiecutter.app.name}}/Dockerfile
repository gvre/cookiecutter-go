# syntax=docker/dockerfile:1
ARG ENVIRONMENT=non-prod

FROM golang:1.19-alpine3.16 AS golang
FROM alpine:3.16 AS alpine
FROM flyway/flyway:9-alpine AS flyway

# Go Builder
FROM golang AS builder
USER root
RUN apk add git gcc musl-dev
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

## Migrations
FROM flyway AS migrations
COPY migrations /migrations

# non-prod environment
FROM builder AS non-prod

## Build binary
FROM non-prod AS non-prod-builder
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=1 go build -buildvcs=true -race -trimpath -ldflags "-extldflags '-static' -s -w" -o main

# prod environment
FROM builder AS prod-builder
RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -buildvcs=true -a -trimpath -ldflags "-extldflags '-static' -s -w" -o main

# Final image alias
FROM ${ENVIRONMENT}-builder AS final

# Final image
FROM alpine
RUN adduser --disabled-password user
COPY --from=final /app/main /main
USER user

EXPOSE 8080
CMD ["/main"]

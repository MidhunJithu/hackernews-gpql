  FROM golang:alpine AS build-base
  WORKDIR /app
  COPY go.mod go.sum ./
  RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go mod download


  FROM build-base  AS dev
  RUN go install github.com/air-verse/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest
  COPY . .
  CMD ["air"]

  # FROM golang:alpine AS dev
  # WORKDIR /app
  # RUN go install github.com/air-verse/air@latest
  # COPY go.mod go.sum ./
  # RUN go mod download
  # COPY . .
  # CMD ["air", "-c", ".air.toml"]

  FROM build-base AS build-prod
  RUN useradd -u 1001 nonroot
  COPY . .
  RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o go-graphql-hackernews

  FROM scratch AS release
  WORKDIR /
  COPY --from=build-prod /etc/passwd /etc/passwd
  COPY --from=build-prod /app/go-graphql-hackernews .
  USER nonroot
  EXPOSE 8080
  CMD ["./go-graphql-hackernews"]
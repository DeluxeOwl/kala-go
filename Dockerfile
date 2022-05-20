FROM golang:1.18 as builder

WORKDIR /src
COPY . .

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /app -a -ldflags '-linkmode external -extldflags "-static"' ./cmd/kala.go

FROM scratch
COPY --from=builder /app /app

EXPOSE 1323

ENTRYPOINT ["/app"]
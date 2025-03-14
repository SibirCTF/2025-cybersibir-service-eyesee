FROM golang:alpine AS builder

RUN apk update --no-cache

WORKDIR /build

ADD ../service/go.mod .
ADD ../service/go.sum .
RUN go mod download

COPY ../service .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /app/eyesee main.go

FROM alpine:3.21.3
WORKDIR /app
RUN apk --no-cache add curl
COPY --from=builder /app/eyesee /app/
CMD ["/app/eyesee", "-minio_endpoint", "eyesee-storage:9000", "-pin", "2591", "-reflect", "-debug"]

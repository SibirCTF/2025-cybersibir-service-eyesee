FROM alpine:3.21.3
WORKDIR /app
RUN apk --no-cache add curl
ADD eyesee /app/
CMD ["/app/eyesee", "-minio_endpoint", "eyesee-storage:9000"]

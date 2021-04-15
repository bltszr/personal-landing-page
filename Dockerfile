FROM golang:1.16.3 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...


FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./main"]

EXPOSE 3000
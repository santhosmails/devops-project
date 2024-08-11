# Builder for compiling the source code
FROM golang:1.22 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main .

# Final Stage - Distroless image
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/password ./password
EXPOSE 8080
ENTRYPOINT ["./main"]

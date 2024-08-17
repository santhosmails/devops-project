# Builder for compiling the source code
FROM golang:1.22 AS builder
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
# CGO_ENABLED=0 and GOOS=linux to ensure that the Go binary is statically linked and compatible with the distroless image
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final Stage - Distroless image
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/password ./password
EXPOSE 8080
ENTRYPOINT ["./main"]

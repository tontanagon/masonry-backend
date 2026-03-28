FROM golang:1.25.6-alpine AS builder

WORKDIR /app

# Copy go module files first for better caching
COPY ./public/go.mod ./public/go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY ./public/ .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /masonry-backend

# Use a minimal image for the final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /masonry-backend .

EXPOSE 8080

CMD ["./masonry-backend"]
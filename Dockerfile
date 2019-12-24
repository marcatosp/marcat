FROM golang:1.13.5 as builder

LABEL maintainer="Min Iddamal"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:3.10.3

RUN apk add --no-cache ca-certificates

WORKDIR /root/

# Copy the pre-build binary from the prev stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]

# Stage 1: Build the Golang application
FROM golang:1.18-alpine as builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app

# Stage 2: Create the production image
FROM scratch
COPY --from=builder /go/bin/app /app
CMD ["/app"]

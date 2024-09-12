FROM golang:1.23 AS builder
WORKDIR /app
# RUN apt-get install -y ca-certificates
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o students-api cmd/main.go cmd/container.go

FROM scratch AS runner
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/.env .
COPY --from=builder /app/students-api .
EXPOSE 8080
ENTRYPOINT ["./students-api"]
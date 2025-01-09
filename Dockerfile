FROM golang:1.23.0-alphine AS builder 
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .env
EXPOSE 8080
CMD [ "./main" ]
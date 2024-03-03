FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server ./main.go

FROM alpine
WORKDIR /app
COPY --from=0 /app/server .
CMD ["./server"]
